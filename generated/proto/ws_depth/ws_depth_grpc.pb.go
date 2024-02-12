// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: protocols/ws_depth.proto

package ws_depth

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
	DepthEventService_StreamDepthEvent_FullMethodName = "/depthevent.DepthEventService/StreamDepthEvent"
)

// DepthEventServiceClient is the client API for DepthEventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DepthEventServiceClient interface {
	StreamDepthEvent(ctx context.Context, opts ...grpc.CallOption) (DepthEventService_StreamDepthEventClient, error)
}

type depthEventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDepthEventServiceClient(cc grpc.ClientConnInterface) DepthEventServiceClient {
	return &depthEventServiceClient{cc}
}

func (c *depthEventServiceClient) StreamDepthEvent(ctx context.Context, opts ...grpc.CallOption) (DepthEventService_StreamDepthEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &DepthEventService_ServiceDesc.Streams[0], DepthEventService_StreamDepthEvent_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &depthEventServiceStreamDepthEventClient{stream}
	return x, nil
}

type DepthEventService_StreamDepthEventClient interface {
	Send(*WsDepthEvent) error
	CloseAndRecv() (*StreamDepthEventResponse, error)
	grpc.ClientStream
}

type depthEventServiceStreamDepthEventClient struct {
	grpc.ClientStream
}

func (x *depthEventServiceStreamDepthEventClient) Send(m *WsDepthEvent) error {
	return x.ClientStream.SendMsg(m)
}

func (x *depthEventServiceStreamDepthEventClient) CloseAndRecv() (*StreamDepthEventResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamDepthEventResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DepthEventServiceServer is the server API for DepthEventService service.
// All implementations must embed UnimplementedDepthEventServiceServer
// for forward compatibility
type DepthEventServiceServer interface {
	StreamDepthEvent(DepthEventService_StreamDepthEventServer) error
	mustEmbedUnimplementedDepthEventServiceServer()
}

// UnimplementedDepthEventServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDepthEventServiceServer struct {
}

func (UnimplementedDepthEventServiceServer) StreamDepthEvent(DepthEventService_StreamDepthEventServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamDepthEvent not implemented")
}
func (UnimplementedDepthEventServiceServer) mustEmbedUnimplementedDepthEventServiceServer() {}

// UnsafeDepthEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DepthEventServiceServer will
// result in compilation errors.
type UnsafeDepthEventServiceServer interface {
	mustEmbedUnimplementedDepthEventServiceServer()
}

func RegisterDepthEventServiceServer(s grpc.ServiceRegistrar, srv DepthEventServiceServer) {
	s.RegisterService(&DepthEventService_ServiceDesc, srv)
}

func _DepthEventService_StreamDepthEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DepthEventServiceServer).StreamDepthEvent(&depthEventServiceStreamDepthEventServer{stream})
}

type DepthEventService_StreamDepthEventServer interface {
	SendAndClose(*StreamDepthEventResponse) error
	Recv() (*WsDepthEvent, error)
	grpc.ServerStream
}

type depthEventServiceStreamDepthEventServer struct {
	grpc.ServerStream
}

func (x *depthEventServiceStreamDepthEventServer) SendAndClose(m *StreamDepthEventResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *depthEventServiceStreamDepthEventServer) Recv() (*WsDepthEvent, error) {
	m := new(WsDepthEvent)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DepthEventService_ServiceDesc is the grpc.ServiceDesc for DepthEventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DepthEventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "depthevent.DepthEventService",
	HandlerType: (*DepthEventServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamDepthEvent",
			Handler:       _DepthEventService_StreamDepthEvent_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "protocols/ws_depth.proto",
}