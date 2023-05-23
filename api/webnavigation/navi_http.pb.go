// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.2
// - protoc             v3.21.12
// source: webnavigation/navi.proto

package webnavigation

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationNaviCreateNavi = "/api.webnavigation.Navi/CreateNavi"
const OperationNaviDeleteNavi = "/api.webnavigation.Navi/DeleteNavi"
const OperationNaviGetGuestSettings = "/api.webnavigation.Navi/GetGuestSettings"
const OperationNaviListNavi = "/api.webnavigation.Navi/ListNavi"
const OperationNaviSetNotRemindInfo = "/api.webnavigation.Navi/SetNotRemindInfo"
const OperationNaviSortNavi = "/api.webnavigation.Navi/SortNavi"

type NaviHTTPServer interface {
	// CreateNavi 分类添加
	CreateNavi(context.Context, *CreateNaviRequest) (*CreateNaviReply, error)
	// DeleteNavi 分类删除
	DeleteNavi(context.Context, *DeleteNaviRequest) (*DeleteNaviReply, error)
	// GetGuestSettings 主要是看"同一个cookie和IP"是否"勾选下次不再提醒"，我理解的，首次登陆也可以有多次，即每次关掉网页再打开即可视为首次登陆
	GetGuestSettings(context.Context, *GetGuestSettingsRequest) (*GetGuestSettingsReply, error)
	// ListNavi 分类列表
	ListNavi(context.Context, *ListNaviRequest) (*ListNaviReply, error)
	// SetNotRemindInfo 记录"同一个cookie和IP"已经"勾选下次不再提醒"-也可记录在redis里
	SetNotRemindInfo(context.Context, *SetNotRemindInfoRequest) (*SetNotRemindInfoReply, error)
	// SortNavi 分类排序
	SortNavi(context.Context, *SortNaviRequest) (*SortNaviReply, error)
}

func RegisterNaviHTTPServer(s *http.Server, srv NaviHTTPServer) {
	r := s.Route("/")
	r.POST("/api/web-navigation/v1/create_navi", _Navi_CreateNavi0_HTTP_Handler(srv))
	r.POST("/api/web-navigation/v1/delete_navi", _Navi_DeleteNavi0_HTTP_Handler(srv))
	r.POST("/api/web-navigation/v1/sort_navi", _Navi_SortNavi0_HTTP_Handler(srv))
	r.GET("/api/web-navigation/v1/list_navi", _Navi_ListNavi0_HTTP_Handler(srv))
	r.GET("/api/web-navigation/v1/get_guest_settings", _Navi_GetGuestSettings0_HTTP_Handler(srv))
	r.POST("/api/web-navigation/v1/set_not_remind_info", _Navi_SetNotRemindInfo0_HTTP_Handler(srv))
}

func _Navi_CreateNavi0_HTTP_Handler(srv NaviHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateNaviRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNaviCreateNavi)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateNavi(ctx, req.(*CreateNaviRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateNaviReply)
		return ctx.Result(200, reply)
	}
}

func _Navi_DeleteNavi0_HTTP_Handler(srv NaviHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteNaviRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNaviDeleteNavi)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteNavi(ctx, req.(*DeleteNaviRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteNaviReply)
		return ctx.Result(200, reply)
	}
}

func _Navi_SortNavi0_HTTP_Handler(srv NaviHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SortNaviRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNaviSortNavi)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SortNavi(ctx, req.(*SortNaviRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SortNaviReply)
		return ctx.Result(200, reply)
	}
}

func _Navi_ListNavi0_HTTP_Handler(srv NaviHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListNaviRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNaviListNavi)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListNavi(ctx, req.(*ListNaviRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListNaviReply)
		return ctx.Result(200, reply)
	}
}

func _Navi_GetGuestSettings0_HTTP_Handler(srv NaviHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetGuestSettingsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNaviGetGuestSettings)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetGuestSettings(ctx, req.(*GetGuestSettingsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetGuestSettingsReply)
		return ctx.Result(200, reply)
	}
}

func _Navi_SetNotRemindInfo0_HTTP_Handler(srv NaviHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetNotRemindInfoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNaviSetNotRemindInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetNotRemindInfo(ctx, req.(*SetNotRemindInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SetNotRemindInfoReply)
		return ctx.Result(200, reply)
	}
}

type NaviHTTPClient interface {
	CreateNavi(ctx context.Context, req *CreateNaviRequest, opts ...http.CallOption) (rsp *CreateNaviReply, err error)
	DeleteNavi(ctx context.Context, req *DeleteNaviRequest, opts ...http.CallOption) (rsp *DeleteNaviReply, err error)
	GetGuestSettings(ctx context.Context, req *GetGuestSettingsRequest, opts ...http.CallOption) (rsp *GetGuestSettingsReply, err error)
	ListNavi(ctx context.Context, req *ListNaviRequest, opts ...http.CallOption) (rsp *ListNaviReply, err error)
	SetNotRemindInfo(ctx context.Context, req *SetNotRemindInfoRequest, opts ...http.CallOption) (rsp *SetNotRemindInfoReply, err error)
	SortNavi(ctx context.Context, req *SortNaviRequest, opts ...http.CallOption) (rsp *SortNaviReply, err error)
}

type NaviHTTPClientImpl struct {
	cc *http.Client
}

func NewNaviHTTPClient(client *http.Client) NaviHTTPClient {
	return &NaviHTTPClientImpl{client}
}

func (c *NaviHTTPClientImpl) CreateNavi(ctx context.Context, in *CreateNaviRequest, opts ...http.CallOption) (*CreateNaviReply, error) {
	var out CreateNaviReply
	pattern := "/api/web-navigation/v1/create_navi"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationNaviCreateNavi))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NaviHTTPClientImpl) DeleteNavi(ctx context.Context, in *DeleteNaviRequest, opts ...http.CallOption) (*DeleteNaviReply, error) {
	var out DeleteNaviReply
	pattern := "/api/web-navigation/v1/delete_navi"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationNaviDeleteNavi))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NaviHTTPClientImpl) GetGuestSettings(ctx context.Context, in *GetGuestSettingsRequest, opts ...http.CallOption) (*GetGuestSettingsReply, error) {
	var out GetGuestSettingsReply
	pattern := "/api/web-navigation/v1/get_guest_settings"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNaviGetGuestSettings))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NaviHTTPClientImpl) ListNavi(ctx context.Context, in *ListNaviRequest, opts ...http.CallOption) (*ListNaviReply, error) {
	var out ListNaviReply
	pattern := "/api/web-navigation/v1/list_navi"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNaviListNavi))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NaviHTTPClientImpl) SetNotRemindInfo(ctx context.Context, in *SetNotRemindInfoRequest, opts ...http.CallOption) (*SetNotRemindInfoReply, error) {
	var out SetNotRemindInfoReply
	pattern := "/api/web-navigation/v1/set_not_remind_info"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationNaviSetNotRemindInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *NaviHTTPClientImpl) SortNavi(ctx context.Context, in *SortNaviRequest, opts ...http.CallOption) (*SortNaviReply, error) {
	var out SortNaviReply
	pattern := "/api/web-navigation/v1/sort_navi"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationNaviSortNavi))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}