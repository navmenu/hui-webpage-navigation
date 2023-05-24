// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: webnavigation/navi.proto

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
	Navi_CreateNavi_FullMethodName       = "/api.webnavigation.Navi/CreateNavi"
	Navi_DeleteNavi_FullMethodName       = "/api.webnavigation.Navi/DeleteNavi"
	Navi_SortNavi_FullMethodName         = "/api.webnavigation.Navi/SortNavi"
	Navi_ListNavi_FullMethodName         = "/api.webnavigation.Navi/ListNavi"
	Navi_GetGuestSettings_FullMethodName = "/api.webnavigation.Navi/GetGuestSettings"
	Navi_SetNotRemindInfo_FullMethodName = "/api.webnavigation.Navi/SetNotRemindInfo"
)

// NaviClient is the client API for Navi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NaviClient interface {
	// 分类添加
	CreateNavi(ctx context.Context, in *CreateNaviRequest, opts ...grpc.CallOption) (*CreateNaviReply, error)
	// 分类删除
	DeleteNavi(ctx context.Context, in *DeleteNaviRequest, opts ...grpc.CallOption) (*DeleteNaviReply, error)
	// 分类排序，策略是把新的顺序完整的发过来，假如需要其它策略可以再改代码
	SortNavi(ctx context.Context, in *SortNaviRequest, opts ...grpc.CallOption) (*SortNaviReply, error)
	// 分类列表，目前看来获得分类列表的时候必然要获得内容，因此一起返回，假如不需要可以再改
	ListNavi(ctx context.Context, in *ListNaviRequest, opts ...grpc.CallOption) (*ListNaviReply, error)
	// 主要是看"同一个cookie和IP"是否"勾选下次不再提醒"
	// 我理解的，首次登陆也可以有多次，即每次关掉网页再打开即可视为首次登陆
	// 即使没有设置，也会返回一个默认值的结果，比如false和0等
	// 需要在请求里带cookie
	GetGuestSettings(ctx context.Context, in *GetGuestSettingsRequest, opts ...grpc.CallOption) (*GetGuestSettingsReply, error)
	// 记录"同一个cookie和IP"已经"勾选下次不再提醒"-也可记录在redis里
	// 需要在请求里带cookie
	SetNotRemindInfo(ctx context.Context, in *SetNotRemindInfoRequest, opts ...grpc.CallOption) (*SetNotRemindInfoReply, error)
}

type naviClient struct {
	cc grpc.ClientConnInterface
}

func NewNaviClient(cc grpc.ClientConnInterface) NaviClient {
	return &naviClient{cc}
}

func (c *naviClient) CreateNavi(ctx context.Context, in *CreateNaviRequest, opts ...grpc.CallOption) (*CreateNaviReply, error) {
	out := new(CreateNaviReply)
	err := c.cc.Invoke(ctx, Navi_CreateNavi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *naviClient) DeleteNavi(ctx context.Context, in *DeleteNaviRequest, opts ...grpc.CallOption) (*DeleteNaviReply, error) {
	out := new(DeleteNaviReply)
	err := c.cc.Invoke(ctx, Navi_DeleteNavi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *naviClient) SortNavi(ctx context.Context, in *SortNaviRequest, opts ...grpc.CallOption) (*SortNaviReply, error) {
	out := new(SortNaviReply)
	err := c.cc.Invoke(ctx, Navi_SortNavi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *naviClient) ListNavi(ctx context.Context, in *ListNaviRequest, opts ...grpc.CallOption) (*ListNaviReply, error) {
	out := new(ListNaviReply)
	err := c.cc.Invoke(ctx, Navi_ListNavi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *naviClient) GetGuestSettings(ctx context.Context, in *GetGuestSettingsRequest, opts ...grpc.CallOption) (*GetGuestSettingsReply, error) {
	out := new(GetGuestSettingsReply)
	err := c.cc.Invoke(ctx, Navi_GetGuestSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *naviClient) SetNotRemindInfo(ctx context.Context, in *SetNotRemindInfoRequest, opts ...grpc.CallOption) (*SetNotRemindInfoReply, error) {
	out := new(SetNotRemindInfoReply)
	err := c.cc.Invoke(ctx, Navi_SetNotRemindInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NaviServer is the server API for Navi service.
// All implementations must embed UnimplementedNaviServer
// for forward compatibility
type NaviServer interface {
	// 分类添加
	CreateNavi(context.Context, *CreateNaviRequest) (*CreateNaviReply, error)
	// 分类删除
	DeleteNavi(context.Context, *DeleteNaviRequest) (*DeleteNaviReply, error)
	// 分类排序，策略是把新的顺序完整的发过来，假如需要其它策略可以再改代码
	SortNavi(context.Context, *SortNaviRequest) (*SortNaviReply, error)
	// 分类列表，目前看来获得分类列表的时候必然要获得内容，因此一起返回，假如不需要可以再改
	ListNavi(context.Context, *ListNaviRequest) (*ListNaviReply, error)
	// 主要是看"同一个cookie和IP"是否"勾选下次不再提醒"
	// 我理解的，首次登陆也可以有多次，即每次关掉网页再打开即可视为首次登陆
	// 即使没有设置，也会返回一个默认值的结果，比如false和0等
	// 需要在请求里带cookie
	GetGuestSettings(context.Context, *GetGuestSettingsRequest) (*GetGuestSettingsReply, error)
	// 记录"同一个cookie和IP"已经"勾选下次不再提醒"-也可记录在redis里
	// 需要在请求里带cookie
	SetNotRemindInfo(context.Context, *SetNotRemindInfoRequest) (*SetNotRemindInfoReply, error)
	mustEmbedUnimplementedNaviServer()
}

// UnimplementedNaviServer must be embedded to have forward compatible implementations.
type UnimplementedNaviServer struct {
}

func (UnimplementedNaviServer) CreateNavi(context.Context, *CreateNaviRequest) (*CreateNaviReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNavi not implemented")
}
func (UnimplementedNaviServer) DeleteNavi(context.Context, *DeleteNaviRequest) (*DeleteNaviReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNavi not implemented")
}
func (UnimplementedNaviServer) SortNavi(context.Context, *SortNaviRequest) (*SortNaviReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SortNavi not implemented")
}
func (UnimplementedNaviServer) ListNavi(context.Context, *ListNaviRequest) (*ListNaviReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNavi not implemented")
}
func (UnimplementedNaviServer) GetGuestSettings(context.Context, *GetGuestSettingsRequest) (*GetGuestSettingsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuestSettings not implemented")
}
func (UnimplementedNaviServer) SetNotRemindInfo(context.Context, *SetNotRemindInfoRequest) (*SetNotRemindInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetNotRemindInfo not implemented")
}
func (UnimplementedNaviServer) mustEmbedUnimplementedNaviServer() {}

// UnsafeNaviServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NaviServer will
// result in compilation errors.
type UnsafeNaviServer interface {
	mustEmbedUnimplementedNaviServer()
}

func RegisterNaviServer(s grpc.ServiceRegistrar, srv NaviServer) {
	s.RegisterService(&Navi_ServiceDesc, srv)
}

func _Navi_CreateNavi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNaviRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NaviServer).CreateNavi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Navi_CreateNavi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NaviServer).CreateNavi(ctx, req.(*CreateNaviRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Navi_DeleteNavi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNaviRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NaviServer).DeleteNavi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Navi_DeleteNavi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NaviServer).DeleteNavi(ctx, req.(*DeleteNaviRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Navi_SortNavi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SortNaviRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NaviServer).SortNavi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Navi_SortNavi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NaviServer).SortNavi(ctx, req.(*SortNaviRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Navi_ListNavi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNaviRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NaviServer).ListNavi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Navi_ListNavi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NaviServer).ListNavi(ctx, req.(*ListNaviRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Navi_GetGuestSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGuestSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NaviServer).GetGuestSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Navi_GetGuestSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NaviServer).GetGuestSettings(ctx, req.(*GetGuestSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Navi_SetNotRemindInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetNotRemindInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NaviServer).SetNotRemindInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Navi_SetNotRemindInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NaviServer).SetNotRemindInfo(ctx, req.(*SetNotRemindInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Navi_ServiceDesc is the grpc.ServiceDesc for Navi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Navi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.webnavigation.Navi",
	HandlerType: (*NaviServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNavi",
			Handler:    _Navi_CreateNavi_Handler,
		},
		{
			MethodName: "DeleteNavi",
			Handler:    _Navi_DeleteNavi_Handler,
		},
		{
			MethodName: "SortNavi",
			Handler:    _Navi_SortNavi_Handler,
		},
		{
			MethodName: "ListNavi",
			Handler:    _Navi_ListNavi_Handler,
		},
		{
			MethodName: "GetGuestSettings",
			Handler:    _Navi_GetGuestSettings_Handler,
		},
		{
			MethodName: "SetNotRemindInfo",
			Handler:    _Navi_SetNotRemindInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "webnavigation/navi.proto",
}
