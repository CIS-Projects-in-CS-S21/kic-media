// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MediaStorageClient is the client API for MediaStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MediaStorageClient interface {
	//
	// Send a file as a stream of messages, starting with a message containing a File message, then
	// followed by an arbitrary number of messages containing bytes representing the file. The response
	// will then confirm the number of bytes received or provide an error.
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (MediaStorage_UploadFileClient, error)
	//
	// Using the same format as above, the service allows the client to retrieve a stored file.
	DownloadFileByName(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (MediaStorage_DownloadFileByNameClient, error)
	// Check for the existence of a file by filename
	CheckForFileByName(ctx context.Context, in *CheckForFileRequest, opts ...grpc.CallOption) (*CheckForFileResponse, error)
	//
	// Allows for the requesting of files with specific key value pairs as metadata. The strictness can be set
	// such that for example only perfect matches will be returned.
	GetFilesWithMetadata(ctx context.Context, in *GetFilesByMetadataRequest, opts ...grpc.CallOption) (*GetFilesByMetadataResponse, error)
	//
	// Using the same strictness settings as the above, delete particular files with certain metadata.
	DeleteFilesWithMetaData(ctx context.Context, in *DeleteFilesWithMetaDataRequest, opts ...grpc.CallOption) (*DeleteFilesWithMetaDataResponse, error)
}

type mediaStorageClient struct {
	cc grpc.ClientConnInterface
}

func NewMediaStorageClient(cc grpc.ClientConnInterface) MediaStorageClient {
	return &mediaStorageClient{cc}
}

func (c *mediaStorageClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (MediaStorage_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MediaStorage_serviceDesc.Streams[0], "/kic.media.MediaStorage/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &mediaStorageUploadFileClient{stream}
	return x, nil
}

type MediaStorage_UploadFileClient interface {
	Send(*UploadFileRequest) error
	CloseAndRecv() (*UploadFileResponse, error)
	grpc.ClientStream
}

type mediaStorageUploadFileClient struct {
	grpc.ClientStream
}

func (x *mediaStorageUploadFileClient) Send(m *UploadFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mediaStorageUploadFileClient) CloseAndRecv() (*UploadFileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mediaStorageClient) DownloadFileByName(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (MediaStorage_DownloadFileByNameClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MediaStorage_serviceDesc.Streams[1], "/kic.media.MediaStorage/DownloadFileByName", opts...)
	if err != nil {
		return nil, err
	}
	x := &mediaStorageDownloadFileByNameClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MediaStorage_DownloadFileByNameClient interface {
	Recv() (*DownloadFileResponse, error)
	grpc.ClientStream
}

type mediaStorageDownloadFileByNameClient struct {
	grpc.ClientStream
}

func (x *mediaStorageDownloadFileByNameClient) Recv() (*DownloadFileResponse, error) {
	m := new(DownloadFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mediaStorageClient) CheckForFileByName(ctx context.Context, in *CheckForFileRequest, opts ...grpc.CallOption) (*CheckForFileResponse, error) {
	out := new(CheckForFileResponse)
	err := c.cc.Invoke(ctx, "/kic.media.MediaStorage/CheckForFileByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaStorageClient) GetFilesWithMetadata(ctx context.Context, in *GetFilesByMetadataRequest, opts ...grpc.CallOption) (*GetFilesByMetadataResponse, error) {
	out := new(GetFilesByMetadataResponse)
	err := c.cc.Invoke(ctx, "/kic.media.MediaStorage/GetFilesWithMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaStorageClient) DeleteFilesWithMetaData(ctx context.Context, in *DeleteFilesWithMetaDataRequest, opts ...grpc.CallOption) (*DeleteFilesWithMetaDataResponse, error) {
	out := new(DeleteFilesWithMetaDataResponse)
	err := c.cc.Invoke(ctx, "/kic.media.MediaStorage/DeleteFilesWithMetaData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediaStorageServer is the server API for MediaStorage service.
// All implementations must embed UnimplementedMediaStorageServer
// for forward compatibility
type MediaStorageServer interface {
	//
	// Send a file as a stream of messages, starting with a message containing a File message, then
	// followed by an arbitrary number of messages containing bytes representing the file. The response
	// will then confirm the number of bytes received or provide an error.
	UploadFile(MediaStorage_UploadFileServer) error
	//
	// Using the same format as above, the service allows the client to retrieve a stored file.
	DownloadFileByName(*DownloadFileRequest, MediaStorage_DownloadFileByNameServer) error
	// Check for the existence of a file by filename
	CheckForFileByName(context.Context, *CheckForFileRequest) (*CheckForFileResponse, error)
	//
	// Allows for the requesting of files with specific key value pairs as metadata. The strictness can be set
	// such that for example only perfect matches will be returned.
	GetFilesWithMetadata(context.Context, *GetFilesByMetadataRequest) (*GetFilesByMetadataResponse, error)
	//
	// Using the same strictness settings as the above, delete particular files with certain metadata.
	DeleteFilesWithMetaData(context.Context, *DeleteFilesWithMetaDataRequest) (*DeleteFilesWithMetaDataResponse, error)
	mustEmbedUnimplementedMediaStorageServer()
}

// UnimplementedMediaStorageServer must be embedded to have forward compatible implementations.
type UnimplementedMediaStorageServer struct{}

func (UnimplementedMediaStorageServer) UploadFile(MediaStorage_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}

func (UnimplementedMediaStorageServer) DownloadFileByName(*DownloadFileRequest, MediaStorage_DownloadFileByNameServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFileByName not implemented")
}

func (UnimplementedMediaStorageServer) CheckForFileByName(context.Context, *CheckForFileRequest) (*CheckForFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckForFileByName not implemented")
}

func (UnimplementedMediaStorageServer) GetFilesWithMetadata(context.Context, *GetFilesByMetadataRequest) (*GetFilesByMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilesWithMetadata not implemented")
}

func (UnimplementedMediaStorageServer) DeleteFilesWithMetaData(context.Context, *DeleteFilesWithMetaDataRequest) (*DeleteFilesWithMetaDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFilesWithMetaData not implemented")
}
func (UnimplementedMediaStorageServer) mustEmbedUnimplementedMediaStorageServer() {}

// UnsafeMediaStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MediaStorageServer will
// result in compilation errors.
type UnsafeMediaStorageServer interface {
	mustEmbedUnimplementedMediaStorageServer()
}

func RegisterMediaStorageServer(s grpc.ServiceRegistrar, srv MediaStorageServer) {
	s.RegisterService(&_MediaStorage_serviceDesc, srv)
}

func _MediaStorage_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MediaStorageServer).UploadFile(&mediaStorageUploadFileServer{stream})
}

type MediaStorage_UploadFileServer interface {
	SendAndClose(*UploadFileResponse) error
	Recv() (*UploadFileRequest, error)
	grpc.ServerStream
}

type mediaStorageUploadFileServer struct {
	grpc.ServerStream
}

func (x *mediaStorageUploadFileServer) SendAndClose(m *UploadFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mediaStorageUploadFileServer) Recv() (*UploadFileRequest, error) {
	m := new(UploadFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MediaStorage_DownloadFileByName_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MediaStorageServer).DownloadFileByName(m, &mediaStorageDownloadFileByNameServer{stream})
}

type MediaStorage_DownloadFileByNameServer interface {
	Send(*DownloadFileResponse) error
	grpc.ServerStream
}

type mediaStorageDownloadFileByNameServer struct {
	grpc.ServerStream
}

func (x *mediaStorageDownloadFileByNameServer) Send(m *DownloadFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MediaStorage_CheckForFileByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckForFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaStorageServer).CheckForFileByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kic.media.MediaStorage/CheckForFileByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaStorageServer).CheckForFileByName(ctx, req.(*CheckForFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaStorage_GetFilesWithMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilesByMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaStorageServer).GetFilesWithMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kic.media.MediaStorage/GetFilesWithMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaStorageServer).GetFilesWithMetadata(ctx, req.(*GetFilesByMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaStorage_DeleteFilesWithMetaData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFilesWithMetaDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaStorageServer).DeleteFilesWithMetaData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kic.media.MediaStorage/DeleteFilesWithMetaData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaStorageServer).DeleteFilesWithMetaData(ctx, req.(*DeleteFilesWithMetaDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MediaStorage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kic.media.MediaStorage",
	HandlerType: (*MediaStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckForFileByName",
			Handler:    _MediaStorage_CheckForFileByName_Handler,
		},
		{
			MethodName: "GetFilesWithMetadata",
			Handler:    _MediaStorage_GetFilesWithMetadata_Handler,
		},
		{
			MethodName: "DeleteFilesWithMetaData",
			Handler:    _MediaStorage_DeleteFilesWithMetaData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _MediaStorage_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFileByName",
			Handler:       _MediaStorage_DownloadFileByName_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/media.proto",
}
