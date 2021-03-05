package server

import (
	"bytes"
	"context"
	"github.com/kic/media/pkg/database"
	pbmedia "github.com/kic/media/pkg/proto/media"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

// 1 MB
const maxImageSize = 1 << 20

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
	req, err := stream.Recv()
	if err != nil {
		m.logger.Infof("%v", err)
		return status.Errorf(codes.Unknown, "File Data could not be received")
	}
	fileInfo := req.GetFileInfo()
	m.logger.Infof("FileName: %v", fileInfo.FileName)
	m.logger.Infof("FileName: %v", fileInfo.FileLocation)
	m.logger.Infof("FileName: %v", fileInfo.Metadata)

	data := bytes.Buffer{}
	bytesRead := uint64(0)

	for {
		m.logger.Info("Waiting for file byte data")

		req, err := stream.Recv()
		if err == io.EOF {
			m.logger.Info("Stream closed")
			break
		}
		if err != nil {
			m.logger.Infof("%v", err)
			return status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err)
		}

		chunk := req.GetChunk()
		size := len(chunk)

		m.logger.Infof("received a chunk with size: %d", size)

		bytesRead += uint64(size)
		if bytesRead > maxImageSize {
			return status.Errorf(codes.InvalidArgument, "file is too large: %d > %d", bytesRead, maxImageSize)
		}
		_, err = data.Write(chunk)
		if err != nil {

		}
	}

	id, err := m.db.AddFile(context.TODO(), fileInfo)

	res := &pbmedia.UploadFileResponse {
		FileID: id,
		BytesRead: bytesRead,
	}

	err = stream.SendAndClose(res)

	if err != nil {
		m.logger.Infof("%v", err)
		return status.Errorf(codes.Unknown, "cannot send response: %v", err)
	}

	m.logger.Infof("saved image with id: %s, size: %d", res.FileID, bytesRead)

	return nil
}

// Using the same format as above, the service allows the client to retrieve a stored file.
func (m *MediaStorageServer) DownloadFileByName(stuff *pbmedia.DownloadFileRequest, serv pbmedia.MediaStorage_DownloadFileByNameServer) error {
	return nil
}

// Check for the existence of a file by filename
func (m *MediaStorageServer) CheckForFileByName(ctx context.Context, req *pbmedia.CheckForFileRequest) (*pbmedia.CheckForFileResponse, error) {
	info := req.FileInfo
	m.logger.Infof("%v", info.FileName)
	file, err := m.db.GetFileWithName(ctx, req.FileInfo.FileName)
	if err != nil {
		m.logger.Infof("%v", err)
		return &pbmedia.CheckForFileResponse{
			Exists: false,
		}, err
	}
	if file.FileName == "" {
		return &pbmedia.CheckForFileResponse{
			Exists: false,
		}, nil
	}
	m.logger.Infof("%v", file.FileName)
	return &pbmedia.CheckForFileResponse{
		Exists: true,
	}, nil
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
