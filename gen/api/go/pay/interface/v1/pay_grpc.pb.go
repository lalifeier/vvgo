// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: pay/interface/v1/pay.proto

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

// PayClient is the client API for Pay service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PayClient interface {
	// 发起支付
	GoPay(ctx context.Context, in *GoPayReq, opts ...grpc.CallOption) (*GoPayReply, error)
	// 发起退款
	Refund(ctx context.Context, in *GoPayReq, opts ...grpc.CallOption) (*GoPayReply, error)
	// 接口异步通知
	AsyncNotify(ctx context.Context, in *GoPayReq, opts ...grpc.CallOption) (*GoPayReply, error)
	// 接口同步通知
	SyncNotify(ctx context.Context, in *GoPayReq, opts ...grpc.CallOption) (*GoPayReply, error)
	// 交易查询
	QueryTrade(ctx context.Context, in *GoPayReq, opts ...grpc.CallOption) (*GoPayReply, error)
	// 退款查询
	QueryRefund(ctx context.Context, in *GoPayReq, opts ...grpc.CallOption) (*GoPayReply, error)
	// 账单获取
	QueryBill(ctx context.Context, in *GoPayReq, opts ...grpc.CallOption) (*GoPayReply, error)
	// 结算明细
	QuerySettle(ctx context.Context, in *GoPayReq, opts ...grpc.CallOption) (*GoPayReply, error)
}

type payClient struct {
	cc grpc.ClientConnInterface
}

func NewPayClient(cc grpc.ClientConnInterface) PayClient {
	return &payClient{cc}
}

func (c *payClient) GoPay(ctx context.Context, in *GoPayReq, opts ...grpc.CallOption) (*GoPayReply, error) {
	out := new(GoPayReply)
	err := c.cc.Invoke(ctx, "/api.pay.service.v1.Pay/GoPay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payClient) Refund(ctx context.Context, in *GoPayReq, opts ...grpc.CallOption) (*GoPayReply, error) {
	out := new(GoPayReply)
	err := c.cc.Invoke(ctx, "/api.pay.service.v1.Pay/Refund", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payClient) AsyncNotify(ctx context.Context, in *GoPayReq, opts ...grpc.CallOption) (*GoPayReply, error) {
	out := new(GoPayReply)
	err := c.cc.Invoke(ctx, "/api.pay.service.v1.Pay/AsyncNotify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payClient) SyncNotify(ctx context.Context, in *GoPayReq, opts ...grpc.CallOption) (*GoPayReply, error) {
	out := new(GoPayReply)
	err := c.cc.Invoke(ctx, "/api.pay.service.v1.Pay/SyncNotify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payClient) QueryTrade(ctx context.Context, in *GoPayReq, opts ...grpc.CallOption) (*GoPayReply, error) {
	out := new(GoPayReply)
	err := c.cc.Invoke(ctx, "/api.pay.service.v1.Pay/QueryTrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payClient) QueryRefund(ctx context.Context, in *GoPayReq, opts ...grpc.CallOption) (*GoPayReply, error) {
	out := new(GoPayReply)
	err := c.cc.Invoke(ctx, "/api.pay.service.v1.Pay/QueryRefund", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payClient) QueryBill(ctx context.Context, in *GoPayReq, opts ...grpc.CallOption) (*GoPayReply, error) {
	out := new(GoPayReply)
	err := c.cc.Invoke(ctx, "/api.pay.service.v1.Pay/QueryBill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payClient) QuerySettle(ctx context.Context, in *GoPayReq, opts ...grpc.CallOption) (*GoPayReply, error) {
	out := new(GoPayReply)
	err := c.cc.Invoke(ctx, "/api.pay.service.v1.Pay/QuerySettle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayServer is the server API for Pay service.
// All implementations must embed UnimplementedPayServer
// for forward compatibility
type PayServer interface {
	// 发起支付
	GoPay(context.Context, *GoPayReq) (*GoPayReply, error)
	// 发起退款
	Refund(context.Context, *GoPayReq) (*GoPayReply, error)
	// 接口异步通知
	AsyncNotify(context.Context, *GoPayReq) (*GoPayReply, error)
	// 接口同步通知
	SyncNotify(context.Context, *GoPayReq) (*GoPayReply, error)
	// 交易查询
	QueryTrade(context.Context, *GoPayReq) (*GoPayReply, error)
	// 退款查询
	QueryRefund(context.Context, *GoPayReq) (*GoPayReply, error)
	// 账单获取
	QueryBill(context.Context, *GoPayReq) (*GoPayReply, error)
	// 结算明细
	QuerySettle(context.Context, *GoPayReq) (*GoPayReply, error)
	mustEmbedUnimplementedPayServer()
}

// UnimplementedPayServer must be embedded to have forward compatible implementations.
type UnimplementedPayServer struct {
}

func (UnimplementedPayServer) GoPay(context.Context, *GoPayReq) (*GoPayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoPay not implemented")
}
func (UnimplementedPayServer) Refund(context.Context, *GoPayReq) (*GoPayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refund not implemented")
}
func (UnimplementedPayServer) AsyncNotify(context.Context, *GoPayReq) (*GoPayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AsyncNotify not implemented")
}
func (UnimplementedPayServer) SyncNotify(context.Context, *GoPayReq) (*GoPayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncNotify not implemented")
}
func (UnimplementedPayServer) QueryTrade(context.Context, *GoPayReq) (*GoPayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTrade not implemented")
}
func (UnimplementedPayServer) QueryRefund(context.Context, *GoPayReq) (*GoPayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRefund not implemented")
}
func (UnimplementedPayServer) QueryBill(context.Context, *GoPayReq) (*GoPayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryBill not implemented")
}
func (UnimplementedPayServer) QuerySettle(context.Context, *GoPayReq) (*GoPayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySettle not implemented")
}
func (UnimplementedPayServer) mustEmbedUnimplementedPayServer() {}

// UnsafePayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PayServer will
// result in compilation errors.
type UnsafePayServer interface {
	mustEmbedUnimplementedPayServer()
}

func RegisterPayServer(s grpc.ServiceRegistrar, srv PayServer) {
	s.RegisterService(&Pay_ServiceDesc, srv)
}

func _Pay_GoPay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoPayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).GoPay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.pay.service.v1.Pay/GoPay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).GoPay(ctx, req.(*GoPayReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pay_Refund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoPayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).Refund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.pay.service.v1.Pay/Refund",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).Refund(ctx, req.(*GoPayReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pay_AsyncNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoPayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).AsyncNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.pay.service.v1.Pay/AsyncNotify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).AsyncNotify(ctx, req.(*GoPayReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pay_SyncNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoPayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).SyncNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.pay.service.v1.Pay/SyncNotify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).SyncNotify(ctx, req.(*GoPayReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pay_QueryTrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoPayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).QueryTrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.pay.service.v1.Pay/QueryTrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).QueryTrade(ctx, req.(*GoPayReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pay_QueryRefund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoPayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).QueryRefund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.pay.service.v1.Pay/QueryRefund",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).QueryRefund(ctx, req.(*GoPayReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pay_QueryBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoPayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).QueryBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.pay.service.v1.Pay/QueryBill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).QueryBill(ctx, req.(*GoPayReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pay_QuerySettle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoPayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).QuerySettle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.pay.service.v1.Pay/QuerySettle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).QuerySettle(ctx, req.(*GoPayReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Pay_ServiceDesc is the grpc.ServiceDesc for Pay service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pay_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.pay.service.v1.Pay",
	HandlerType: (*PayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GoPay",
			Handler:    _Pay_GoPay_Handler,
		},
		{
			MethodName: "Refund",
			Handler:    _Pay_Refund_Handler,
		},
		{
			MethodName: "AsyncNotify",
			Handler:    _Pay_AsyncNotify_Handler,
		},
		{
			MethodName: "SyncNotify",
			Handler:    _Pay_SyncNotify_Handler,
		},
		{
			MethodName: "QueryTrade",
			Handler:    _Pay_QueryTrade_Handler,
		},
		{
			MethodName: "QueryRefund",
			Handler:    _Pay_QueryRefund_Handler,
		},
		{
			MethodName: "QueryBill",
			Handler:    _Pay_QueryBill_Handler,
		},
		{
			MethodName: "QuerySettle",
			Handler:    _Pay_QuerySettle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pay/interface/v1/pay.proto",
}
