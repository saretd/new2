// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package trv_airport_api

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

// TrvAirportApiServiceClient is the client API for TrvAirportApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrvAirportApiServiceClient interface {
	CreateAirportV1(ctx context.Context, in *CreateAirportV1Request, opts ...grpc.CallOption) (*CreateAirportV1Response, error)
	DescribeAirportV1(ctx context.Context, in *DescribeAirportV1Request, opts ...grpc.CallOption) (*DescribeAirportV1Response, error)
	ListAirportsV1(ctx context.Context, in *ListAirportsV1Request, opts ...grpc.CallOption) (*ListAirportsV1Response, error)
	DeleteAirportV1(ctx context.Context, in *DeleteAirportV1Request, opts ...grpc.CallOption) (*DeleteAirportV1Response, error)
}

type trvAirportApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTrvAirportApiServiceClient(cc grpc.ClientConnInterface) TrvAirportApiServiceClient {
	return &trvAirportApiServiceClient{cc}
}

func (c *trvAirportApiServiceClient) CreateAirportV1(ctx context.Context, in *CreateAirportV1Request, opts ...grpc.CallOption) (*CreateAirportV1Response, error) {
	out := new(CreateAirportV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.trv_airport_api.v1.TrvAirportApiService/CreateAirportV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trvAirportApiServiceClient) DescribeAirportV1(ctx context.Context, in *DescribeAirportV1Request, opts ...grpc.CallOption) (*DescribeAirportV1Response, error) {
	out := new(DescribeAirportV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.trv_airport_api.v1.TrvAirportApiService/DescribeAirportV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trvAirportApiServiceClient) ListAirportsV1(ctx context.Context, in *ListAirportsV1Request, opts ...grpc.CallOption) (*ListAirportsV1Response, error) {
	out := new(ListAirportsV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.trv_airport_api.v1.TrvAirportApiService/ListAirportsV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trvAirportApiServiceClient) DeleteAirportV1(ctx context.Context, in *DeleteAirportV1Request, opts ...grpc.CallOption) (*DeleteAirportV1Response, error) {
	out := new(DeleteAirportV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.trv_airport_api.v1.TrvAirportApiService/DeleteAirportV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrvAirportApiServiceServer is the server API for TrvAirportApiService service.
// All implementations must embed UnimplementedTrvAirportApiServiceServer
// for forward compatibility
type TrvAirportApiServiceServer interface {
	CreateAirportV1(context.Context, *CreateAirportV1Request) (*CreateAirportV1Response, error)
	DescribeAirportV1(context.Context, *DescribeAirportV1Request) (*DescribeAirportV1Response, error)
	ListAirportsV1(context.Context, *ListAirportsV1Request) (*ListAirportsV1Response, error)
	DeleteAirportV1(context.Context, *DeleteAirportV1Request) (*DeleteAirportV1Response, error)
	mustEmbedUnimplementedTrvAirportApiServiceServer()
}

// UnimplementedTrvAirportApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTrvAirportApiServiceServer struct {
}

func (UnimplementedTrvAirportApiServiceServer) CreateAirportV1(context.Context, *CreateAirportV1Request) (*CreateAirportV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAirportV1 not implemented")
}
func (UnimplementedTrvAirportApiServiceServer) DescribeAirportV1(context.Context, *DescribeAirportV1Request) (*DescribeAirportV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeAirportV1 not implemented")
}
func (UnimplementedTrvAirportApiServiceServer) ListAirportsV1(context.Context, *ListAirportsV1Request) (*ListAirportsV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAirportsV1 not implemented")
}
func (UnimplementedTrvAirportApiServiceServer) DeleteAirportV1(context.Context, *DeleteAirportV1Request) (*DeleteAirportV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAirportV1 not implemented")
}
func (UnimplementedTrvAirportApiServiceServer) mustEmbedUnimplementedTrvAirportApiServiceServer() {}

// UnsafeTrvAirportApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrvAirportApiServiceServer will
// result in compilation errors.
type UnsafeTrvAirportApiServiceServer interface {
	mustEmbedUnimplementedTrvAirportApiServiceServer()
}

func RegisterTrvAirportApiServiceServer(s grpc.ServiceRegistrar, srv TrvAirportApiServiceServer) {
	s.RegisterService(&TrvAirportApiService_ServiceDesc, srv)
}

func _TrvAirportApiService_CreateAirportV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAirportV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrvAirportApiServiceServer).CreateAirportV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.trv_airport_api.v1.TrvAirportApiService/CreateAirportV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrvAirportApiServiceServer).CreateAirportV1(ctx, req.(*CreateAirportV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrvAirportApiService_DescribeAirportV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeAirportV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrvAirportApiServiceServer).DescribeAirportV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.trv_airport_api.v1.TrvAirportApiService/DescribeAirportV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrvAirportApiServiceServer).DescribeAirportV1(ctx, req.(*DescribeAirportV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrvAirportApiService_ListAirportsV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAirportsV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrvAirportApiServiceServer).ListAirportsV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.trv_airport_api.v1.TrvAirportApiService/ListAirportsV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrvAirportApiServiceServer).ListAirportsV1(ctx, req.(*ListAirportsV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrvAirportApiService_DeleteAirportV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAirportV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrvAirportApiServiceServer).DeleteAirportV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.trv_airport_api.v1.TrvAirportApiService/DeleteAirportV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrvAirportApiServiceServer).DeleteAirportV1(ctx, req.(*DeleteAirportV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// TrvAirportApiService_ServiceDesc is the grpc.ServiceDesc for TrvAirportApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrvAirportApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ozonmp.trv_airport_api.v1.TrvAirportApiService",
	HandlerType: (*TrvAirportApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAirportV1",
			Handler:    _TrvAirportApiService_CreateAirportV1_Handler,
		},
		{
			MethodName: "DescribeAirportV1",
			Handler:    _TrvAirportApiService_DescribeAirportV1_Handler,
		},
		{
			MethodName: "ListAirportsV1",
			Handler:    _TrvAirportApiService_ListAirportsV1_Handler,
		},
		{
			MethodName: "DeleteAirportV1",
			Handler:    _TrvAirportApiService_DeleteAirportV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ozonmp/trv_airport_api/v1/trv_airport_api.proto",
}
