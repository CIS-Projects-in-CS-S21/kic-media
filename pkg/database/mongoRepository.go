package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"

	"go.uber.org/zap"

	"go.mongodb.org/mongo-driver/mongo"

	pbcommon "github.com/kic/media/pkg/proto/common"
	pbmedia "github.com/kic/media/pkg/proto/media"
)

const (
	fileCollectionName = "files"
)

type MongoRepository struct {
	client         *mongo.Client
	fileCollection *mongo.Collection

	logger *zap.SugaredLogger
}

func NewMongoRepository(client *mongo.Client, logger *zap.SugaredLogger) *MongoRepository {
	return &MongoRepository{
		client: client,
		logger: logger,
	}
}

func (m *MongoRepository) SetCollections(databaseName string) {
	m.fileCollection = m.client.Database(databaseName).Collection(fileCollectionName)
}

func (m *MongoRepository) AddFile(ctx context.Context, file *pbcommon.File) (string, error) {
	res, err := m.fileCollection.InsertOne(context.TODO(), file)
	if err != nil {
		m.logger.Infof("%v", err)
		return "", err
	}
	var toReturn string
	toReturn = res.InsertedID.(primitive.ObjectID).Hex()

	return toReturn, err
}

func (m *MongoRepository) GetFilesWithMetadata(
	ctx context.Context,
	meta map[string]string,
	strict pbmedia.MetadataStrictness,
) ([]*pbcommon.File, error) {
	return nil, nil
}

func (m *MongoRepository) GetFileWithName(ctx context.Context, fileName string) (*pbcommon.File, error) {
	filter := bson.M{
		"filename": fileName,
	}

	res := m.fileCollection.FindOne(context.TODO(), filter)

	if res.Err() == mongo.ErrNoDocuments {
		return nil, res.Err()
	}

	file := &pbcommon.File{}

	err := res.Decode(file)

	if err != nil {
		log.Println("Failed to decode file info")
	}

	log.Printf("Devoded: %v", file)
	return file, nil
}

func (m *MongoRepository) DeleteFilesWithMetadata(
	ctx context.Context,
	meta map[string]string,
	strict pbmedia.MetadataStrictness,
) error {
	return nil
}
