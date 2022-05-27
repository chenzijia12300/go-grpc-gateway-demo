// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: Goods.proto

package proto

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

// GoodsServiceClient is the client API for GoodsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodsServiceClient interface {
	SaveGoods(ctx context.Context, in *Goods, opts ...grpc.CallOption) (*GoodsId, error)
	ModifyGoods(ctx context.Context, in *Goods, opts ...grpc.CallOption) (*GoodsId, error)
	DeleteGoods(ctx context.Context, in *GoodsIds, opts ...grpc.CallOption) (*DelGoodsResponse, error)
	GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsResponse, error)
}

type goodsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodsServiceClient(cc grpc.ClientConnInterface) GoodsServiceClient {
	return &goodsServiceClient{cc}
}

func (c *goodsServiceClient) SaveGoods(ctx context.Context, in *Goods, opts ...grpc.CallOption) (*GoodsId, error) {
	out := new(GoodsId)
	err := c.cc.Invoke(ctx, "/proto.GoodsService/SaveGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) ModifyGoods(ctx context.Context, in *Goods, opts ...grpc.CallOption) (*GoodsId, error) {
	out := new(GoodsId)
	err := c.cc.Invoke(ctx, "/proto.GoodsService/ModifyGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) DeleteGoods(ctx context.Context, in *GoodsIds, opts ...grpc.CallOption) (*DelGoodsResponse, error) {
	out := new(DelGoodsResponse)
	err := c.cc.Invoke(ctx, "/proto.GoodsService/DeleteGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsServiceClient) GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsResponse, error) {
	out := new(GetGoodsResponse)
	err := c.cc.Invoke(ctx, "/proto.GoodsService/GetGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsServiceServer is the server API for GoodsService service.
// All implementations should embed UnimplementedGoodsServiceServer
// for forward compatibility
type GoodsServiceServer interface {
	SaveGoods(context.Context, *Goods) (*GoodsId, error)
	ModifyGoods(context.Context, *Goods) (*GoodsId, error)
	DeleteGoods(context.Context, *GoodsIds) (*DelGoodsResponse, error)
	GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsResponse, error)
}

// UnimplementedGoodsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedGoodsServiceServer struct {
}

func (UnimplementedGoodsServiceServer) SaveGoods(context.Context, *Goods) (*GoodsId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveGoods not implemented")
}
func (UnimplementedGoodsServiceServer) ModifyGoods(context.Context, *Goods) (*GoodsId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyGoods not implemented")
}
func (UnimplementedGoodsServiceServer) DeleteGoods(context.Context, *GoodsIds) (*DelGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGoods not implemented")
}
func (UnimplementedGoodsServiceServer) GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoods not implemented")
}

// UnsafeGoodsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsServiceServer will
// result in compilation errors.
type UnsafeGoodsServiceServer interface {
	mustEmbedUnimplementedGoodsServiceServer()
}

func RegisterGoodsServiceServer(s grpc.ServiceRegistrar, srv GoodsServiceServer) {
	s.RegisterService(&GoodsService_ServiceDesc, srv)
}

func _GoodsService_SaveGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Goods)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).SaveGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GoodsService/SaveGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).SaveGoods(ctx, req.(*Goods))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_ModifyGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Goods)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).ModifyGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GoodsService/ModifyGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).ModifyGoods(ctx, req.(*Goods))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_DeleteGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).DeleteGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GoodsService/DeleteGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).DeleteGoods(ctx, req.(*GoodsIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsService_GetGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServiceServer).GetGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GoodsService/GetGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServiceServer).GetGoods(ctx, req.(*GetGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoodsService_ServiceDesc is the grpc.ServiceDesc for GoodsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoodsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GoodsService",
	HandlerType: (*GoodsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveGoods",
			Handler:    _GoodsService_SaveGoods_Handler,
		},
		{
			MethodName: "ModifyGoods",
			Handler:    _GoodsService_ModifyGoods_Handler,
		},
		{
			MethodName: "DeleteGoods",
			Handler:    _GoodsService_DeleteGoods_Handler,
		},
		{
			MethodName: "GetGoods",
			Handler:    _GoodsService_GetGoods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Goods.proto",
}