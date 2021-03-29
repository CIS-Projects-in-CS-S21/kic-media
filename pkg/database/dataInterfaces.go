package database

import (
	"context"

	pbcommon "github.com/kic/media/pkg/proto/common"
	pbmedia "github.com/kic/media/pkg/proto/media"
)

// Repository - interface for a data provider that interfaces between the database backend and the grpc server
// enables the repository pattern so that we can swap out the database backend easily
type Repository interface {
	AddFile(context.Context, *pbcommon.File) (string, error)
	GetFilesWithMetadata(context.Context, map[string]string, pbmedia.MetadataStrictness) ([]*pbcommon.File, error)
	GetFileWithName(context.Context, string) (*pbcommon.File, error)
	DeleteFilesWithMetadata(context.Context, map[string]string, pbmedia.MetadataStrictness) error
	UpdateFilesWithMetadata(ctx context.Context, targetMetaData map[string]string, desiredMetaData map[string]string, strict pbmedia.MetadataStrictness, updateFlag pbmedia.UpdateFlag) error
}
