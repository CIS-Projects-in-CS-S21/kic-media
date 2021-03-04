package server

import (
	"context"
	"fmt"
	"go.uber.org/zap"

	"github.com/kic/media/pkg/database"
	pbmedia "github.com/kic/media/pkg/proto/media"
)

type MediaStorageServer struct {
	// required by interface for backwards compatibility with streaming methods
	pbmedia.UnimplementedMediaStorageServer
	db database.Repository
	logger *zap.SugaredLogger
}

func NewMediaStorageServer(db database.Repository, logger *zap.SugaredLogger) *MediaStorageServer {
	return &MediaStorageServer{
		UnimplementedMediaStorageServer: pbmedia.UnimplementedMediaStorageServer{},
		db:                              db,
		logger: 						 logger,
	}
}

func (m *MediaStorageServer) UploadFile(stream pbmedia.MediaStorage_UploadFileServer) error {
	return nil
}

// Using the same format as above, the service allows the client to retrieve a stored file.
func (m *MediaStorageServer) DownloadFileByName(stuff *pbmedia.DownloadFileRequest, serv pbmedia.MediaStorage_DownloadFileByNameServer) error {
	return nil
}

// Check for the existence of a file by filename
func (m *MediaStorageServer) CheckForFileByName(ctx context.Context, req *pbmedia.CheckForFileRequest) (*pbmedia.CheckForFileResponse, error) {
	info := req.FileInfo
	out := fmt.Sprintf("%v", info.FileLocation)
	m.logger.Info(out)
	return &pbmedia.CheckForFileResponse{}, nil
}

// Allows for the requesting of files with specific key value pairs as metadata. The strictness can be set
// such that for example only perfect matches will be returned.
func (m *MediaStorageServer) GetFilesWithMetadata(context.Context, *pbmedia.GetFilesByMetadataRequest) (*pbmedia.GetFilesByMetadataResponse, error) {
	return nil, nil
}

// Using the same strictness settings as the above, delete particular files with certain metadata.
func (m *MediaStorageServer) DeleteFilesWithMetaData(context.Context, *pbmedia.DeleteFilesWithMetaDataRequest) (*pbmedia.DeleteFilesWithMetaDataResponse, error) {
	return nil, nil
}
