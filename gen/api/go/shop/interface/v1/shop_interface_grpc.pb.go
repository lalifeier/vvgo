// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: shop/interface/v1/shop_interface.proto

package v1

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

// ShopInterfaceClient is the client API for ShopInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopInterfaceClient interface {
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	//	rpc QuickLogin(LoginReq) returns (LoginResp) {
	//	  option (google.api.http) = {
	//	    post : "/api/quickLogin"
	//	    body : "*"
	//	  };
	//	}
	GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
}

type shopInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewShopInterfaceClient(cc grpc.ClientConnInterface) ShopInterfaceClient {
	return &shopInterfaceClient{cc}
}

func (c *shopInterfaceClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, "/api.sms.interface.v1.ShopInterface/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/api.sms.interface.v1.ShopInterface/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopInterfaceClient) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	out := new(GetUserInfoResp)
	err := c.cc.Invoke(ctx, "/api.sms.interface.v1.ShopInterface/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopInterfaceServer is the server API for ShopInterface service.
// All implementations must embed UnimplementedShopInterfaceServer
// for forward compatibility
type ShopInterfaceServer interface {
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	Login(context.Context, *LoginReq) (*LoginResp, error)
	//	rpc QuickLogin(LoginReq) returns (LoginResp) {
	//	  option (google.api.http) = {
	//	    post : "/api/quickLogin"
	//	    body : "*"
	//	  };
	//	}
	GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error)
	mustEmbedUnimplementedShopInterfaceServer()
}

// UnimplementedShopInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedShopInterfaceServer struct {
}

func (UnimplementedShopInterfaceServer) Register(context.Context, *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedShopInterfaceServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedShopInterfaceServer) GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedShopInterfaceServer) mustEmbedUnimplementedShopInterfaceServer() {}

// UnsafeShopInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopInterfaceServer will
// result in compilation errors.
type UnsafeShopInterfaceServer interface {
	mustEmbedUnimplementedShopInterfaceServer()
}

func RegisterShopInterfaceServer(s grpc.ServiceRegistrar, srv ShopInterfaceServer) {
	s.RegisterService(&ShopInterface_ServiceDesc, srv)
}

func _ShopInterface_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sms.interface.v1.ShopInterface/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sms.interface.v1.ShopInterface/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopInterface_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopInterfaceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sms.interface.v1.ShopInterface/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopInterfaceServer).GetUserInfo(ctx, req.(*GetUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ShopInterface_ServiceDesc is the grpc.ServiceDesc for ShopInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShopInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.sms.interface.v1.ShopInterface",
	HandlerType: (*ShopInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _ShopInterface_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _ShopInterface_Login_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _ShopInterface_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop/interface/v1/shop_interface.proto",
}
