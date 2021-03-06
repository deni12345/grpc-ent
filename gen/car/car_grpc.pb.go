// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: car.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CarServiceClient is the client API for CarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CarServiceClient interface {
	GetCarInfo(ctx context.Context, in *GetCarInfoRequest, opts ...grpc.CallOption) (*GetCarInfoResponse, error)
	StreamCarsInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (CarService_StreamCarsInfoClient, error)
	CreateCarInfo(ctx context.Context, in *CreateCarInfoRequest, opts ...grpc.CallOption) (*CreateCarInfoResponse, error)
	GetCarsInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetCarsInfoResponse, error)
}

type carServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCarServiceClient(cc grpc.ClientConnInterface) CarServiceClient {
	return &carServiceClient{cc}
}

func (c *carServiceClient) GetCarInfo(ctx context.Context, in *GetCarInfoRequest, opts ...grpc.CallOption) (*GetCarInfoResponse, error) {
	out := new(GetCarInfoResponse)
	err := c.cc.Invoke(ctx, "/rpc.car.v1.CarService/GetCarInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carServiceClient) StreamCarsInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (CarService_StreamCarsInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &CarService_ServiceDesc.Streams[0], "/rpc.car.v1.CarService/StreamCarsInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &carServiceStreamCarsInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CarService_StreamCarsInfoClient interface {
	Recv() (*Car, error)
	grpc.ClientStream
}

type carServiceStreamCarsInfoClient struct {
	grpc.ClientStream
}

func (x *carServiceStreamCarsInfoClient) Recv() (*Car, error) {
	m := new(Car)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *carServiceClient) CreateCarInfo(ctx context.Context, in *CreateCarInfoRequest, opts ...grpc.CallOption) (*CreateCarInfoResponse, error) {
	out := new(CreateCarInfoResponse)
	err := c.cc.Invoke(ctx, "/rpc.car.v1.CarService/CreateCarInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carServiceClient) GetCarsInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetCarsInfoResponse, error) {
	out := new(GetCarsInfoResponse)
	err := c.cc.Invoke(ctx, "/rpc.car.v1.CarService/GetCarsInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarServiceServer is the server API for CarService service.
// All implementations must embed UnimplementedCarServiceServer
// for forward compatibility
type CarServiceServer interface {
	GetCarInfo(context.Context, *GetCarInfoRequest) (*GetCarInfoResponse, error)
	StreamCarsInfo(*emptypb.Empty, CarService_StreamCarsInfoServer) error
	CreateCarInfo(context.Context, *CreateCarInfoRequest) (*CreateCarInfoResponse, error)
	GetCarsInfo(context.Context, *emptypb.Empty) (*GetCarsInfoResponse, error)
	mustEmbedUnimplementedCarServiceServer()
}

// UnimplementedCarServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCarServiceServer struct {
}

func (UnimplementedCarServiceServer) GetCarInfo(context.Context, *GetCarInfoRequest) (*GetCarInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCarInfo not implemented")
}
func (UnimplementedCarServiceServer) StreamCarsInfo(*emptypb.Empty, CarService_StreamCarsInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamCarsInfo not implemented")
}
func (UnimplementedCarServiceServer) CreateCarInfo(context.Context, *CreateCarInfoRequest) (*CreateCarInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCarInfo not implemented")
}
func (UnimplementedCarServiceServer) GetCarsInfo(context.Context, *emptypb.Empty) (*GetCarsInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCarsInfo not implemented")
}
func (UnimplementedCarServiceServer) mustEmbedUnimplementedCarServiceServer() {}

// UnsafeCarServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CarServiceServer will
// result in compilation errors.
type UnsafeCarServiceServer interface {
	mustEmbedUnimplementedCarServiceServer()
}

func RegisterCarServiceServer(s grpc.ServiceRegistrar, srv CarServiceServer) {
	s.RegisterService(&CarService_ServiceDesc, srv)
}

func _CarService_GetCarInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCarInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServiceServer).GetCarInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.car.v1.CarService/GetCarInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServiceServer).GetCarInfo(ctx, req.(*GetCarInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarService_StreamCarsInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CarServiceServer).StreamCarsInfo(m, &carServiceStreamCarsInfoServer{stream})
}

type CarService_StreamCarsInfoServer interface {
	Send(*Car) error
	grpc.ServerStream
}

type carServiceStreamCarsInfoServer struct {
	grpc.ServerStream
}

func (x *carServiceStreamCarsInfoServer) Send(m *Car) error {
	return x.ServerStream.SendMsg(m)
}

func _CarService_CreateCarInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCarInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServiceServer).CreateCarInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.car.v1.CarService/CreateCarInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServiceServer).CreateCarInfo(ctx, req.(*CreateCarInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarService_GetCarsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServiceServer).GetCarsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.car.v1.CarService/GetCarsInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServiceServer).GetCarsInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CarService_ServiceDesc is the grpc.ServiceDesc for CarService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CarService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.car.v1.CarService",
	HandlerType: (*CarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCarInfo",
			Handler:    _CarService_GetCarInfo_Handler,
		},
		{
			MethodName: "CreateCarInfo",
			Handler:    _CarService_CreateCarInfo_Handler,
		},
		{
			MethodName: "GetCarsInfo",
			Handler:    _CarService_GetCarsInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamCarsInfo",
			Handler:       _CarService_StreamCarsInfo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "car.proto",
}
