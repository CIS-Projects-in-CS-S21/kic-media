package database

import (
	"context"
	pbcommon "github.com/kic/media/pkg/proto/common"
	pbmedia "github.com/kic/media/pkg/proto/media"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
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

func (m *MongoRepository) GetFileWithName(ctx context.Context, fileName string) (*pbcommon.File, error) {
	filter := bson.M{
		"filename": fileName,
	}

	res := m.fileCollection.FindOne(ctx, filter)

	if res.Err() == mongo.ErrNoDocuments {
		return nil, res.Err()
	}

	file := &pbcommon.File{}

	err := res.Decode(file)

	if err != nil {
		m.logger.Debugf("Failed to decode file info")
	}

	m.logger.Debugf("Decoded: %v", file)

	return file, nil
}


func compareMetadataStrict(target, stored map[string]string) bool {
	good := true

	for key, element := range target {
		if val, ok := stored[key]; ok {
			if val != element {
				good = false
				break
			}
		} else {
			good = false
			break
		}
	}
	return good
}

func compareMetadataCasual(target, stored map[string]string) bool {
	good := false

	for key, element := range target {
		if val, ok := stored[key]; ok {
			if val == element {
				good = true
				break
			}
		}
	}
	return good
}

func (m *MongoRepository) GetFilesWithMetadata(
	ctx context.Context,
	meta map[string]string,
	strict pbmedia.MetadataStrictness,
) ([]*pbcommon.File, error) {
	toReturn := make([]*pbcommon.File, 0)

	filter := bson.M{}

	cur, err := m.fileCollection.Find(ctx, filter)
	if err != nil {
		m.logger.Errorf("Error finding files: %v", err)
	}

	for cur.Next(context.Background()) {
		file := &pbcommon.File{}
		err = cur.Decode(file)
		if err != nil {
			m.logger.Errorf("Error decoding file: %v", err)
			return toReturn, err
		}
		var res bool
		if strict == pbmedia.MetadataStrictness_STRICT {
			res = compareMetadataStrict(meta, file.Metadata)
		} else if strict == pbmedia.MetadataStrictness_CASUAL {
			res = compareMetadataCasual(meta, file.Metadata)
		}
		if res {
			toReturn = append(toReturn, file)
		}
	}
	return toReturn, nil
}

type HexId struct {
	ID primitive.ObjectID `bson:"_id"`
}

func (m *MongoRepository) DeleteFilesWithMetadata(
	ctx context.Context,
	meta map[string]string,
	strict pbmedia.MetadataStrictness,
) error {
	filter := bson.M{}

	cur, err := m.fileCollection.Find(ctx, filter)
	if err != nil {
		m.logger.Errorf("Error finding files: %v", err)
	}

	for cur.Next(context.Background()) {
		file := &pbcommon.File{}
		err = cur.Decode(file)
		if err != nil {
			m.logger.Errorf("Error decoding file: %v", err)
			return err
		}
		var res bool
		if strict == pbmedia.MetadataStrictness_STRICT {
			res = compareMetadataStrict(meta, file.Metadata)
		} else if strict == pbmedia.MetadataStrictness_CASUAL {
			res = compareMetadataCasual(meta, file.Metadata)
		}
		if res {
			var hexId HexId
			err = cur.Decode(&hexId)
			if err != nil {
				m.logger.Errorf("Error decoding file: %v", err)
				return err
			}
			_, err := m.fileCollection.DeleteOne(ctx, bson.M{"_id": hexId.ID})
			if err != nil {
				m.logger.Errorf("Error deleting file: %v", err)
				return err
			}
			m.logger.Infof("Deleted file: %v", file)

		}
	}
	return nil
}


func (m *MongoRepository) UpdateFilesWithMetadata(
	ctx context.Context,
	meta map[string]string,
	strict pbmedia.MetadataStrictness,
) error {

	filter := bson.M{}

	cur, err := m.fileCollection.Find(ctx, filter)
	if err != nil {
		m.logger.Errorf("Error finding files: %v", err)
	}

	for cur.Next(context.Background()) {

	}

	return nil;
}


func (m *MongoRepository) AddCommentToFile(
	ctx context.Context,
	meta map[string]string,
	strict pbmedia.MetadataStrictness,
) error {
	return nil;
}