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

// MediaStorageServer - Implements the generated interface to be a media storage server handler
type MediaStorageServer struct {
	// required by interface for backwards compatibility with streaming methods
	pbmedia.UnimplementedMediaStorageServer
	db         database.Repository
	cloudStore cloudstorage.CloudStorage
	logger     *zap.SugaredLogger
}

// NewMediaStorageServer - create a new instance of a MediaStorageServer struct. Requires a connected database
// driver which implements the database.Repository interface, a cloud store which implements the
// cloudstorage.CloudStorage interface, and a logger instance
func NewMediaStorageServer(db database.Repository, cloudStore cloudstorage.CloudStorage, logger *zap.SugaredLogger) *MediaStorageServer {
	return &MediaStorageServer{
		UnimplementedMediaStorageServer: pbmedia.UnimplementedMediaStorageServer{},
		db:                              db,
		cloudStore:                      cloudStore,
		logger:                          logger,
	}
}

// UploadFile - allows a client to upload a file in a stream the server which will then
// store the file information in the database, and store the file in cloud storage
func (m *MediaStorageServer) UploadFile(stream pbmedia.MediaStorage_UploadFileServer) error {
	req, err := stream.Recv()
	if err != nil {
		m.logger.Infof("%v", err)
		return status.Errorf(codes.Unknown, "File Data could not be received")
	}
	fileInfo := req.GetFileInfo()
	m.logger.Debugf("FileName: %v", fileInfo.FileName)
	m.logger.Debugf("FileLocation: %v", fileInfo.FileLocation)
	m.logger.Debugf("Metadata: %v", fileInfo.Metadata)

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

// DownloadFileByName - allows a client to request a file for download from cloud storage
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
		stream.Send(nil) // returning nil response when the filename is left blank
		return status.Errorf(codes.NotFound, "File name field left empty")
	}

	buffer, err := m.cloudStore.DownloadFile(file.FileName)

	if err != nil {
		m.logger.Infof("%v", err)
		return status.Errorf(codes.Internal, "Could not access file in storage: %v", err)
	}

	// Stream a number of bytes to our
	numMessages := int(math.Ceil(float64(buffer.Len()) / float64(packetSize)))

	for i := 0; i < numMessages; i++ {
		toSend := buffer.Next(packetSize)

		err := stream.Send(&pbmedia.DownloadFileResponse{
			Chunk: toSend,
		})

		if err != nil {
			m.logger.Infof("%v", err)
			return status.Errorf(codes.Internal, "cannot send chunk data: %v", err)
		}
	}

	return nil
}

// CheckForFileByName - allows a client to check if the database has a particular file stored in it
func (m *MediaStorageServer) CheckForFileByName(ctx context.Context, req *pbmedia.CheckForFileRequest) (*pbmedia.CheckForFileResponse, error) {
	info := req.FileInfo
	m.logger.Debugf("Checking for %v", info.FileName)
	file, err := m.db.GetFileWithName(ctx, req.FileInfo.FileName)

	// we failed to find the file either with an error or by returning a nil file
	if file == nil || err != nil {
		m.logger.Infof("%v", err)
		return &pbmedia.CheckForFileResponse{
			Exists: false,
		}, status.Errorf(codes.NotFound, "File not found")
	}

	m.logger.Debugf("Found this file: %v", file.FileName)
	return &pbmedia.CheckForFileResponse{
		Exists: true,
	}, nil
}

// GetFilesWithMetadata - Allows for the requesting of files with specific key value pairs as metadata.
// The strictness can be set such that for example only perfect matches will be returned.
func (m *MediaStorageServer) GetFilesWithMetadata(
	ctx context.Context,
	req *pbmedia.GetFilesByMetadataRequest,
) (*pbmedia.GetFilesByMetadataResponse, error) {
	fileSlice, err := m.db.GetFilesWithMetadata(ctx, req.DesiredMetadata, req.Strictness)

	if err != nil {
		m.logger.Infof("%v", err)
		return &pbmedia.GetFilesByMetadataResponse{
			FileInfos: nil,
		}, status.Errorf(codes.Internal, "Error finding or decoding files with metadata")
	}

	return &pbmedia.GetFilesByMetadataResponse{
		FileInfos: fileSlice,
	}, nil
}

// DeleteFilesWithMetaData - Delete files in the database with the passed metadata
func (m *MediaStorageServer) DeleteFilesWithMetaData(
	ctx context.Context,
	req *pbmedia.DeleteFilesWithMetaDataRequest,
) (*pbmedia.DeleteFilesWithMetaDataResponse, error) {
	err := m.db.DeleteFilesWithMetadata(ctx, req.Metadata, req.Strictness)

	if err != nil {
		m.logger.Infof("%v", err)
		return &pbmedia.DeleteFilesWithMetaDataResponse{
			Success: false,
		}, status.Errorf(codes.Internal, "Error finding or deleting files")
	}

	return &pbmedia.DeleteFilesWithMetaDataResponse{Success: true}, nil
}

func (m *MediaStorageServer) UpdateFilesWithMetadata(
	ctx context.Context,
	req *pbmedia.UpdateFilesWithMetadataRequest,
) (*pbmedia.UpdateFilesWithMetadataResponse, error) {
	err := m.db.UpdateFilesWithMetadata(ctx, req.FilterMetadata, req.DesiredMetadata, req.Strictness, req.UpdateFlag)

	// If error, return empty response and err
	if err != nil {
		return &pbmedia.UpdateFilesWithMetadataResponse{
		}, status.Errorf(codes.Internal, "Error updating file metadata")
	}

	res := &pbmedia.UpdateFilesWithMetadataResponse{NumFilesUpdated: 1}

	return res, err
}
