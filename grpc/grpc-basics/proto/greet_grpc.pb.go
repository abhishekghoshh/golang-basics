// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: greet.proto

package proto

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

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetServiceClient interface {
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	GreetServerStreaming(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (GreetService_GreetServerStreamingClient, error)
	GrretClientStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_GrretClientStreamingClient, error)
	GreetBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetBidirectionalStreamingClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) GreetServerStreaming(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (GreetService_GreetServerStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[0], "/greet.GreetService/GreetServerStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetServerStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_GreetServerStreamingClient interface {
	Recv() (*GreetResponse, error)
	grpc.ClientStream
}

type greetServiceGreetServerStreamingClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetServerStreamingClient) Recv() (*GreetResponse, error) {
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) GrretClientStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_GrretClientStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[1], "/greet.GreetService/GrretClientStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGrretClientStreamingClient{stream}
	return x, nil
}

type GreetService_GrretClientStreamingClient interface {
	Send(*GreetRequest) error
	CloseAndRecv() (*GreetResponse, error)
	grpc.ClientStream
}

type greetServiceGrretClientStreamingClient struct {
	grpc.ClientStream
}

func (x *greetServiceGrretClientStreamingClient) Send(m *GreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceGrretClientStreamingClient) CloseAndRecv() (*GreetResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) GreetBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetBidirectionalStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[2], "/greet.GreetService/GreetBidirectionalStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetBidirectionalStreamingClient{stream}
	return x, nil
}

type GreetService_GreetBidirectionalStreamingClient interface {
	Send(*GreetRequest) error
	Recv() (*GreetResponse, error)
	grpc.ClientStream
}

type greetServiceGreetBidirectionalStreamingClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetBidirectionalStreamingClient) Send(m *GreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceGreetBidirectionalStreamingClient) Recv() (*GreetResponse, error) {
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
// All implementations must embed UnimplementedGreetServiceServer
// for forward compatibility
type GreetServiceServer interface {
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	GreetServerStreaming(*GreetRequest, GreetService_GreetServerStreamingServer) error
	GrretClientStreaming(GreetService_GrretClientStreamingServer) error
	GreetBidirectionalStreaming(GreetService_GreetBidirectionalStreamingServer) error
	mustEmbedUnimplementedGreetServiceServer()
}

// UnimplementedGreetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (UnimplementedGreetServiceServer) Greet(context.Context, *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (UnimplementedGreetServiceServer) GreetServerStreaming(*GreetRequest, GreetService_GreetServerStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetServerStreaming not implemented")
}
func (UnimplementedGreetServiceServer) GrretClientStreaming(GreetService_GrretClientStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method GrretClientStreaming not implemented")
}
func (UnimplementedGreetServiceServer) GreetBidirectionalStreaming(GreetService_GreetBidirectionalStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetBidirectionalStreaming not implemented")
}
func (UnimplementedGreetServiceServer) mustEmbedUnimplementedGreetServiceServer() {}

// UnsafeGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServiceServer will
// result in compilation errors.
type UnsafeGreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
	s.RegisterService(&GreetService_ServiceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_GreetServerStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).GreetServerStreaming(m, &greetServiceGreetServerStreamingServer{stream})
}

type GreetService_GreetServerStreamingServer interface {
	Send(*GreetResponse) error
	grpc.ServerStream
}

type greetServiceGreetServerStreamingServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetServerStreamingServer) Send(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_GrretClientStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).GrretClientStreaming(&greetServiceGrretClientStreamingServer{stream})
}

type GreetService_GrretClientStreamingServer interface {
	SendAndClose(*GreetResponse) error
	Recv() (*GreetRequest, error)
	grpc.ServerStream
}

type greetServiceGrretClientStreamingServer struct {
	grpc.ServerStream
}

func (x *greetServiceGrretClientStreamingServer) SendAndClose(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceGrretClientStreamingServer) Recv() (*GreetRequest, error) {
	m := new(GreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_GreetBidirectionalStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).GreetBidirectionalStreaming(&greetServiceGreetBidirectionalStreamingServer{stream})
}

type GreetService_GreetBidirectionalStreamingServer interface {
	Send(*GreetResponse) error
	Recv() (*GreetRequest, error)
	grpc.ServerStream
}

type greetServiceGreetBidirectionalStreamingServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetBidirectionalStreamingServer) Send(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceGreetBidirectionalStreamingServer) Recv() (*GreetRequest, error) {
	m := new(GreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetService_ServiceDesc is the grpc.ServiceDesc for GreetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetServerStreaming",
			Handler:       _GreetService_GreetServerStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GrretClientStreaming",
			Handler:       _GreetService_GrretClientStreaming_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GreetBidirectionalStreaming",
			Handler:       _GreetService_GreetBidirectionalStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "greet.proto",
}
