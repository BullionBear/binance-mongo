// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: protocols/wspdepth.proto

package wspdepth

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
	PartialDepthEventService_StreamPartialDepthEvent_FullMethodName = "/pdepthevent.PartialDepthEventService/StreamPartialDepthEvent"
)

// PartialDepthEventServiceClient is the client API for PartialDepthEventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PartialDepthEventServiceClient interface {
	StreamPartialDepthEvent(ctx context.Context, opts ...grpc.CallOption) (PartialDepthEventService_StreamPartialDepthEventClient, error)
}

type partialDepthEventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPartialDepthEventServiceClient(cc grpc.ClientConnInterface) PartialDepthEventServiceClient {
	return &partialDepthEventServiceClient{cc}
}

func (c *partialDepthEventServiceClient) StreamPartialDepthEvent(ctx context.Context, opts ...grpc.CallOption) (PartialDepthEventService_StreamPartialDepthEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &PartialDepthEventService_ServiceDesc.Streams[0], PartialDepthEventService_StreamPartialDepthEvent_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &partialDepthEventServiceStreamPartialDepthEventClient{stream}
	return x, nil
}

type PartialDepthEventService_StreamPartialDepthEventClient interface {
	Send(*WsPartialDepthEvent) error
	CloseAndRecv() (*StreamDepthEventResponse, error)
	grpc.ClientStream
}

type partialDepthEventServiceStreamPartialDepthEventClient struct {
	grpc.ClientStream
}

func (x *partialDepthEventServiceStreamPartialDepthEventClient) Send(m *WsPartialDepthEvent) error {
	return x.ClientStream.SendMsg(m)
}

func (x *partialDepthEventServiceStreamPartialDepthEventClient) CloseAndRecv() (*StreamDepthEventResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamDepthEventResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PartialDepthEventServiceServer is the server API for PartialDepthEventService service.
// All implementations must embed UnimplementedPartialDepthEventServiceServer
// for forward compatibility
type PartialDepthEventServiceServer interface {
	StreamPartialDepthEvent(PartialDepthEventService_StreamPartialDepthEventServer) error
	mustEmbedUnimplementedPartialDepthEventServiceServer()
}

// UnimplementedPartialDepthEventServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPartialDepthEventServiceServer struct {
}

func (UnimplementedPartialDepthEventServiceServer) StreamPartialDepthEvent(PartialDepthEventService_StreamPartialDepthEventServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamPartialDepthEvent not implemented")
}
func (UnimplementedPartialDepthEventServiceServer) mustEmbedUnimplementedPartialDepthEventServiceServer() {
}

// UnsafePartialDepthEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PartialDepthEventServiceServer will
// result in compilation errors.
type UnsafePartialDepthEventServiceServer interface {
	mustEmbedUnimplementedPartialDepthEventServiceServer()
}

func RegisterPartialDepthEventServiceServer(s grpc.ServiceRegistrar, srv PartialDepthEventServiceServer) {
	s.RegisterService(&PartialDepthEventService_ServiceDesc, srv)
}

func _PartialDepthEventService_StreamPartialDepthEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PartialDepthEventServiceServer).StreamPartialDepthEvent(&partialDepthEventServiceStreamPartialDepthEventServer{stream})
}

type PartialDepthEventService_StreamPartialDepthEventServer interface {
	SendAndClose(*StreamDepthEventResponse) error
	Recv() (*WsPartialDepthEvent, error)
	grpc.ServerStream
}

type partialDepthEventServiceStreamPartialDepthEventServer struct {
	grpc.ServerStream
}

func (x *partialDepthEventServiceStreamPartialDepthEventServer) SendAndClose(m *StreamDepthEventResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *partialDepthEventServiceStreamPartialDepthEventServer) Recv() (*WsPartialDepthEvent, error) {
	m := new(WsPartialDepthEvent)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PartialDepthEventService_ServiceDesc is the grpc.ServiceDesc for PartialDepthEventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PartialDepthEventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pdepthevent.PartialDepthEventService",
	HandlerType: (*PartialDepthEventServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamPartialDepthEvent",
			Handler:       _PartialDepthEventService_StreamPartialDepthEvent_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "protocols/wspdepth.proto",
}
