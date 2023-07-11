// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: webnavigation/navi_lvl2.proto

package webnavigation

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
	NaviLvl2_CreateNaviLvl2_FullMethodName = "/api.webnavigation.NaviLvl2/CreateNaviLvl2"
	NaviLvl2_DeleteNaviLvl2_FullMethodName = "/api.webnavigation.NaviLvl2/DeleteNaviLvl2"
	NaviLvl2_SortNaviLvl2_FullMethodName   = "/api.webnavigation.NaviLvl2/SortNaviLvl2"
	NaviLvl2_ListNaviLvl2_FullMethodName   = "/api.webnavigation.NaviLvl2/ListNaviLvl2"
)

// NaviLvl2Client is the client API for NaviLvl2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NaviLvl2Client interface {
	// 分类内容添加
	CreateNaviLvl2(ctx context.Context, in *CreateNaviLvl2Request, opts ...grpc.CallOption) (*CreateNaviLvl2Reply, error)
	// 分类内容删除
	DeleteNaviLvl2(ctx context.Context, in *DeleteNaviLvl2Request, opts ...grpc.CallOption) (*DeleteNaviLvl2Reply, error)
	// 分类内容排序
	SortNaviLvl2(ctx context.Context, in *SortNaviLvl2Request, opts ...grpc.CallOption) (*SortNaviLvl2Reply, error)
	// 分类内容列表，这个接口可能是多余的，因为获得分类的时候已经顺带获取了内容
	ListNaviLvl2(ctx context.Context, in *ListNaviLvl2Request, opts ...grpc.CallOption) (*ListNaviLvl2Reply, error)
}

type naviLvl2Client struct {
	cc grpc.ClientConnInterface
}

func NewNaviLvl2Client(cc grpc.ClientConnInterface) NaviLvl2Client {
	return &naviLvl2Client{cc}
}

func (c *naviLvl2Client) CreateNaviLvl2(ctx context.Context, in *CreateNaviLvl2Request, opts ...grpc.CallOption) (*CreateNaviLvl2Reply, error) {
	out := new(CreateNaviLvl2Reply)
	err := c.cc.Invoke(ctx, NaviLvl2_CreateNaviLvl2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *naviLvl2Client) DeleteNaviLvl2(ctx context.Context, in *DeleteNaviLvl2Request, opts ...grpc.CallOption) (*DeleteNaviLvl2Reply, error) {
	out := new(DeleteNaviLvl2Reply)
	err := c.cc.Invoke(ctx, NaviLvl2_DeleteNaviLvl2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *naviLvl2Client) SortNaviLvl2(ctx context.Context, in *SortNaviLvl2Request, opts ...grpc.CallOption) (*SortNaviLvl2Reply, error) {
	out := new(SortNaviLvl2Reply)
	err := c.cc.Invoke(ctx, NaviLvl2_SortNaviLvl2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *naviLvl2Client) ListNaviLvl2(ctx context.Context, in *ListNaviLvl2Request, opts ...grpc.CallOption) (*ListNaviLvl2Reply, error) {
	out := new(ListNaviLvl2Reply)
	err := c.cc.Invoke(ctx, NaviLvl2_ListNaviLvl2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NaviLvl2Server is the server API for NaviLvl2 service.
// All implementations must embed UnimplementedNaviLvl2Server
// for forward compatibility
type NaviLvl2Server interface {
	// 分类内容添加
	CreateNaviLvl2(context.Context, *CreateNaviLvl2Request) (*CreateNaviLvl2Reply, error)
	// 分类内容删除
	DeleteNaviLvl2(context.Context, *DeleteNaviLvl2Request) (*DeleteNaviLvl2Reply, error)
	// 分类内容排序
	SortNaviLvl2(context.Context, *SortNaviLvl2Request) (*SortNaviLvl2Reply, error)
	// 分类内容列表，这个接口可能是多余的，因为获得分类的时候已经顺带获取了内容
	ListNaviLvl2(context.Context, *ListNaviLvl2Request) (*ListNaviLvl2Reply, error)
	mustEmbedUnimplementedNaviLvl2Server()
}

// UnimplementedNaviLvl2Server must be embedded to have forward compatible implementations.
type UnimplementedNaviLvl2Server struct {
}

func (UnimplementedNaviLvl2Server) CreateNaviLvl2(context.Context, *CreateNaviLvl2Request) (*CreateNaviLvl2Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNaviLvl2 not implemented")
}
func (UnimplementedNaviLvl2Server) DeleteNaviLvl2(context.Context, *DeleteNaviLvl2Request) (*DeleteNaviLvl2Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNaviLvl2 not implemented")
}
func (UnimplementedNaviLvl2Server) SortNaviLvl2(context.Context, *SortNaviLvl2Request) (*SortNaviLvl2Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SortNaviLvl2 not implemented")
}
func (UnimplementedNaviLvl2Server) ListNaviLvl2(context.Context, *ListNaviLvl2Request) (*ListNaviLvl2Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNaviLvl2 not implemented")
}
func (UnimplementedNaviLvl2Server) mustEmbedUnimplementedNaviLvl2Server() {}

// UnsafeNaviLvl2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NaviLvl2Server will
// result in compilation errors.
type UnsafeNaviLvl2Server interface {
	mustEmbedUnimplementedNaviLvl2Server()
}

func RegisterNaviLvl2Server(s grpc.ServiceRegistrar, srv NaviLvl2Server) {
	s.RegisterService(&NaviLvl2_ServiceDesc, srv)
}

func _NaviLvl2_CreateNaviLvl2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNaviLvl2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NaviLvl2Server).CreateNaviLvl2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NaviLvl2_CreateNaviLvl2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NaviLvl2Server).CreateNaviLvl2(ctx, req.(*CreateNaviLvl2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NaviLvl2_DeleteNaviLvl2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNaviLvl2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NaviLvl2Server).DeleteNaviLvl2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NaviLvl2_DeleteNaviLvl2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NaviLvl2Server).DeleteNaviLvl2(ctx, req.(*DeleteNaviLvl2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NaviLvl2_SortNaviLvl2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SortNaviLvl2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NaviLvl2Server).SortNaviLvl2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NaviLvl2_SortNaviLvl2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NaviLvl2Server).SortNaviLvl2(ctx, req.(*SortNaviLvl2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NaviLvl2_ListNaviLvl2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNaviLvl2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NaviLvl2Server).ListNaviLvl2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NaviLvl2_ListNaviLvl2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NaviLvl2Server).ListNaviLvl2(ctx, req.(*ListNaviLvl2Request))
	}
	return interceptor(ctx, in, info, handler)
}

// NaviLvl2_ServiceDesc is the grpc.ServiceDesc for NaviLvl2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NaviLvl2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.webnavigation.NaviLvl2",
	HandlerType: (*NaviLvl2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNaviLvl2",
			Handler:    _NaviLvl2_CreateNaviLvl2_Handler,
		},
		{
			MethodName: "DeleteNaviLvl2",
			Handler:    _NaviLvl2_DeleteNaviLvl2_Handler,
		},
		{
			MethodName: "SortNaviLvl2",
			Handler:    _NaviLvl2_SortNaviLvl2_Handler,
		},
		{
			MethodName: "ListNaviLvl2",
			Handler:    _NaviLvl2_ListNaviLvl2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "webnavigation/navi_lvl2.proto",
}
