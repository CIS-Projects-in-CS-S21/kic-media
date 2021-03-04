package server

import (
	"context"

	"github.com/kic/media/pkg/database"
	pbmedia "github.com/kic/media/pkg/proto/media"
)

type MediaStorageServer struct {
	// required by interface for backwards compatibility with streaming methods
	pbmedia.UnimplementedMediaStorageServer
	db database.Repository
}

func (m *MediaStorageServer) UploadFile(stream pbmedia.MediaStorage_UploadFileServer) error {
	return nil
}

// Using the same format as above, the service allows the client to retrieve a stored file.
func (m *MediaStorageServer) DownloadFileByName(stuff *pbmedia.DownloadFileRequest, serv pbmedia.MediaStorage_DownloadFileByNameServer) error {
	return nil
}

// Check for the existence of a file by filename
func (m *MediaStorageServer) CheckForFileByName(context.Context, *pbmedia.CheckForFileRequest) (*pbmedia.CheckForFileResponse, error) {
	return nil, nil
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
