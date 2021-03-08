package database_test

import (
	"context"
	"github.com/kic/media/pkg/database"
	pbcommon "github.com/kic/media/pkg/proto/common"
	"os"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/kic/media/internal/setup"
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

	r, mongoClient := setup.DBRepositorySetup(log, "test-mongo-storage")

	repo = r

	defer mongoClient.Disconnect(context.Background())

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestMongoRepository_AddFile(t *testing.T) {
	t.Fail()
}

func TestMongoRepository_GetFilesWithMetadata(t *testing.T) {
	t.Fail()
}

func TestMongoRepository_GetFileWithName(t *testing.T) {
	t.Fail()
}

func TestMongoRepository_DeleteFilesWithMetadata(t *testing.T) {
	t.Fail()
}
