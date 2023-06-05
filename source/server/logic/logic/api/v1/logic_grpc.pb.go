// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: v1/logic.proto

package logic

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Logic_GetConnectorUrl_FullMethodName = "/api.v1.logic.logic.Logic/GetConnectorUrl"
	Logic_GetToolOptions_FullMethodName  = "/api.v1.logic.logic.Logic/GetToolOptions"
	Logic_GetPreEmojis_FullMethodName    = "/api.v1.logic.logic.Logic/GetPreEmojis"
	Logic_UploadFile_FullMethodName      = "/api.v1.logic.logic.Logic/UploadFile"
	Logic_DownloadFile_FullMethodName    = "/api.v1.logic.logic.Logic/DownloadFile"
)

// LogicClient is the client API for Logic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogicClient interface {
	GetConnectorUrl(ctx context.Context, in *GetConnectorUrlRequest, opts ...grpc.CallOption) (*GetConnectorUrlResponse, error)
	GetToolOptions(ctx context.Context, in *GetToolOptionsRequest, opts ...grpc.CallOption) (*GetToolOptionsResponse, error)
	GetPreEmojis(ctx context.Context, in *GetPreEmojisRequest, opts ...grpc.CallOption) (*GetPreEmojisResponse, error)
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (Logic_UploadFileClient, error)
	DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (Logic_DownloadFileClient, error)
}

type logicClient struct {
	cc grpc.ClientConnInterface
}

func NewLogicClient(cc grpc.ClientConnInterface) LogicClient {
	return &logicClient{cc}
}

func (c *logicClient) GetConnectorUrl(ctx context.Context, in *GetConnectorUrlRequest, opts ...grpc.CallOption) (*GetConnectorUrlResponse, error) {
	out := new(GetConnectorUrlResponse)
	err := c.cc.Invoke(ctx, Logic_GetConnectorUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) GetToolOptions(ctx context.Context, in *GetToolOptionsRequest, opts ...grpc.CallOption) (*GetToolOptionsResponse, error) {
	out := new(GetToolOptionsResponse)
	err := c.cc.Invoke(ctx, Logic_GetToolOptions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) GetPreEmojis(ctx context.Context, in *GetPreEmojisRequest, opts ...grpc.CallOption) (*GetPreEmojisResponse, error) {
	out := new(GetPreEmojisResponse)
	err := c.cc.Invoke(ctx, Logic_GetPreEmojis_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (Logic_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &Logic_ServiceDesc.Streams[0], Logic_UploadFile_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &logicUploadFileClient{stream}
	return x, nil
}

type Logic_UploadFileClient interface {
	Send(*UploadFileRequest) error
	CloseAndRecv() (*UploadFileResponse, error)
	grpc.ClientStream
}

type logicUploadFileClient struct {
	grpc.ClientStream
}

func (x *logicUploadFileClient) Send(m *UploadFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *logicUploadFileClient) CloseAndRecv() (*UploadFileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logicClient) DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (Logic_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &Logic_ServiceDesc.Streams[1], Logic_DownloadFile_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &logicDownloadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Logic_DownloadFileClient interface {
	Recv() (*DownloadFileResponse, error)
	grpc.ClientStream
}

type logicDownloadFileClient struct {
	grpc.ClientStream
}

func (x *logicDownloadFileClient) Recv() (*DownloadFileResponse, error) {
	m := new(DownloadFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogicServer is the server API for Logic service.
// All implementations must embed UnimplementedLogicServer
// for forward compatibility
type LogicServer interface {
	GetConnectorUrl(context.Context, *GetConnectorUrlRequest) (*GetConnectorUrlResponse, error)
	GetToolOptions(context.Context, *GetToolOptionsRequest) (*GetToolOptionsResponse, error)
	GetPreEmojis(context.Context, *GetPreEmojisRequest) (*GetPreEmojisResponse, error)
	UploadFile(Logic_UploadFileServer) error
	DownloadFile(*DownloadFileRequest, Logic_DownloadFileServer) error
	mustEmbedUnimplementedLogicServer()
}

// UnimplementedLogicServer must be embedded to have forward compatible implementations.
type UnimplementedLogicServer struct {
}

func (UnimplementedLogicServer) GetConnectorUrl(context.Context, *GetConnectorUrlRequest) (*GetConnectorUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectorUrl not implemented")
}
func (UnimplementedLogicServer) GetToolOptions(context.Context, *GetToolOptionsRequest) (*GetToolOptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToolOptions not implemented")
}
func (UnimplementedLogicServer) GetPreEmojis(context.Context, *GetPreEmojisRequest) (*GetPreEmojisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPreEmojis not implemented")
}
func (UnimplementedLogicServer) UploadFile(Logic_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedLogicServer) DownloadFile(*DownloadFileRequest, Logic_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedLogicServer) mustEmbedUnimplementedLogicServer() {}

// UnsafeLogicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogicServer will
// result in compilation errors.
type UnsafeLogicServer interface {
	mustEmbedUnimplementedLogicServer()
}

func RegisterLogicServer(s grpc.ServiceRegistrar, srv LogicServer) {
	s.RegisterService(&Logic_ServiceDesc, srv)
}

func _Logic_GetConnectorUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConnectorUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).GetConnectorUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Logic_GetConnectorUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).GetConnectorUrl(ctx, req.(*GetConnectorUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_GetToolOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetToolOptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).GetToolOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Logic_GetToolOptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).GetToolOptions(ctx, req.(*GetToolOptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_GetPreEmojis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPreEmojisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicServer).GetPreEmojis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Logic_GetPreEmojis_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicServer).GetPreEmojis(ctx, req.(*GetPreEmojisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logic_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LogicServer).UploadFile(&logicUploadFileServer{stream})
}

type Logic_UploadFileServer interface {
	SendAndClose(*UploadFileResponse) error
	Recv() (*UploadFileRequest, error)
	grpc.ServerStream
}

type logicUploadFileServer struct {
	grpc.ServerStream
}

func (x *logicUploadFileServer) SendAndClose(m *UploadFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *logicUploadFileServer) Recv() (*UploadFileRequest, error) {
	m := new(UploadFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Logic_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogicServer).DownloadFile(m, &logicDownloadFileServer{stream})
}

type Logic_DownloadFileServer interface {
	Send(*DownloadFileResponse) error
	grpc.ServerStream
}

type logicDownloadFileServer struct {
	grpc.ServerStream
}

func (x *logicDownloadFileServer) Send(m *DownloadFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Logic_ServiceDesc is the grpc.ServiceDesc for Logic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Logic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.logic.logic.Logic",
	HandlerType: (*LogicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConnectorUrl",
			Handler:    _Logic_GetConnectorUrl_Handler,
		},
		{
			MethodName: "GetToolOptions",
			Handler:    _Logic_GetToolOptions_Handler,
		},
		{
			MethodName: "GetPreEmojis",
			Handler:    _Logic_GetPreEmojis_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _Logic_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _Logic_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v1/logic.proto",
}
