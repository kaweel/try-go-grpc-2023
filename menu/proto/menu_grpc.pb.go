// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: menu/proto/menu.proto

package menu

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

const (
	Menu_ListBeers_FullMethodName = "/menu.Menu/ListBeers"
)

// MenuClient is the client API for Menu service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MenuClient interface {
	// List a beers
	ListBeers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListBeersResponse, error)
}

type menuClient struct {
	cc grpc.ClientConnInterface
}

func NewMenuClient(cc grpc.ClientConnInterface) MenuClient {
	return &menuClient{cc}
}

func (c *menuClient) ListBeers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListBeersResponse, error) {
	out := new(ListBeersResponse)
	err := c.cc.Invoke(ctx, Menu_ListBeers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MenuServer is the server API for Menu service.
// All implementations must embed UnimplementedMenuServer
// for forward compatibility
type MenuServer interface {
	// List a beers
	ListBeers(context.Context, *emptypb.Empty) (*ListBeersResponse, error)
	mustEmbedUnimplementedMenuServer()
}

// UnimplementedMenuServer must be embedded to have forward compatible implementations.
type UnimplementedMenuServer struct {
}

func (UnimplementedMenuServer) ListBeers(context.Context, *emptypb.Empty) (*ListBeersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBeers not implemented")
}
func (UnimplementedMenuServer) mustEmbedUnimplementedMenuServer() {}

// UnsafeMenuServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MenuServer will
// result in compilation errors.
type UnsafeMenuServer interface {
	mustEmbedUnimplementedMenuServer()
}

func RegisterMenuServer(s grpc.ServiceRegistrar, srv MenuServer) {
	s.RegisterService(&Menu_ServiceDesc, srv)
}

func _Menu_ListBeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServer).ListBeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Menu_ListBeers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServer).ListBeers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Menu_ServiceDesc is the grpc.ServiceDesc for Menu service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Menu_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "menu.Menu",
	HandlerType: (*MenuServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBeers",
			Handler:    _Menu_ListBeers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "menu/proto/menu.proto",
}
