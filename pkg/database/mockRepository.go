package database

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pbcommon "github.com/kic/media/pkg/proto/common"
	pbmedia "github.com/kic/media/pkg/proto/media"
)

type MockRepository struct {
	fileCollection map[int]*pbcommon.File

	idCounter int

	logger *zap.SugaredLogger
}

func NewMockRepository(fileCollection map[int]*pbcommon.File, logger *zap.SugaredLogger) *MockRepository {
	return &MockRepository{
		fileCollection: fileCollection,
		idCounter:      len(fileCollection),
		logger:         logger,
	}
}

func (m *MockRepository) AddFile(ctx context.Context, file *pbcommon.File) (string, error) {
	m.fileCollection[m.idCounter] = file
	var toReturn string
	toReturn = fmt.Sprint(m.idCounter)
	m.idCounter++

	return toReturn, nil
}

func (m *MockRepository) GetFileWithName(ctx context.Context, fileName string) (*pbcommon.File, error) {
	var toReturn *pbcommon.File
	toReturn = nil

	for _, val := range m.fileCollection {
		if val.FileName == fileName {
			toReturn = val
		}
	}

	if toReturn == nil {
		return nil, status.Errorf(codes.NotFound, "File not found")
	}

	return toReturn, nil
}

func (m *MockRepository) GetFilesWithMetadata(
	ctx context.Context,
	meta map[string]string,
	strict pbmedia.MetadataStrictness,
) ([]*pbcommon.File, error) {
	toReturn := make([]*pbcommon.File, 0)

	for _, val := range m.fileCollection {
		var res bool
		if strict == pbmedia.MetadataStrictness_STRICT {
			res = compareMetadataStrict(meta, val.Metadata)
		} else if strict == pbmedia.MetadataStrictness_CASUAL {
			res = compareMetadataCasual(meta, val.Metadata)
		}
		if res {
			toReturn = append(toReturn, val)
		}
	}

	return toReturn, nil
}

func (m *MockRepository) DeleteFilesWithMetadata(
	ctx context.Context,
	meta map[string]string,
	strict pbmedia.MetadataStrictness,
) error {

	for key, val := range m.fileCollection {
		var res bool
		if strict == pbmedia.MetadataStrictness_STRICT {
			res = compareMetadataStrict(meta, val.Metadata)
		} else if strict == pbmedia.MetadataStrictness_CASUAL {
			res = compareMetadataCasual(meta, val.Metadata)
		}
		if res {
			delete(m.fileCollection, key)
		}
	}

	return nil
}

func (m *MockRepository) UpdateFilesWithMetadata(
	ctx context.Context,
	targetMetaData map[string]string,
	desiredMetaData map[string]string,
	strict pbmedia.MetadataStrictness,
	updateFlag pbmedia.UpdateFlag,
) error {
	var err error
	for key, val := range m.fileCollection {
		var res bool
		if strict == pbmedia.MetadataStrictness_STRICT {
			res = compareMetadataStrict(targetMetaData, val.Metadata)
		} else if strict == pbmedia.MetadataStrictness_CASUAL {
			res = compareMetadataCasual(targetMetaData, val.Metadata)
		}
		if res {
			if updateFlag == pbmedia.UpdateFlag_OVERWRITE { // if overwriting metadata
				for key, value := range desiredMetaData {
					val.Metadata[key] = value
				}
			} else { // if we wish to append metadata
				err = appendMetaData(val.Metadata, desiredMetaData) // appending new metadata to existing metadata

				if err != nil {
					return err
				}

			}

			m.fileCollection[key].Metadata = val.Metadata
		}
		return err
	}
	return nil
}
