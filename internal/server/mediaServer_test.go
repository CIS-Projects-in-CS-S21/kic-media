package server_test

import (
	"bytes"
	"context"
	"github.com/kic/media/internal/server"
	"github.com/kic/media/pkg/cloudstorage"
	"google.golang.org/grpc/reflection"
	"io"
	"io/ioutil"
	"net"
	"os"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"

	"github.com/kic/media/internal/setup"
	"github.com/kic/media/pkg/database"
	"github.com/kic/media/pkg/logging"
	pbcommon "github.com/kic/media/pkg/proto/common"
	pbmedia "github.com/kic/media/pkg/proto/media"
)

var log *zap.SugaredLogger
var client pbmedia.MediaStorageClient

const testDataPath = "../../test_data"

type testGetFilesWithMetadata struct {
	req       *pbmedia.GetFilesByMetadataRequest
	res       *pbmedia.GetFilesByMetadataResponse
	shouldErr bool
}

type testUploadFile struct {
	filePath   string
	uploadPath string
	checkPath  string
	shouldErr  bool
}

func prepDBForTests(db database.Repository) {

	filesToAdd := []*pbcommon.File{
		{
			FileName:     "tester1",
			FileLocation: "test",
			Metadata: map[string]string{
				"UID":  "12345",
				"type": "video",
			},
		},
		{
			FileName:     "tester2",
			FileLocation: "test",
			Metadata: map[string]string{
				"UID":  "12345",
				"type": "image",
			},
		},
		{
			FileName:     "tester3",
			FileLocation: "test",
			Metadata: map[string]string{
				"UID":      "12345",
				"comments": "12",
			},
		},
		{
			FileName:     "tester4",
			FileLocation: "test",
			Metadata: map[string]string{
				"UID":      "deleteMe",
				"deleteMe": "true",
			},
		},
		{
			FileName:     "tester5",
			FileLocation: "test",
			Metadata: map[string]string{
				"UID":      "deleteMe",
				"deleteMe": "true",
			},
		},
		{
			FileName:     "tester6",
			FileLocation: "test",
			Metadata: map[string]string{
				"UID":      "deleteMe",
				"deleteMe": "true",
			},
		},
		{
			FileName:     "tester7",
			FileLocation: "test",
			Metadata: map[string]string{
				"UID":  "123",
				"type": "image",
			},
		},
		{
			FileName:     "Animals-Dog-icon.png",
			FileLocation: "test",
			Metadata:     map[string]string{},
		},
		{
			FileName:     "term.png",
			FileLocation: "test",
			Metadata:     map[string]string{},
		},
	}

	for _, file := range filesToAdd {
		id, err := db.AddFile(context.Background(), file)
		log.Debugf("inserted id: %v", id)
		if err != nil {
			log.Debugf("insertion error: %v", err)
		}
	}

}

func TestMain(m *testing.M) {
	time.Sleep(1 * time.Second)
	log = logging.CreateLogger(zapcore.DebugLevel)

	repo, mongoClient := setup.DBRepositorySetup(log, "test-media-storage")

	prepDBForTests(repo)

	ListenAddress := "localhost:50051"

	listener, err := net.Listen("tcp", ListenAddress)
	if err != nil {
		log.Fatalf("Unable to listen on %v: %v", ListenAddress, err)
	}

	grpcServer := grpc.NewServer()

	cloudStorage := cloudstorage.NewMockCloudStorage(testDataPath, log)

	mediaService := server.NewMediaStorageServer(repo, cloudStorage, log)
	pbmedia.RegisterMediaStorageServer(grpcServer, mediaService)

	reflection.Register(grpcServer)

	go func() {
		defer listener.Close()
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	log.Infof("Server started on %v", ListenAddress)

	defer grpcServer.Stop()
	defer mongoClient.Disconnect(context.Background())

	conn, err := grpc.Dial(ListenAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client = pbmedia.NewMediaStorageClient(conn)

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestMediaStorageServer_CheckForFileByName(t *testing.T) {
	correctRequests := []*pbmedia.CheckForFileRequest{
		{
			FileInfo: &pbcommon.File{
				FileName:     "tester1",
				FileLocation: "test",
				Metadata:     nil,
			},
		},
		{
			FileInfo: &pbcommon.File{
				FileName:     "tester2",
				FileLocation: "test",
				Metadata:     nil,
			},
		},
	}

	incorrectRequests := []*pbmedia.CheckForFileRequest{
		{
			FileInfo: &pbcommon.File{
				FileName:     "notThere1",
				FileLocation: "test",
				Metadata:     nil,
			},
		},
		{
			FileInfo: &pbcommon.File{
				FileName:     "notThere2",
				FileLocation: "test",
				Metadata:     nil,
			},
		},
	}

	for _, req := range correctRequests {
		res, err := client.CheckForFileByName(context.Background(), req)
		log.Debugf("%v %v", res, err)
		if err != nil || res.Exists != true {
			t.Errorf("File that should be in db is reported as not there")
		}
	}

	for _, req := range incorrectRequests {
		res, err := client.CheckForFileByName(context.Background(), req)
		log.Debugf("%v %v", res, err)
		if err == nil {
			t.Errorf("File that should not be in db is reported as there")
		}
	}

}

func TestMediaStorageServer_DeleteFilesWithMetaData(t *testing.T) {

	meta := map[string]string{
		"UID":  "deleteMe",
	}
	_, err := client.DeleteFilesWithMetaData(context.Background(), &pbmedia.DeleteFilesWithMetaDataRequest{
		Metadata: meta,
		Strictness: pbmedia.MetadataStrictness_STRICT,
	})

	if err != nil {
		t.Errorf("Failed to delete files: %v", err)
	}

	checkResp, err := client.GetFilesWithMetadata(context.Background(), &pbmedia.GetFilesByMetadataRequest{
		DesiredMetadata: meta,
		Strictness: pbmedia.MetadataStrictness_STRICT,
	})

	if err != nil {
		t.Errorf("Failed to delete files: %v", err)
	}

	if len(checkResp.FileInfos) > 0 {
		t.Error("Failed to delete files")
	}
}

func compareFileLists(l1, l2 []*pbcommon.File) bool {
	if len(l1) != len(l2) {
		return false
	}

	for index := range l1 {
		if l1[index].FileName != l2[index].FileName {
			return false
		}
		if l1[index].FileLocation != l2[index].FileLocation {
			return false
		}

		for key, element := range l1[index].Metadata {
			if val, ok := l2[index].Metadata[key]; ok {
				if val != element {
					return false
				}
			} else {
				return false
			}
		}
	}

	return true
}

func TestMediaStorageServer_GetFilesWithMetadata(t *testing.T) {
	tests := []testGetFilesWithMetadata{
		{
			req: &pbmedia.GetFilesByMetadataRequest{
				DesiredMetadata: map[string]string{
					"type": "image",
				},
				Strictness: pbmedia.MetadataStrictness_STRICT,
			},
			res: &pbmedia.GetFilesByMetadataResponse{
				FileInfos: []*pbcommon.File{
					{
						FileName:     "tester2",
						FileLocation: "test",
						Metadata: map[string]string{
							"UID":  "12345",
							"type": "image",
						},
					},
					{
						FileName:     "tester7",
						FileLocation: "test",
						Metadata: map[string]string{
							"UID":  "123",
							"type": "image",
						},
					},
				},
			},
			shouldErr: false,
		},

		{
			req: &pbmedia.GetFilesByMetadataRequest{
				DesiredMetadata: map[string]string{
					"UID":  "12345",
					"type": "image",
				},
				Strictness: pbmedia.MetadataStrictness_CASUAL,
			},
			res: &pbmedia.GetFilesByMetadataResponse{
				FileInfos: []*pbcommon.File{
					{
						FileName:     "tester1",
						FileLocation: "test",
						Metadata: map[string]string{
							"UID":  "12345",
							"type": "video",
						},
					},
					{
						FileName:     "tester2",
						FileLocation: "test",
						Metadata: map[string]string{
							"UID":  "12345",
							"type": "image",
						},
					},
					{
						FileName:     "tester3",
						FileLocation: "test",
						Metadata: map[string]string{
							"UID":      "12345",
							"comments": "12",
						},
					},
					{
						FileName:     "tester7",
						FileLocation: "test",
						Metadata: map[string]string{
							"UID":  "123",
							"type": "image",
						},
					},
				},
			},
			shouldErr: false,
		},
	}

	for i, test := range tests {
		resp, err := client.GetFilesWithMetadata(context.Background(), test.req)

		if err == nil && test.shouldErr {
			t.Errorf("Test %v should err but did not", i)
		} else if err != nil {
			t.Errorf("Test %v should not err but did: %v", i, err)
		} else if !test.shouldErr && !compareFileLists(resp.FileInfos, test.res.FileInfos) {
			t.Errorf("Test %v did not get the correct response\nresp: %v\ndesired: %v\n", i, resp, test.res)
		}
	}
}

func TestMediaStorageServer_DownloadFileByName(t *testing.T) {

	tests := []testUploadFile{
		{
			uploadPath: "Animals-Dog-icon.png",
			checkPath:  "../../test_data/Animals-Dog-icon.png",
			shouldErr:  false,
		},
		{
			uploadPath: "term.png",
			checkPath:  "../../test_data/term.png",
			shouldErr:  false,
		},
	}

	for i, test := range tests {
		stream, err := client.DownloadFileByName(context.Background(), &pbmedia.DownloadFileRequest{FileInfo: &pbcommon.File{
			FileName: test.uploadPath,
		}})

		if err != nil {
			t.Errorf("Test %v download file failure: %v", i, err)
		}

		var buf []byte
		buff := bytes.NewBuffer(buf)

		for {
			recv, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Errorf("Test %v download file failure: %v", i, err)
			}
			buff.Write(recv.GetChunk())
		}

		fo, err := os.Open(test.checkPath)
		if err != nil {
			t.Errorf("Test %v download file failure: %v", i, err)
		}
		rec, err := ioutil.ReadAll(fo)
		if err != nil {
			t.Errorf("Test %v upload file failure: %v", i, err)
		}

		if bytes.Compare(buff.Bytes(), rec) != 0 {
			t.Errorf("Test %v download file failure", i)
		}
	}
}

func TestMediaStorageServer_UploadFile(t *testing.T) {
	testFiles := []*pbcommon.File{
		&pbcommon.File{
			FileName:     "Animals-Dog-icon.png",
			FileLocation: "../../test_data/Animals-Dog-icon.png",
			Metadata:     nil,
			DateStored:   &pbcommon.Date{
				Year:  2021,
				Month: 4,
				Day:   12,
			},
		},

		&pbcommon.File{
			FileName:     "term.png",
			FileLocation: "../../test_data/term.png",
			Metadata:     nil,
			DateStored:   &pbcommon.Date{
				Year:  2021,
				Month: 4,
				Day:   12,
			},
		},
	}



	for i, test := range testFiles {
		resp, err := sendFile(test, test.FileLocation)
		if err != nil || resp.BytesRead == 0 {
			t.Errorf("Test %v upload file failure", i)
		}
		fi, err := os.Open(test.FileLocation)
		if err != nil {
			t.Errorf("Test %v upload file failure: %v", i, err)
		}
		fo, err := os.Open(test.FileLocation)
		if err != nil {
			t.Errorf("Test %v upload file failure: %v", i, err)
		}
		sent, err := ioutil.ReadAll(fi)
		if err != nil {
			t.Errorf("Test %v upload file failure: %v", i, err)
		}
		rec, err := ioutil.ReadAll(fo)
		if err != nil {
			t.Errorf("Test %v upload file failure: %v", i, err)
		}

		if bytes.Compare(sent, rec) != 0 {
			t.Errorf("Test %v upload file failure", i)
		}
	}
}

func sendFile(fileInfo *pbcommon.File, uploadName string) (*pbmedia.UploadFileResponse, error) {
	file, err := os.Open(fileInfo.FileLocation)
	if err != nil {
		log.Fatal("cannot open file: ", err)
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fileBytes, err := ioutil.ReadFile(fileInfo.FileLocation)

	req := &pbmedia.UploadFileRequest{
		FileInfo: fileInfo,
		File:     fileBytes,
	}

	res, err := client.UploadFile(ctx, req)
	if err != nil {
		log.Fatal("cannot upload file: ", err)
	}


	return res, err
}
