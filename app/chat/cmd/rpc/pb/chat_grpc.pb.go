// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: chat.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TableService_StoreAddItem_FullMethodName         = "/pb.TableService/StoreAddItem"
	TableService_SyncAddItem_FullMethodName          = "/pb.TableService/SyncAddItem"
	TableService_Send_FullMethodName                 = "/pb.TableService/Send"
	TableService_GetSyncMessage_FullMethodName       = "/pb.TableService/GetSyncMessage"
	TableService_GetHistoryMessage_FullMethodName    = "/pb.TableService/GetHistoryMessage"
	TableService_SearchHistoryMEssage_FullMethodName = "/pb.TableService/SearchHistoryMEssage"
)

// TableServiceClient is the client API for TableService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 负责阿里云table-store相关操作
type TableServiceClient interface {
	// has been depricated
	StoreAddItem(ctx context.Context, in *StoreAddItemReq, opts ...grpc.CallOption) (*StoreAddItemResp, error)
	SyncAddItem(ctx context.Context, in *SyncAddItemReq, opts ...grpc.CallOption) (*SyncAddItemResp, error)
	Send(ctx context.Context, in *SendReq, opts ...grpc.CallOption) (*SendResp, error)
	GetSyncMessage(ctx context.Context, in *GetSyncMessageReq, opts ...grpc.CallOption) (*GetSyncMessageResp, error)
	GetHistoryMessage(ctx context.Context, in *GetHistoryMessageReq, opts ...grpc.CallOption) (*GetHistoryMessageResp, error)
	SearchHistoryMEssage(ctx context.Context, in *SearchHistoryMessageReq, opts ...grpc.CallOption) (*SearchHistoryMessageResp, error)
}

type tableServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTableServiceClient(cc grpc.ClientConnInterface) TableServiceClient {
	return &tableServiceClient{cc}
}

func (c *tableServiceClient) StoreAddItem(ctx context.Context, in *StoreAddItemReq, opts ...grpc.CallOption) (*StoreAddItemResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StoreAddItemResp)
	err := c.cc.Invoke(ctx, TableService_StoreAddItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableServiceClient) SyncAddItem(ctx context.Context, in *SyncAddItemReq, opts ...grpc.CallOption) (*SyncAddItemResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SyncAddItemResp)
	err := c.cc.Invoke(ctx, TableService_SyncAddItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableServiceClient) Send(ctx context.Context, in *SendReq, opts ...grpc.CallOption) (*SendResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendResp)
	err := c.cc.Invoke(ctx, TableService_Send_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableServiceClient) GetSyncMessage(ctx context.Context, in *GetSyncMessageReq, opts ...grpc.CallOption) (*GetSyncMessageResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSyncMessageResp)
	err := c.cc.Invoke(ctx, TableService_GetSyncMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableServiceClient) GetHistoryMessage(ctx context.Context, in *GetHistoryMessageReq, opts ...grpc.CallOption) (*GetHistoryMessageResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHistoryMessageResp)
	err := c.cc.Invoke(ctx, TableService_GetHistoryMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tableServiceClient) SearchHistoryMEssage(ctx context.Context, in *SearchHistoryMessageReq, opts ...grpc.CallOption) (*SearchHistoryMessageResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchHistoryMessageResp)
	err := c.cc.Invoke(ctx, TableService_SearchHistoryMEssage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TableServiceServer is the server API for TableService service.
// All implementations must embed UnimplementedTableServiceServer
// for forward compatibility.
//
// 负责阿里云table-store相关操作
type TableServiceServer interface {
	// has been depricated
	StoreAddItem(context.Context, *StoreAddItemReq) (*StoreAddItemResp, error)
	SyncAddItem(context.Context, *SyncAddItemReq) (*SyncAddItemResp, error)
	Send(context.Context, *SendReq) (*SendResp, error)
	GetSyncMessage(context.Context, *GetSyncMessageReq) (*GetSyncMessageResp, error)
	GetHistoryMessage(context.Context, *GetHistoryMessageReq) (*GetHistoryMessageResp, error)
	SearchHistoryMEssage(context.Context, *SearchHistoryMessageReq) (*SearchHistoryMessageResp, error)
	mustEmbedUnimplementedTableServiceServer()
}

// UnimplementedTableServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTableServiceServer struct{}

func (UnimplementedTableServiceServer) StoreAddItem(context.Context, *StoreAddItemReq) (*StoreAddItemResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreAddItem not implemented")
}
func (UnimplementedTableServiceServer) SyncAddItem(context.Context, *SyncAddItemReq) (*SyncAddItemResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncAddItem not implemented")
}
func (UnimplementedTableServiceServer) Send(context.Context, *SendReq) (*SendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedTableServiceServer) GetSyncMessage(context.Context, *GetSyncMessageReq) (*GetSyncMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSyncMessage not implemented")
}
func (UnimplementedTableServiceServer) GetHistoryMessage(context.Context, *GetHistoryMessageReq) (*GetHistoryMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistoryMessage not implemented")
}
func (UnimplementedTableServiceServer) SearchHistoryMEssage(context.Context, *SearchHistoryMessageReq) (*SearchHistoryMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchHistoryMEssage not implemented")
}
func (UnimplementedTableServiceServer) mustEmbedUnimplementedTableServiceServer() {}
func (UnimplementedTableServiceServer) testEmbeddedByValue()                      {}

// UnsafeTableServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TableServiceServer will
// result in compilation errors.
type UnsafeTableServiceServer interface {
	mustEmbedUnimplementedTableServiceServer()
}

func RegisterTableServiceServer(s grpc.ServiceRegistrar, srv TableServiceServer) {
	// If the following call pancis, it indicates UnimplementedTableServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TableService_ServiceDesc, srv)
}

func _TableService_StoreAddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreAddItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TableServiceServer).StoreAddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TableService_StoreAddItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TableServiceServer).StoreAddItem(ctx, req.(*StoreAddItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TableService_SyncAddItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncAddItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TableServiceServer).SyncAddItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TableService_SyncAddItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TableServiceServer).SyncAddItem(ctx, req.(*SyncAddItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TableService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TableServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TableService_Send_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TableServiceServer).Send(ctx, req.(*SendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TableService_GetSyncMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSyncMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TableServiceServer).GetSyncMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TableService_GetSyncMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TableServiceServer).GetSyncMessage(ctx, req.(*GetSyncMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TableService_GetHistoryMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHistoryMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TableServiceServer).GetHistoryMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TableService_GetHistoryMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TableServiceServer).GetHistoryMessage(ctx, req.(*GetHistoryMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TableService_SearchHistoryMEssage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchHistoryMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TableServiceServer).SearchHistoryMEssage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TableService_SearchHistoryMEssage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TableServiceServer).SearchHistoryMEssage(ctx, req.(*SearchHistoryMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TableService_ServiceDesc is the grpc.ServiceDesc for TableService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TableService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TableService",
	HandlerType: (*TableServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreAddItem",
			Handler:    _TableService_StoreAddItem_Handler,
		},
		{
			MethodName: "SyncAddItem",
			Handler:    _TableService_SyncAddItem_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _TableService_Send_Handler,
		},
		{
			MethodName: "GetSyncMessage",
			Handler:    _TableService_GetSyncMessage_Handler,
		},
		{
			MethodName: "GetHistoryMessage",
			Handler:    _TableService_GetHistoryMessage_Handler,
		},
		{
			MethodName: "SearchHistoryMEssage",
			Handler:    _TableService_SearchHistoryMEssage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
