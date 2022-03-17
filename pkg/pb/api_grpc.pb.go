// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api.proto

package pb

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

// PunchrServiceClient is the client API for PunchrService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PunchrServiceClient interface {
	// Register takes punchr client information and saves them to the database.
	// This should be called upon start of a client.
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// GetAddrInfo returns peer address information that should be used to attempt
	// a hole punch. Clients should call this endpoint periodically.
	GetAddrInfo(ctx context.Context, in *GetAddrInfoRequest, opts ...grpc.CallOption) (*GetAddrInfoResponse, error)
	// TrackHolePunch takes measurement data from the client and persists them in the database
	TrackHolePunch(ctx context.Context, in *TrackHolePunchRequest, opts ...grpc.CallOption) (*TrackHolePunchResponse, error)
}

type punchrServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPunchrServiceClient(cc grpc.ClientConnInterface) PunchrServiceClient {
	return &punchrServiceClient{cc}
}

func (c *punchrServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/PunchrService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *punchrServiceClient) GetAddrInfo(ctx context.Context, in *GetAddrInfoRequest, opts ...grpc.CallOption) (*GetAddrInfoResponse, error) {
	out := new(GetAddrInfoResponse)
	err := c.cc.Invoke(ctx, "/PunchrService/GetAddrInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *punchrServiceClient) TrackHolePunch(ctx context.Context, in *TrackHolePunchRequest, opts ...grpc.CallOption) (*TrackHolePunchResponse, error) {
	out := new(TrackHolePunchResponse)
	err := c.cc.Invoke(ctx, "/PunchrService/TrackHolePunch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PunchrServiceServer is the server API for PunchrService service.
// All implementations must embed UnimplementedPunchrServiceServer
// for forward compatibility
type PunchrServiceServer interface {
	// Register takes punchr client information and saves them to the database.
	// This should be called upon start of a client.
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// GetAddrInfo returns peer address information that should be used to attempt
	// a hole punch. Clients should call this endpoint periodically.
	GetAddrInfo(context.Context, *GetAddrInfoRequest) (*GetAddrInfoResponse, error)
	// TrackHolePunch takes measurement data from the client and persists them in the database
	TrackHolePunch(context.Context, *TrackHolePunchRequest) (*TrackHolePunchResponse, error)
	mustEmbedUnimplementedPunchrServiceServer()
}

// UnimplementedPunchrServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPunchrServiceServer struct {
}

func (UnimplementedPunchrServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedPunchrServiceServer) GetAddrInfo(context.Context, *GetAddrInfoRequest) (*GetAddrInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddrInfo not implemented")
}
func (UnimplementedPunchrServiceServer) TrackHolePunch(context.Context, *TrackHolePunchRequest) (*TrackHolePunchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrackHolePunch not implemented")
}
func (UnimplementedPunchrServiceServer) mustEmbedUnimplementedPunchrServiceServer() {}

// UnsafePunchrServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PunchrServiceServer will
// result in compilation errors.
type UnsafePunchrServiceServer interface {
	mustEmbedUnimplementedPunchrServiceServer()
}

func RegisterPunchrServiceServer(s grpc.ServiceRegistrar, srv PunchrServiceServer) {
	s.RegisterService(&PunchrService_ServiceDesc, srv)
}

func _PunchrService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PunchrServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PunchrService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PunchrServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PunchrService_GetAddrInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddrInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PunchrServiceServer).GetAddrInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PunchrService/GetAddrInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PunchrServiceServer).GetAddrInfo(ctx, req.(*GetAddrInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PunchrService_TrackHolePunch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackHolePunchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PunchrServiceServer).TrackHolePunch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PunchrService/TrackHolePunch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PunchrServiceServer).TrackHolePunch(ctx, req.(*TrackHolePunchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PunchrService_ServiceDesc is the grpc.ServiceDesc for PunchrService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PunchrService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PunchrService",
	HandlerType: (*PunchrServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _PunchrService_Register_Handler,
		},
		{
			MethodName: "GetAddrInfo",
			Handler:    _PunchrService_GetAddrInfo_Handler,
		},
		{
			MethodName: "TrackHolePunch",
			Handler:    _PunchrService_TrackHolePunch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
