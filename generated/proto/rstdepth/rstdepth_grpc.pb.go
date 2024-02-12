// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: protocols/rstdepth.proto

package rstdepth

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
	DepthResponseService_StreamDepthResponse_FullMethodName = "/depthresponse.DepthResponseService/StreamDepthResponse"
)

// DepthResponseServiceClient is the client API for DepthResponseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DepthResponseServiceClient interface {
	StreamDepthResponse(ctx context.Context, opts ...grpc.CallOption) (DepthResponseService_StreamDepthResponseClient, error)
}

type depthResponseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDepthResponseServiceClient(cc grpc.ClientConnInterface) DepthResponseServiceClient {
	return &depthResponseServiceClient{cc}
}

func (c *depthResponseServiceClient) StreamDepthResponse(ctx context.Context, opts ...grpc.CallOption) (DepthResponseService_StreamDepthResponseClient, error) {
	stream, err := c.cc.NewStream(ctx, &DepthResponseService_ServiceDesc.Streams[0], DepthResponseService_StreamDepthResponse_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &depthResponseServiceStreamDepthResponseClient{stream}
	return x, nil
}

type DepthResponseService_StreamDepthResponseClient interface {
	Send(*DepthResponse) error
	CloseAndRecv() (*StreamDepthResponseReply, error)
	grpc.ClientStream
}

type depthResponseServiceStreamDepthResponseClient struct {
	grpc.ClientStream
}

func (x *depthResponseServiceStreamDepthResponseClient) Send(m *DepthResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *depthResponseServiceStreamDepthResponseClient) CloseAndRecv() (*StreamDepthResponseReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamDepthResponseReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DepthResponseServiceServer is the server API for DepthResponseService service.
// All implementations must embed UnimplementedDepthResponseServiceServer
// for forward compatibility
type DepthResponseServiceServer interface {
	StreamDepthResponse(DepthResponseService_StreamDepthResponseServer) error
	mustEmbedUnimplementedDepthResponseServiceServer()
}

// UnimplementedDepthResponseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDepthResponseServiceServer struct {
}

func (UnimplementedDepthResponseServiceServer) StreamDepthResponse(DepthResponseService_StreamDepthResponseServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamDepthResponse not implemented")
}
func (UnimplementedDepthResponseServiceServer) mustEmbedUnimplementedDepthResponseServiceServer() {}

// UnsafeDepthResponseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DepthResponseServiceServer will
// result in compilation errors.
type UnsafeDepthResponseServiceServer interface {
	mustEmbedUnimplementedDepthResponseServiceServer()
}

func RegisterDepthResponseServiceServer(s grpc.ServiceRegistrar, srv DepthResponseServiceServer) {
	s.RegisterService(&DepthResponseService_ServiceDesc, srv)
}

func _DepthResponseService_StreamDepthResponse_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DepthResponseServiceServer).StreamDepthResponse(&depthResponseServiceStreamDepthResponseServer{stream})
}

type DepthResponseService_StreamDepthResponseServer interface {
	SendAndClose(*StreamDepthResponseReply) error
	Recv() (*DepthResponse, error)
	grpc.ServerStream
}

type depthResponseServiceStreamDepthResponseServer struct {
	grpc.ServerStream
}

func (x *depthResponseServiceStreamDepthResponseServer) SendAndClose(m *StreamDepthResponseReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *depthResponseServiceStreamDepthResponseServer) Recv() (*DepthResponse, error) {
	m := new(DepthResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DepthResponseService_ServiceDesc is the grpc.ServiceDesc for DepthResponseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DepthResponseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "depthresponse.DepthResponseService",
	HandlerType: (*DepthResponseServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamDepthResponse",
			Handler:       _DepthResponseService_StreamDepthResponse_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "protocols/rstdepth.proto",
}
