// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: router/service.proto

package router

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
	Router_Route_FullMethodName          = "/router.Router/Route"
	Router_SetRoutingRule_FullMethodName = "/router.Router/SetRoutingRule"
)

// RouterClient is the client API for Router service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouterClient interface {
	Route(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error)
	SetRoutingRule(ctx context.Context, in *SetRoutingRuleRequest, opts ...grpc.CallOption) (*SetRoutingRuleResponse, error)
}

type routerClient struct {
	cc grpc.ClientConnInterface
}

func NewRouterClient(cc grpc.ClientConnInterface) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) Route(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error) {
	out := new(RouteResponse)
	err := c.cc.Invoke(ctx, Router_Route_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) SetRoutingRule(ctx context.Context, in *SetRoutingRuleRequest, opts ...grpc.CallOption) (*SetRoutingRuleResponse, error) {
	out := new(SetRoutingRuleResponse)
	err := c.cc.Invoke(ctx, Router_SetRoutingRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouterServer is the server API for Router service.
// All implementations must embed UnimplementedRouterServer
// for forward compatibility
type RouterServer interface {
	Route(context.Context, *RouteRequest) (*RouteResponse, error)
	SetRoutingRule(context.Context, *SetRoutingRuleRequest) (*SetRoutingRuleResponse, error)
	mustEmbedUnimplementedRouterServer()
}

// UnimplementedRouterServer must be embedded to have forward compatible implementations.
type UnimplementedRouterServer struct {
}

func (UnimplementedRouterServer) Route(context.Context, *RouteRequest) (*RouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Route not implemented")
}
func (UnimplementedRouterServer) SetRoutingRule(context.Context, *SetRoutingRuleRequest) (*SetRoutingRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRoutingRule not implemented")
}
func (UnimplementedRouterServer) mustEmbedUnimplementedRouterServer() {}

// UnsafeRouterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouterServer will
// result in compilation errors.
type UnsafeRouterServer interface {
	mustEmbedUnimplementedRouterServer()
}

func RegisterRouterServer(s grpc.ServiceRegistrar, srv RouterServer) {
	s.RegisterService(&Router_ServiceDesc, srv)
}

func _Router_Route_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Route(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Router_Route_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Route(ctx, req.(*RouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_SetRoutingRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRoutingRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).SetRoutingRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Router_SetRoutingRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).SetRoutingRule(ctx, req.(*SetRoutingRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Router_ServiceDesc is the grpc.ServiceDesc for Router service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Router_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "router.Router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Route",
			Handler:    _Router_Route_Handler,
		},
		{
			MethodName: "SetRoutingRule",
			Handler:    _Router_SetRoutingRule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "router/service.proto",
}
