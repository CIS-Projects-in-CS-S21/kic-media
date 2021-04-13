package database_test

import (
	"context"
	"github.com/kic/media/pkg/database"
	pbcommon "github.com/kic/media/pkg/proto/common"
	pbmedia "github.com/kic/media/pkg/proto/media"
	"os"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/kic/media/pkg/logging"
)

var log *zap.SugaredLogger

var repo database.Repository

func prepDBForTests() {

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
	}

	for _, file := range filesToAdd {
		id, err := repo.AddFile(context.Background(), file)
		log.Debugf("inserted id: %v", id)
		if err != nil {
			log.Debugf("insertion error: %v", err)
		}
	}
}

func TestMain(m *testing.M) {
	time.Sleep(1 * time.Second)
	log = logging.CreateLogger(zapcore.DebugLevel)

	// r, mongoClient := setup.DBRepositorySetup(log, "test-mongo-storage")

	fileCollection := map[int]*pbcommon.File{}
	r := database.NewMockRepository(fileCollection, log)

	repo = r

	prepDBForTests()

	// defer mongoClient.Disconnect(context.Background())

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestMongoRepository_GetFileWithName(t *testing.T) {
	filesToCheck := []*pbcommon.File{
		{
			FileName:     "tester1",
		},
		{
			FileName:     "tester2",
		},
		{
			FileName:     "tester3",
		},
	}

	notThereFiles := []*pbcommon.File{
		{
			FileName:     "notThere123",
		},
	}

	for i, file := range filesToCheck {
		file, err := repo.GetFileWithName(context.Background(), file.FileName)

		if err != nil || file == nil {
			t.Errorf("Test %v failed with err: %v", i, err)
		}
	}

	for i, file := range notThereFiles {
		file, err := repo.GetFileWithName(context.Background(), file.FileName)

		if err == nil || file != nil {
			t.Errorf("Test %v succeeded but should not have", i)
		}
	}
}

func TestMongoRepository_DeleteFilesWithMetadata(t *testing.T) {
	meta := map[string]string{
		"deleteMe": "true",
	}
	err := repo.DeleteFilesWithMetadata(context.Background(), meta, pbmedia.MetadataStrictness_STRICT)

	if err != nil {
		t.Errorf("Failed to delete files: %v", err)
	}

	files, err := repo.GetFilesWithMetadata(context.Background(), meta, pbmedia.MetadataStrictness_STRICT)

	if err != nil {
		t.Errorf("Failed to delete files: %v", err)
	}

	if len(files) > 0 {
		t.Error("Failed to delete files")
	}
}

func TestMongoRepository_GetFilesWithMetadata(t *testing.T) {
	meta := map[string]string{
		"UID": "12345",
	}
	files, err := repo.GetFilesWithMetadata(context.TODO(), meta, pbmedia.MetadataStrictness_CASUAL)

	if err != nil {
		t.Errorf("Failed to get files: %v", err)
	}

	if len(files) == 0 {
		t.Error("Failed to get files")
	}
}

func TestMongoRepository_UpdateFilesWithMetadata(t *testing.T) {
	target := map[string]string{
		"UID": "12345",
	}

	desired := map[string]string{
		"UID": "54321",
	}
	err := repo.UpdateFilesWithMetadata(context.TODO(), target, desired, pbmedia.MetadataStrictness_CASUAL, pbmedia.UpdateFlag_OVERWRITE)

	if err != nil {
		t.Errorf("Failed to update files: %v", err)
	}

	files, err := repo.GetFilesWithMetadata(context.Background(), desired, pbmedia.MetadataStrictness_STRICT)

	if err != nil {
		t.Errorf("Failed to delete files: %v", err)
	}

	if len(files) == 0 {
		t.Error("Failed to update files")
	}
}
