// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: pg/service.proto

package pg

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
	PaymentGateway_AddClient_FullMethodName     = "/pg.PaymentGateway/AddClient"
	PaymentGateway_DeleteClient_FullMethodName  = "/pg.PaymentGateway/DeleteClient"
	PaymentGateway_HasClient_FullMethodName     = "/pg.PaymentGateway/HasClient"
	PaymentGateway_ListPayModes_FullMethodName  = "/pg.PaymentGateway/ListPayModes"
	PaymentGateway_AddPayMode_FullMethodName    = "/pg.PaymentGateway/AddPayMode"
	PaymentGateway_RemovePayMode_FullMethodName = "/pg.PaymentGateway/RemovePayMode"
	PaymentGateway_MakePayment_FullMethodName   = "/pg.PaymentGateway/MakePayment"
)

// PaymentGatewayClient is the client API for PaymentGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentGatewayClient interface {
	AddClient(ctx context.Context, in *AddClientRequest, opts ...grpc.CallOption) (*AddClientResponse, error)
	DeleteClient(ctx context.Context, in *DeleteClientRequest, opts ...grpc.CallOption) (*DeleteClientResponse, error)
	HasClient(ctx context.Context, in *HasClientRequest, opts ...grpc.CallOption) (*HasClientResponse, error)
	ListPayModes(ctx context.Context, in *ListPayModesRequest, opts ...grpc.CallOption) (*ListPayModesResponse, error)
	AddPayMode(ctx context.Context, in *AddPayModeRequest, opts ...grpc.CallOption) (*AddPayModeResponse, error)
	RemovePayMode(ctx context.Context, in *RemovePayModeRequest, opts ...grpc.CallOption) (*RemovePayModeResponse, error)
	MakePayment(ctx context.Context, in *MakePaymentRequest, opts ...grpc.CallOption) (*MakePaymentResponse, error)
}

type paymentGatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentGatewayClient(cc grpc.ClientConnInterface) PaymentGatewayClient {
	return &paymentGatewayClient{cc}
}

func (c *paymentGatewayClient) AddClient(ctx context.Context, in *AddClientRequest, opts ...grpc.CallOption) (*AddClientResponse, error) {
	out := new(AddClientResponse)
	err := c.cc.Invoke(ctx, PaymentGateway_AddClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayClient) DeleteClient(ctx context.Context, in *DeleteClientRequest, opts ...grpc.CallOption) (*DeleteClientResponse, error) {
	out := new(DeleteClientResponse)
	err := c.cc.Invoke(ctx, PaymentGateway_DeleteClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayClient) HasClient(ctx context.Context, in *HasClientRequest, opts ...grpc.CallOption) (*HasClientResponse, error) {
	out := new(HasClientResponse)
	err := c.cc.Invoke(ctx, PaymentGateway_HasClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayClient) ListPayModes(ctx context.Context, in *ListPayModesRequest, opts ...grpc.CallOption) (*ListPayModesResponse, error) {
	out := new(ListPayModesResponse)
	err := c.cc.Invoke(ctx, PaymentGateway_ListPayModes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayClient) AddPayMode(ctx context.Context, in *AddPayModeRequest, opts ...grpc.CallOption) (*AddPayModeResponse, error) {
	out := new(AddPayModeResponse)
	err := c.cc.Invoke(ctx, PaymentGateway_AddPayMode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayClient) RemovePayMode(ctx context.Context, in *RemovePayModeRequest, opts ...grpc.CallOption) (*RemovePayModeResponse, error) {
	out := new(RemovePayModeResponse)
	err := c.cc.Invoke(ctx, PaymentGateway_RemovePayMode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentGatewayClient) MakePayment(ctx context.Context, in *MakePaymentRequest, opts ...grpc.CallOption) (*MakePaymentResponse, error) {
	out := new(MakePaymentResponse)
	err := c.cc.Invoke(ctx, PaymentGateway_MakePayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentGatewayServer is the server API for PaymentGateway service.
// All implementations must embed UnimplementedPaymentGatewayServer
// for forward compatibility
type PaymentGatewayServer interface {
	AddClient(context.Context, *AddClientRequest) (*AddClientResponse, error)
	DeleteClient(context.Context, *DeleteClientRequest) (*DeleteClientResponse, error)
	HasClient(context.Context, *HasClientRequest) (*HasClientResponse, error)
	ListPayModes(context.Context, *ListPayModesRequest) (*ListPayModesResponse, error)
	AddPayMode(context.Context, *AddPayModeRequest) (*AddPayModeResponse, error)
	RemovePayMode(context.Context, *RemovePayModeRequest) (*RemovePayModeResponse, error)
	MakePayment(context.Context, *MakePaymentRequest) (*MakePaymentResponse, error)
	mustEmbedUnimplementedPaymentGatewayServer()
}

// UnimplementedPaymentGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentGatewayServer struct {
}

func (UnimplementedPaymentGatewayServer) AddClient(context.Context, *AddClientRequest) (*AddClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddClient not implemented")
}
func (UnimplementedPaymentGatewayServer) DeleteClient(context.Context, *DeleteClientRequest) (*DeleteClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClient not implemented")
}
func (UnimplementedPaymentGatewayServer) HasClient(context.Context, *HasClientRequest) (*HasClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasClient not implemented")
}
func (UnimplementedPaymentGatewayServer) ListPayModes(context.Context, *ListPayModesRequest) (*ListPayModesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPayModes not implemented")
}
func (UnimplementedPaymentGatewayServer) AddPayMode(context.Context, *AddPayModeRequest) (*AddPayModeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPayMode not implemented")
}
func (UnimplementedPaymentGatewayServer) RemovePayMode(context.Context, *RemovePayModeRequest) (*RemovePayModeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePayMode not implemented")
}
func (UnimplementedPaymentGatewayServer) MakePayment(context.Context, *MakePaymentRequest) (*MakePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakePayment not implemented")
}
func (UnimplementedPaymentGatewayServer) mustEmbedUnimplementedPaymentGatewayServer() {}

// UnsafePaymentGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentGatewayServer will
// result in compilation errors.
type UnsafePaymentGatewayServer interface {
	mustEmbedUnimplementedPaymentGatewayServer()
}

func RegisterPaymentGatewayServer(s grpc.ServiceRegistrar, srv PaymentGatewayServer) {
	s.RegisterService(&PaymentGateway_ServiceDesc, srv)
}

func _PaymentGateway_AddClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServer).AddClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentGateway_AddClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServer).AddClient(ctx, req.(*AddClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGateway_DeleteClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServer).DeleteClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentGateway_DeleteClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServer).DeleteClient(ctx, req.(*DeleteClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGateway_HasClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServer).HasClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentGateway_HasClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServer).HasClient(ctx, req.(*HasClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGateway_ListPayModes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPayModesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServer).ListPayModes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentGateway_ListPayModes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServer).ListPayModes(ctx, req.(*ListPayModesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGateway_AddPayMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPayModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServer).AddPayMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentGateway_AddPayMode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServer).AddPayMode(ctx, req.(*AddPayModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGateway_RemovePayMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePayModeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServer).RemovePayMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentGateway_RemovePayMode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServer).RemovePayMode(ctx, req.(*RemovePayModeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentGateway_MakePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentGatewayServer).MakePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentGateway_MakePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentGatewayServer).MakePayment(ctx, req.(*MakePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentGateway_ServiceDesc is the grpc.ServiceDesc for PaymentGateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentGateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pg.PaymentGateway",
	HandlerType: (*PaymentGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddClient",
			Handler:    _PaymentGateway_AddClient_Handler,
		},
		{
			MethodName: "DeleteClient",
			Handler:    _PaymentGateway_DeleteClient_Handler,
		},
		{
			MethodName: "HasClient",
			Handler:    _PaymentGateway_HasClient_Handler,
		},
		{
			MethodName: "ListPayModes",
			Handler:    _PaymentGateway_ListPayModes_Handler,
		},
		{
			MethodName: "AddPayMode",
			Handler:    _PaymentGateway_AddPayMode_Handler,
		},
		{
			MethodName: "RemovePayMode",
			Handler:    _PaymentGateway_RemovePayMode_Handler,
		},
		{
			MethodName: "MakePayment",
			Handler:    _PaymentGateway_MakePayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pg/service.proto",
}
