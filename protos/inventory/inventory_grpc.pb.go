// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.15.8
// source: inventory.proto

package inventory

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

// InventoryClient is the client API for Inventory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryClient interface {
	GetStock(ctx context.Context, in *LevelRequest, opts ...grpc.CallOption) (*StockItem, error)
	ChangeStock(ctx context.Context, in *AmendRequest, opts ...grpc.CallOption) (*AmendResponse, error)
	CheckShort(ctx context.Context, in *ShortRequest, opts ...grpc.CallOption) (*ShortList, error)
	GetStore(ctx context.Context, in *ShortRequest, opts ...grpc.CallOption) (*ShortList, error)
}

type inventoryClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryClient(cc grpc.ClientConnInterface) InventoryClient {
	return &inventoryClient{cc}
}

func (c *inventoryClient) GetStock(ctx context.Context, in *LevelRequest, opts ...grpc.CallOption) (*StockItem, error) {
	out := new(StockItem)
	err := c.cc.Invoke(ctx, "/Inventory/GetStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) ChangeStock(ctx context.Context, in *AmendRequest, opts ...grpc.CallOption) (*AmendResponse, error) {
	out := new(AmendResponse)
	err := c.cc.Invoke(ctx, "/Inventory/ChangeStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) CheckShort(ctx context.Context, in *ShortRequest, opts ...grpc.CallOption) (*ShortList, error) {
	out := new(ShortList)
	err := c.cc.Invoke(ctx, "/Inventory/CheckShort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryClient) GetStore(ctx context.Context, in *ShortRequest, opts ...grpc.CallOption) (*ShortList, error) {
	out := new(ShortList)
	err := c.cc.Invoke(ctx, "/Inventory/GetStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServer is the server API for Inventory service.
// All implementations should embed UnimplementedInventoryServer
// for forward compatibility
type InventoryServer interface {
	GetStock(context.Context, *LevelRequest) (*StockItem, error)
	ChangeStock(context.Context, *AmendRequest) (*AmendResponse, error)
	CheckShort(context.Context, *ShortRequest) (*ShortList, error)
	GetStore(context.Context, *ShortRequest) (*ShortList, error)
}

// UnimplementedInventoryServer should be embedded to have forward compatible implementations.
type UnimplementedInventoryServer struct {
}

func (UnimplementedInventoryServer) GetStock(context.Context, *LevelRequest) (*StockItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStock not implemented")
}
func (UnimplementedInventoryServer) ChangeStock(context.Context, *AmendRequest) (*AmendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeStock not implemented")
}
func (UnimplementedInventoryServer) CheckShort(context.Context, *ShortRequest) (*ShortList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckShort not implemented")
}
func (UnimplementedInventoryServer) GetStore(context.Context, *ShortRequest) (*ShortList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStore not implemented")
}

// UnsafeInventoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryServer will
// result in compilation errors.
type UnsafeInventoryServer interface {
	mustEmbedUnimplementedInventoryServer()
}

func RegisterInventoryServer(s grpc.ServiceRegistrar, srv InventoryServer) {
	s.RegisterService(&Inventory_ServiceDesc, srv)
}

func _Inventory_GetStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).GetStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Inventory/GetStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).GetStock(ctx, req.(*LevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_ChangeStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AmendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).ChangeStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Inventory/ChangeStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).ChangeStock(ctx, req.(*AmendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_CheckShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).CheckShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Inventory/CheckShort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).CheckShort(ctx, req.(*ShortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventory_GetStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).GetStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Inventory/GetStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).GetStore(ctx, req.(*ShortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Inventory_ServiceDesc is the grpc.ServiceDesc for Inventory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Inventory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Inventory",
	HandlerType: (*InventoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStock",
			Handler:    _Inventory_GetStock_Handler,
		},
		{
			MethodName: "ChangeStock",
			Handler:    _Inventory_ChangeStock_Handler,
		},
		{
			MethodName: "CheckShort",
			Handler:    _Inventory_CheckShort_Handler,
		},
		{
			MethodName: "GetStore",
			Handler:    _Inventory_GetStore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory.proto",
}
