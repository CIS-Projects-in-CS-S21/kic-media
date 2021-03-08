package server_test

import (
	"context"
	"github.com/kic/media/internal/server"
	"github.com/kic/media/pkg/cloudstorage"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"reflect"
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

const testDataPath = "testdata"

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
		if err != nil || res.Exists == true {
			t.Errorf("File that should not be in db is reported as there")
		}
	}

}

func TestMediaStorageServer_DeleteFilesWithMetaData(t *testing.T) {
	t.Fail()
}

type testGetFilesWithMetadata struct {
	req *pbmedia.GetFilesByMetadataRequest
	res *pbmedia.GetFilesByMetadataResponse
	shouldErr bool
}

func TestMediaStorageServer_GetFilesWithMetadataStrict(t *testing.T) {

	tests := []testGetFilesWithMetadata{
		{
			req: &pbmedia.GetFilesByMetadataRequest{
				DesiredMetadata: map[string]string{
					"type": "image",
				},
				Strictness:      pbmedia.MetadataStrictness_STRICT,
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
				},
			},
			shouldErr: false,
		},

		{
			req: &pbmedia.GetFilesByMetadataRequest{
				DesiredMetadata: map[string]string{
					"UID": "12345",
					"type": "image",
				},
				Strictness:      pbmedia.MetadataStrictness_CASUAL,
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
		} else if !test.shouldErr && !reflect.DeepEqual(resp, test.res) {
			t.Errorf("Test %v did not get the correct response\nresp: %v\ndesired: %v\n", i, resp, test.res)
		}
	}
}

func TestMediaStorageServer_DownloadFileByName(t *testing.T) {
	t.Fail()
}

func TestMediaStorageServer_UploadFile(t *testing.T) {
	t.Fail()
}
