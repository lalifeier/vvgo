// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.4

package v1

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

type TaoKeHTTPServer interface {
	GetCarouseList(context.Context, *GetCarouseListReq) (*GetCarouseListResp, error)
	GetSuperCategory(context.Context, *GetSuperCategoryReq) (*GetSuperCategoryResp, error)
}

func RegisterTaoKeHTTPServer(s *http.Server, srv TaoKeHTTPServer) {
	r := s.Route("/")
	r.GET("/api/goods/topic/carouse-list", _TaoKe_GetCarouseList0_HTTP_Handler(srv))
	r.GET("/api/category/get-super-category", _TaoKe_GetSuperCategory0_HTTP_Handler(srv))
}

func _TaoKe_GetCarouseList0_HTTP_Handler(srv TaoKeHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCarouseListReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.taoke.interface.v1.TaoKe/GetCarouseList")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCarouseList(ctx, req.(*GetCarouseListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCarouseListResp)
		return ctx.Result(200, reply)
	}
}

func _TaoKe_GetSuperCategory0_HTTP_Handler(srv TaoKeHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSuperCategoryReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.taoke.interface.v1.TaoKe/GetSuperCategory")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSuperCategory(ctx, req.(*GetSuperCategoryReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSuperCategoryResp)
		return ctx.Result(200, reply)
	}
}

type TaoKeHTTPClient interface {
	GetCarouseList(ctx context.Context, req *GetCarouseListReq, opts ...http.CallOption) (rsp *GetCarouseListResp, err error)
	GetSuperCategory(ctx context.Context, req *GetSuperCategoryReq, opts ...http.CallOption) (rsp *GetSuperCategoryResp, err error)
}

type TaoKeHTTPClientImpl struct {
	cc *http.Client
}

func NewTaoKeHTTPClient(client *http.Client) TaoKeHTTPClient {
	return &TaoKeHTTPClientImpl{client}
}

func (c *TaoKeHTTPClientImpl) GetCarouseList(ctx context.Context, in *GetCarouseListReq, opts ...http.CallOption) (*GetCarouseListResp, error) {
	var out GetCarouseListResp
	pattern := "/api/goods/topic/carouse-list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.taoke.interface.v1.TaoKe/GetCarouseList"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *TaoKeHTTPClientImpl) GetSuperCategory(ctx context.Context, in *GetSuperCategoryReq, opts ...http.CallOption) (*GetSuperCategoryResp, error) {
	var out GetSuperCategoryResp
	pattern := "/api/category/get-super-category"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.taoke.interface.v1.TaoKe/GetSuperCategory"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
