package server

import (
	"bytes"
	"context"
	"io"
	"math"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/kic/media/pkg/cloudstorage"
	"github.com/kic/media/pkg/database"
	pbmedia "github.com/kic/media/pkg/proto/media"
)

const (
	// 2 MB
	maxImageSize = 2 << 20
	// Size of each message when we server-side stream a file to a client
	packetSize = 1024
)

type MediaStorageServer struct {
	// required by interface for backwards compatibility with streaming methods
	pbmedia.UnimplementedMediaStorageServer
	db         database.Repository
	cloudStore cloudstorage.CloudStorage
	logger     *zap.SugaredLogger
}

func NewMediaStorageServer(db database.Repository, cloudStore cloudstorage.CloudStorage, logger *zap.SugaredLogger) *MediaStorageServer {
	return &MediaStorageServer{
		UnimplementedMediaStorageServer: pbmedia.UnimplementedMediaStorageServer{},
		db:                              db,
		cloudStore:                      cloudStore,
		logger:                          logger,
	}
}

func (m *MediaStorageServer) UploadFile(stream pbmedia.MediaStorage_UploadFileServer) error {
	req, err := stream.Recv()
	if err != nil {
		m.logger.Debugf("%v", err)
		return status.Errorf(codes.Unknown, "File Data could not be received")
	}
	fileInfo := req.GetFileInfo()
	m.logger.Debugf("FileName: %v", fileInfo.FileName)
	m.logger.Debugf("FileName: %v", fileInfo.FileLocation)
	m.logger.Debugf("FileName: %v", fileInfo.Metadata)

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

		m.logger.Debugf("received a chunk with size: %d", size)

		bytesRead += uint64(size)
		if bytesRead > maxImageSize {
			m.logger.Infof("%v", err)
			return status.Errorf(codes.InvalidArgument, "file is too large: %d > %d", bytesRead, maxImageSize)
		}
		_, err = data.Write(chunk)
		if err != nil {
			m.logger.Infof("%v", err)
			return status.Errorf(codes.Internal, "Could not write file")
		}
	}

	err = m.cloudStore.UploadFile(fileInfo.FileName, data)

	if err != nil {
		m.logger.Infof("%v", err)
		return status.Errorf(codes.Internal, "Could not write file")
	}

	id, err := m.db.AddFile(context.TODO(), fileInfo)

	res := &pbmedia.UploadFileResponse{
		FileID:    id,
		BytesRead: bytesRead,
	}

	err = stream.SendAndClose(res)

	if err != nil {
		m.logger.Infof("%v", err)
		return status.Errorf(codes.Unknown, "cannot send response: %v", err)
	}

	m.logger.Debugf("saved image with id: %s, size: %d", res.FileID, bytesRead)

	return nil
}

// Using the same format as above, the service allows the client to retrieve a stored file.
func (m *MediaStorageServer) DownloadFileByName(
	req *pbmedia.DownloadFileRequest,
	stream pbmedia.MediaStorage_DownloadFileByNameServer,
) error {
	fileInfo := req.GetFileInfo()

	file, err := m.db.GetFileWithName(context.TODO(), fileInfo.FileName)

	if err != nil {
		return err
	}

	if file.FileName == "" {
		stream.Send(&pbmedia.DownloadFileResponse{
			Data: &pbmedia.DownloadFileResponse_Error{
				Error: pbmedia.DownloadFileByNameError_FILE_NOT_FOUND,
			},
		})
	}

	buffer, err := m.cloudStore.DownloadFile(file.FileName)

	if err != nil {
		m.logger.Debugf("%v", err)
		return status.Errorf(codes.Internal, "Could not access file in storage: %v", err)
	}

	// Stream a number of bytes to our
	numMessages := int(math.Ceil(float64(buffer.Len()) / float64(packetSize)))

	for i := 0; i < numMessages; i++ {
		toSend := buffer.Next(packetSize)

		err := stream.Send(&pbmedia.DownloadFileResponse{
			Data: &pbmedia.DownloadFileResponse_Chunk{
				Chunk: toSend,
			},
		})

		if err != nil {
			m.logger.Debugf("%v", err)
			return status.Errorf(codes.Internal, "cannot send chunk data: %v", err)
		}
	}

	return nil
}

// Check for the existence of a file by filename
func (m *MediaStorageServer) CheckForFileByName(ctx context.Context, req *pbmedia.CheckForFileRequest) (*pbmedia.CheckForFileResponse, error) {
	info := req.FileInfo
	m.logger.Debugf("Checking for %v", info.FileName)
	file, err := m.db.GetFileWithName(ctx, req.FileInfo.FileName)

	// we failed to find the file either with an error or by returning a file with an empty name
	if file == nil || err != nil {
		m.logger.Debugf("%v", err)
		return &pbmedia.CheckForFileResponse{
			Exists: false,
		}, nil
	}

	m.logger.Debugf("Found this file: %v", file.FileName)
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
