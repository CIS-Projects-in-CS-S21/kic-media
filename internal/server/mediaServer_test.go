package server_test

import (
	"context"
	"github.com/kic/media/internal/server"
	"github.com/kic/media/pkg/cloudstorage"
	"os"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/kic/media/pkg/database"
	"github.com/kic/media/pkg/logging"
	pbcommon "github.com/kic/media/pkg/proto/common"
	pbmedia "github.com/kic/media/pkg/proto/media"
)

var log *zap.SugaredLogger
var mediaService *server.MediaStorageServer

const testDataPath = "../../test_data"

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

	mp := make(map[int]*pbcommon.File)
	repo := database.NewMockRepository(mp, log)

	prepDBForTests(repo)

	cloudStorage := cloudstorage.NewMockCloudStorage(testDataPath, log)

	mediaService = server.NewMediaStorageServer(repo, cloudStorage, log)

	exitVal := m.Run()

	os.Exit(exitVal)
}

func Test_ShouldUploadFile(t *testing.T) {
	resp, err := mediaService.UploadFile(context.Background(), &pbmedia.UploadFileRequest{
		FileInfo: &pbcommon.File{
			FileName:     "werwesdf",
			FileLocation: "werwesdf",
			Metadata:     nil,
			DateStored:   nil,
		},
		FileURI: "as456",
	})
	if err != nil {
		t.Errorf("Upload should not fail")
	}
	if resp.BytesRead != 5 {
		t.Errorf("Wrong number of bytes returned")
	}
}

func Test_ShouldFailUploadFile(t *testing.T) {
	_, err := mediaService.UploadFile(context.Background(), &pbmedia.UploadFileRequest{
		FileInfo: &pbcommon.File{
			FileName:     "",
			FileLocation: "",
			Metadata:     nil,
			DateStored:   nil,
		},
		FileURI: "as456",
	})
	if err == nil {
		t.Errorf("Upload should fail")
	}
}

func Test_ShouldCheckForFileByName(t *testing.T) {
	resp, err := mediaService.CheckForFileByName(context.Background(), &pbmedia.CheckForFileRequest{FileInfo: &pbcommon.File{
		FileName:     "tester1",
		FileLocation: "test",
		Metadata: map[string]string{
			"UID":  "12345",
			"type": "video",
		},
	}})
	if err != nil {
		t.Errorf("Check should not fail")
	}
	if resp.GetExists() != true {
		t.Errorf("Check should not fail")
	}
}

func Test_ShouldFailCheckForFileByName(t *testing.T) {
	resp, err := mediaService.CheckForFileByName(context.Background(), &pbmedia.CheckForFileRequest{FileInfo: &pbcommon.File{
		FileName:     "fakefile123",
		FileLocation: "fakefile123",
		Metadata: map[string]string{
			"UID":  "12345",
			"type": "video",
		},
	}})
	if err == nil {
		t.Errorf("Check should err with FNF")
	}
	if resp.GetExists() != false {
		t.Errorf("Check should return false")
	}
}

func Test_ShouldGetFileByMetadata(t *testing.T) {
	resp, err := mediaService.GetFilesWithMetadata(context.Background(), &pbmedia.GetFilesByMetadataRequest{
		DesiredMetadata: map[string]string{
			"type": "image",
		},
		Strictness: pbmedia.MetadataStrictness_STRICT,
	})

	if err != nil {
		t.Errorf("Should not get error for valid files")
	}

	files := resp.GetFileInfos()

	if len(files) != 2 {
		t.Errorf("Wrong number of files returned")
	}

	if files[0].FileName != "tester7" && files[0].FileName != "tester2" {
		t.Errorf("Recieved file with wrong name")
	}

	if files[1].FileName != "tester7" && files[1].FileName != "tester2" {
		t.Errorf("Recieved file with wrong name")
	}

}

func Test_ShouldNotGetFileByMetadata(t *testing.T) {
	resp, _ := mediaService.GetFilesWithMetadata(context.Background(), &pbmedia.GetFilesByMetadataRequest{
		DesiredMetadata: map[string]string{
			"type": "sdfasdfadsf",
		},
		Strictness: pbmedia.MetadataStrictness_STRICT,
	})

	if len(resp.GetFileInfos()) != 0 {
		t.Errorf("Should get no files with bad metadata")
	}
}

func Test_ShouldDeleteFileWithMetadata(t *testing.T) {
	resp, err := mediaService.DeleteFilesWithMetaData(context.Background(), &pbmedia.DeleteFilesWithMetaDataRequest{
		Metadata: map[string]string{
			"deleteMe": "true",
		},
		Strictness: 0,
	})
	if err != nil {
		t.Errorf("Should get no error with good metadata")
	}

	if resp.Success != true {
		t.Errorf("Should have deleted files with good metadata")
	}
}

func Test_ShouldNotDeleteFileWithMetadata(t *testing.T) {
	resp, _ := mediaService.DeleteFilesWithMetaData(context.Background(), &pbmedia.DeleteFilesWithMetaDataRequest{
		Metadata: map[string]string{
			"deleteMe": "false",
		},
		Strictness: 0,
	})

	if resp.Success != true {
		t.Errorf("Should have no deleted files with bad metadata")
	}
}
