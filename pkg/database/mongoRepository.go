package database

import (
	"context"
	"errors"
	pbcommon "github.com/kic/media/pkg/proto/common"
	pbmedia "github.com/kic/media/pkg/proto/media"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"strings"
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
	targetMetaData map[string]string,
	desiredMetaData map[string]string,
	strict pbmedia.MetadataStrictness,
	updateFlag pbmedia.UpdateFlag,
) error {
	filter := bson.M{}

	cur, err := m.fileCollection.Find(ctx, filter)
	if err != nil {
		m.logger.Errorf("Error finding files: %v", err)
	}

	for cur.Next(context.TODO()) {
		file := &pbcommon.File{}
		err = cur.Decode(file)
		if err != nil {
			m.logger.Errorf("Error decoding file: %v", err)
			return err
		}

		var res bool
		if strict == pbmedia.MetadataStrictness_STRICT {
			res = compareMetadataStrict(targetMetaData, file.Metadata)
		} else if strict == pbmedia.MetadataStrictness_CASUAL {
			res = compareMetadataCasual(targetMetaData, file.Metadata)
		}
		if res {
			if updateFlag == pbmedia.UpdateFlag_OVERWRITE { // if overwriting metadata
				for key, value := range desiredMetaData {
					file.Metadata[key] = value
				}
			} else { // if we wish to append metadata
				err = appendMetaData(file.Metadata, desiredMetaData) // appending new metadata to existing metadata

				if err != nil {
					return err
				}
			}

			// getting id of document to update
			var hexId HexId
			err = cur.Decode(&hexId)

			// filter to be used to locate the entry to be updated
			filter := bson.M{
				"_id": hexId.ID,
			}

			// update structure which will be used to update entry in MongoDB database
			update := bson.M{
				"$set": bson.M{
					"metadata": file.Metadata,
				},
			}

			// updating the db with the new data
			_, err := m.fileCollection.UpdateOne(
				context.Background(),
				filter,
				update,
			)

			if err != nil {
				return err
			}

		}
	}
	return nil
}

func appendMetaData(target, appendage map[string]string) error {
	for key, value := range appendage { // iterating through entries in map of new metadata
		if val, ok := target[key]; ok {
			if val[0] == '[' { // if it's a list
				bracketPos := strings.Index(val, "]") // getting position of ]
				target[key] = val[:bracketPos] + "," + val + "]"
			} else if val[0] == '{' { // if it's a dictionary
				bracketPos := strings.Index(val, "}") // getting position of }
				target[key] = val[:bracketPos] + "," + val + "}"
			} else { // if it's neither a list nor a dictionary
				return errors.New("trying to append to a scalar value")
			}
		} else {
			target[key] = value
		}
	}
	return nil
}
