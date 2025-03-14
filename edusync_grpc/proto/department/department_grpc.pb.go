// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: department.proto

package department

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
	DepartmentService_CreateDepartment_FullMethodName  = "/edusync_grpc.DepartmentService/CreateDepartment"
	DepartmentService_GetAllDepartments_FullMethodName = "/edusync_grpc.DepartmentService/GetAllDepartments"
	DepartmentService_GetDepartmentById_FullMethodName = "/edusync_grpc.DepartmentService/GetDepartmentById"
	DepartmentService_UpdateDepartment_FullMethodName  = "/edusync_grpc.DepartmentService/UpdateDepartment"
	DepartmentService_DeleteDepartment_FullMethodName  = "/edusync_grpc.DepartmentService/DeleteDepartment"
)

// DepartmentServiceClient is the client API for DepartmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DepartmentServiceClient interface {
	CreateDepartment(ctx context.Context, in *CreateDepartmentRequest, opts ...grpc.CallOption) (*CreateDepartmentResponse, error)
	GetAllDepartments(ctx context.Context, in *GetAllDepartmentsRequest, opts ...grpc.CallOption) (*GetAllDepartmentsResponse, error)
	GetDepartmentById(ctx context.Context, in *GetDepartmentByIdRequest, opts ...grpc.CallOption) (*GetDepartmentByIdResponse, error)
	UpdateDepartment(ctx context.Context, in *UpdateDepartmentRequest, opts ...grpc.CallOption) (*UpdateDepartmentResponse, error)
	DeleteDepartment(ctx context.Context, in *DeleteDepartmentRequest, opts ...grpc.CallOption) (*DeleteDepartmentResponse, error)
}

type departmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDepartmentServiceClient(cc grpc.ClientConnInterface) DepartmentServiceClient {
	return &departmentServiceClient{cc}
}

func (c *departmentServiceClient) CreateDepartment(ctx context.Context, in *CreateDepartmentRequest, opts ...grpc.CallOption) (*CreateDepartmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDepartmentResponse)
	err := c.cc.Invoke(ctx, DepartmentService_CreateDepartment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentServiceClient) GetAllDepartments(ctx context.Context, in *GetAllDepartmentsRequest, opts ...grpc.CallOption) (*GetAllDepartmentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllDepartmentsResponse)
	err := c.cc.Invoke(ctx, DepartmentService_GetAllDepartments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentServiceClient) GetDepartmentById(ctx context.Context, in *GetDepartmentByIdRequest, opts ...grpc.CallOption) (*GetDepartmentByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDepartmentByIdResponse)
	err := c.cc.Invoke(ctx, DepartmentService_GetDepartmentById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentServiceClient) UpdateDepartment(ctx context.Context, in *UpdateDepartmentRequest, opts ...grpc.CallOption) (*UpdateDepartmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateDepartmentResponse)
	err := c.cc.Invoke(ctx, DepartmentService_UpdateDepartment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentServiceClient) DeleteDepartment(ctx context.Context, in *DeleteDepartmentRequest, opts ...grpc.CallOption) (*DeleteDepartmentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDepartmentResponse)
	err := c.cc.Invoke(ctx, DepartmentService_DeleteDepartment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DepartmentServiceServer is the server API for DepartmentService service.
// All implementations must embed UnimplementedDepartmentServiceServer
// for forward compatibility.
type DepartmentServiceServer interface {
	CreateDepartment(context.Context, *CreateDepartmentRequest) (*CreateDepartmentResponse, error)
	GetAllDepartments(context.Context, *GetAllDepartmentsRequest) (*GetAllDepartmentsResponse, error)
	GetDepartmentById(context.Context, *GetDepartmentByIdRequest) (*GetDepartmentByIdResponse, error)
	UpdateDepartment(context.Context, *UpdateDepartmentRequest) (*UpdateDepartmentResponse, error)
	DeleteDepartment(context.Context, *DeleteDepartmentRequest) (*DeleteDepartmentResponse, error)
	mustEmbedUnimplementedDepartmentServiceServer()
}

// UnimplementedDepartmentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDepartmentServiceServer struct{}

func (UnimplementedDepartmentServiceServer) CreateDepartment(context.Context, *CreateDepartmentRequest) (*CreateDepartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDepartment not implemented")
}
func (UnimplementedDepartmentServiceServer) GetAllDepartments(context.Context, *GetAllDepartmentsRequest) (*GetAllDepartmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDepartments not implemented")
}
func (UnimplementedDepartmentServiceServer) GetDepartmentById(context.Context, *GetDepartmentByIdRequest) (*GetDepartmentByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepartmentById not implemented")
}
func (UnimplementedDepartmentServiceServer) UpdateDepartment(context.Context, *UpdateDepartmentRequest) (*UpdateDepartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDepartment not implemented")
}
func (UnimplementedDepartmentServiceServer) DeleteDepartment(context.Context, *DeleteDepartmentRequest) (*DeleteDepartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDepartment not implemented")
}
func (UnimplementedDepartmentServiceServer) mustEmbedUnimplementedDepartmentServiceServer() {}
func (UnimplementedDepartmentServiceServer) testEmbeddedByValue()                           {}

// UnsafeDepartmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DepartmentServiceServer will
// result in compilation errors.
type UnsafeDepartmentServiceServer interface {
	mustEmbedUnimplementedDepartmentServiceServer()
}

func RegisterDepartmentServiceServer(s grpc.ServiceRegistrar, srv DepartmentServiceServer) {
	// If the following call pancis, it indicates UnimplementedDepartmentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DepartmentService_ServiceDesc, srv)
}

func _DepartmentService_CreateDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).CreateDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentService_CreateDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).CreateDepartment(ctx, req.(*CreateDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentService_GetAllDepartments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllDepartmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).GetAllDepartments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentService_GetAllDepartments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).GetAllDepartments(ctx, req.(*GetAllDepartmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentService_GetDepartmentById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDepartmentByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).GetDepartmentById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentService_GetDepartmentById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).GetDepartmentById(ctx, req.(*GetDepartmentByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentService_UpdateDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).UpdateDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentService_UpdateDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).UpdateDepartment(ctx, req.(*UpdateDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DepartmentService_DeleteDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServiceServer).DeleteDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DepartmentService_DeleteDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServiceServer).DeleteDepartment(ctx, req.(*DeleteDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DepartmentService_ServiceDesc is the grpc.ServiceDesc for DepartmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DepartmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "edusync_grpc.DepartmentService",
	HandlerType: (*DepartmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDepartment",
			Handler:    _DepartmentService_CreateDepartment_Handler,
		},
		{
			MethodName: "GetAllDepartments",
			Handler:    _DepartmentService_GetAllDepartments_Handler,
		},
		{
			MethodName: "GetDepartmentById",
			Handler:    _DepartmentService_GetDepartmentById_Handler,
		},
		{
			MethodName: "UpdateDepartment",
			Handler:    _DepartmentService_UpdateDepartment_Handler,
		},
		{
			MethodName: "DeleteDepartment",
			Handler:    _DepartmentService_DeleteDepartment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "department.proto",
}
