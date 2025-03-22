// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: pb/recycle_hub.proto

package recyclehub

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
	WasteTypeService_GetAllWasteTypes_FullMethodName = "/pb.recyclehub.WasteTypeService/GetAllWasteTypes"
)

// WasteTypeServiceClient is the client API for WasteTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// WasteTypeService defines the gRPC service for managing waste types.
type WasteTypeServiceClient interface {
	// GetAllWasteTypes retrieves all waste types.
	GetAllWasteTypes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*WasteTypeList, error)
}

type wasteTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWasteTypeServiceClient(cc grpc.ClientConnInterface) WasteTypeServiceClient {
	return &wasteTypeServiceClient{cc}
}

func (c *wasteTypeServiceClient) GetAllWasteTypes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*WasteTypeList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WasteTypeList)
	err := c.cc.Invoke(ctx, WasteTypeService_GetAllWasteTypes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WasteTypeServiceServer is the server API for WasteTypeService service.
// All implementations must embed UnimplementedWasteTypeServiceServer
// for forward compatibility.
//
// WasteTypeService defines the gRPC service for managing waste types.
type WasteTypeServiceServer interface {
	// GetAllWasteTypes retrieves all waste types.
	GetAllWasteTypes(context.Context, *Empty) (*WasteTypeList, error)
	mustEmbedUnimplementedWasteTypeServiceServer()
}

// UnimplementedWasteTypeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWasteTypeServiceServer struct{}

func (UnimplementedWasteTypeServiceServer) GetAllWasteTypes(context.Context, *Empty) (*WasteTypeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllWasteTypes not implemented")
}
func (UnimplementedWasteTypeServiceServer) mustEmbedUnimplementedWasteTypeServiceServer() {}
func (UnimplementedWasteTypeServiceServer) testEmbeddedByValue()                          {}

// UnsafeWasteTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WasteTypeServiceServer will
// result in compilation errors.
type UnsafeWasteTypeServiceServer interface {
	mustEmbedUnimplementedWasteTypeServiceServer()
}

func RegisterWasteTypeServiceServer(s grpc.ServiceRegistrar, srv WasteTypeServiceServer) {
	// If the following call pancis, it indicates UnimplementedWasteTypeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WasteTypeService_ServiceDesc, srv)
}

func _WasteTypeService_GetAllWasteTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WasteTypeServiceServer).GetAllWasteTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WasteTypeService_GetAllWasteTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WasteTypeServiceServer).GetAllWasteTypes(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// WasteTypeService_ServiceDesc is the grpc.ServiceDesc for WasteTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WasteTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.recyclehub.WasteTypeService",
	HandlerType: (*WasteTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllWasteTypes",
			Handler:    _WasteTypeService_GetAllWasteTypes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/recycle_hub.proto",
}

const (
	RecycleHubService_CreateRecycleHub_FullMethodName  = "/pb.recyclehub.RecycleHubService/CreateRecycleHub"
	RecycleHubService_GetAllRecycleHubs_FullMethodName = "/pb.recyclehub.RecycleHubService/GetAllRecycleHubs"
	RecycleHubService_GetRecycleHubByID_FullMethodName = "/pb.recyclehub.RecycleHubService/GetRecycleHubByID"
	RecycleHubService_UpdateRecycleHub_FullMethodName  = "/pb.recyclehub.RecycleHubService/UpdateRecycleHub"
	RecycleHubService_DeleteRecycleHub_FullMethodName  = "/pb.recyclehub.RecycleHubService/DeleteRecycleHub"
)

// RecycleHubServiceClient is the client API for RecycleHubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// RecycleHubService defines the gRPC service for managing recycle hubs.
type RecycleHubServiceClient interface {
	// CreateRecycleHub creates a new recycle hub.
	CreateRecycleHub(ctx context.Context, in *CreateRecycleHubRequest, opts ...grpc.CallOption) (*RecycleHubResponse, error)
	// GetAllRecycleHubs retrieves all recycle hubs.
	GetAllRecycleHubs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RecycleHubList, error)
	// GetRecycleHubByID retrieves a recycle hub by its ID.
	GetRecycleHubByID(ctx context.Context, in *GetRecycleHubByIDRequest, opts ...grpc.CallOption) (*RecycleHubResponse, error)
	// UpdateRecycleHub updates an existing recycle hub.
	UpdateRecycleHub(ctx context.Context, in *UpdateRecycleHubRequest, opts ...grpc.CallOption) (*RecycleHubResponse, error)
	// DeleteRecycleHub deletes a recycle hub by its ID.
	DeleteRecycleHub(ctx context.Context, in *DeleteRecycleHubRequest, opts ...grpc.CallOption) (*Empty, error)
}

type recycleHubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecycleHubServiceClient(cc grpc.ClientConnInterface) RecycleHubServiceClient {
	return &recycleHubServiceClient{cc}
}

func (c *recycleHubServiceClient) CreateRecycleHub(ctx context.Context, in *CreateRecycleHubRequest, opts ...grpc.CallOption) (*RecycleHubResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RecycleHubResponse)
	err := c.cc.Invoke(ctx, RecycleHubService_CreateRecycleHub_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recycleHubServiceClient) GetAllRecycleHubs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RecycleHubList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RecycleHubList)
	err := c.cc.Invoke(ctx, RecycleHubService_GetAllRecycleHubs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recycleHubServiceClient) GetRecycleHubByID(ctx context.Context, in *GetRecycleHubByIDRequest, opts ...grpc.CallOption) (*RecycleHubResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RecycleHubResponse)
	err := c.cc.Invoke(ctx, RecycleHubService_GetRecycleHubByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recycleHubServiceClient) UpdateRecycleHub(ctx context.Context, in *UpdateRecycleHubRequest, opts ...grpc.CallOption) (*RecycleHubResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RecycleHubResponse)
	err := c.cc.Invoke(ctx, RecycleHubService_UpdateRecycleHub_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recycleHubServiceClient) DeleteRecycleHub(ctx context.Context, in *DeleteRecycleHubRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, RecycleHubService_DeleteRecycleHub_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecycleHubServiceServer is the server API for RecycleHubService service.
// All implementations must embed UnimplementedRecycleHubServiceServer
// for forward compatibility.
//
// RecycleHubService defines the gRPC service for managing recycle hubs.
type RecycleHubServiceServer interface {
	// CreateRecycleHub creates a new recycle hub.
	CreateRecycleHub(context.Context, *CreateRecycleHubRequest) (*RecycleHubResponse, error)
	// GetAllRecycleHubs retrieves all recycle hubs.
	GetAllRecycleHubs(context.Context, *Empty) (*RecycleHubList, error)
	// GetRecycleHubByID retrieves a recycle hub by its ID.
	GetRecycleHubByID(context.Context, *GetRecycleHubByIDRequest) (*RecycleHubResponse, error)
	// UpdateRecycleHub updates an existing recycle hub.
	UpdateRecycleHub(context.Context, *UpdateRecycleHubRequest) (*RecycleHubResponse, error)
	// DeleteRecycleHub deletes a recycle hub by its ID.
	DeleteRecycleHub(context.Context, *DeleteRecycleHubRequest) (*Empty, error)
	mustEmbedUnimplementedRecycleHubServiceServer()
}

// UnimplementedRecycleHubServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRecycleHubServiceServer struct{}

func (UnimplementedRecycleHubServiceServer) CreateRecycleHub(context.Context, *CreateRecycleHubRequest) (*RecycleHubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecycleHub not implemented")
}
func (UnimplementedRecycleHubServiceServer) GetAllRecycleHubs(context.Context, *Empty) (*RecycleHubList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRecycleHubs not implemented")
}
func (UnimplementedRecycleHubServiceServer) GetRecycleHubByID(context.Context, *GetRecycleHubByIDRequest) (*RecycleHubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecycleHubByID not implemented")
}
func (UnimplementedRecycleHubServiceServer) UpdateRecycleHub(context.Context, *UpdateRecycleHubRequest) (*RecycleHubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecycleHub not implemented")
}
func (UnimplementedRecycleHubServiceServer) DeleteRecycleHub(context.Context, *DeleteRecycleHubRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecycleHub not implemented")
}
func (UnimplementedRecycleHubServiceServer) mustEmbedUnimplementedRecycleHubServiceServer() {}
func (UnimplementedRecycleHubServiceServer) testEmbeddedByValue()                           {}

// UnsafeRecycleHubServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecycleHubServiceServer will
// result in compilation errors.
type UnsafeRecycleHubServiceServer interface {
	mustEmbedUnimplementedRecycleHubServiceServer()
}

func RegisterRecycleHubServiceServer(s grpc.ServiceRegistrar, srv RecycleHubServiceServer) {
	// If the following call pancis, it indicates UnimplementedRecycleHubServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RecycleHubService_ServiceDesc, srv)
}

func _RecycleHubService_CreateRecycleHub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecycleHubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecycleHubServiceServer).CreateRecycleHub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecycleHubService_CreateRecycleHub_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecycleHubServiceServer).CreateRecycleHub(ctx, req.(*CreateRecycleHubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecycleHubService_GetAllRecycleHubs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecycleHubServiceServer).GetAllRecycleHubs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecycleHubService_GetAllRecycleHubs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecycleHubServiceServer).GetAllRecycleHubs(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecycleHubService_GetRecycleHubByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecycleHubByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecycleHubServiceServer).GetRecycleHubByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecycleHubService_GetRecycleHubByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecycleHubServiceServer).GetRecycleHubByID(ctx, req.(*GetRecycleHubByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecycleHubService_UpdateRecycleHub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRecycleHubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecycleHubServiceServer).UpdateRecycleHub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecycleHubService_UpdateRecycleHub_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecycleHubServiceServer).UpdateRecycleHub(ctx, req.(*UpdateRecycleHubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecycleHubService_DeleteRecycleHub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecycleHubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecycleHubServiceServer).DeleteRecycleHub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecycleHubService_DeleteRecycleHub_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecycleHubServiceServer).DeleteRecycleHub(ctx, req.(*DeleteRecycleHubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecycleHubService_ServiceDesc is the grpc.ServiceDesc for RecycleHubService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecycleHubService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.recyclehub.RecycleHubService",
	HandlerType: (*RecycleHubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRecycleHub",
			Handler:    _RecycleHubService_CreateRecycleHub_Handler,
		},
		{
			MethodName: "GetAllRecycleHubs",
			Handler:    _RecycleHubService_GetAllRecycleHubs_Handler,
		},
		{
			MethodName: "GetRecycleHubByID",
			Handler:    _RecycleHubService_GetRecycleHubByID_Handler,
		},
		{
			MethodName: "UpdateRecycleHub",
			Handler:    _RecycleHubService_UpdateRecycleHub_Handler,
		},
		{
			MethodName: "DeleteRecycleHub",
			Handler:    _RecycleHubService_DeleteRecycleHub_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/recycle_hub.proto",
}
