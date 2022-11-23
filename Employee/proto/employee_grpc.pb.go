// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: employee.proto

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

// EmployeeServiceClient is the client API for EmployeeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmployeeServiceClient interface {
	CreateEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*EmployeeId, error)
	ReadEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*Employees, error)
	ReadEmployeeById(ctx context.Context, in *EmployeeId, opts ...grpc.CallOption) (*Employees, error)
	ReadEmployeeByFirstName(ctx context.Context, in *EmployeeFirstName, opts ...grpc.CallOption) (*Employees, error)
	ReadEmployeeByLastName(ctx context.Context, in *EmployeeLastName, opts ...grpc.CallOption) (*Employees, error)
	ReadEmployeeByDisplayName(ctx context.Context, in *EmployeeDisplayName, opts ...grpc.CallOption) (*Employees, error)
	ReadEmployeeByAge(ctx context.Context, in *EmployeeAge, opts ...grpc.CallOption) (*Employees, error)
	ReadEmployeeBySalary(ctx context.Context, in *EmployeeSalary, opts ...grpc.CallOption) (*Employees, error)
	ReadEmployeeByDesignation(ctx context.Context, in *EmployeeDesignation, opts ...grpc.CallOption) (*Employees, error)
	ReadEmployeeByDepartment(ctx context.Context, in *EmployeeDepartment, opts ...grpc.CallOption) (*Employees, error)
	UpdateEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*UpdateEmployeeResponse, error)
	DeleteEmployee(ctx context.Context, in *EmployeeId, opts ...grpc.CallOption) (*DeleteEmployeeResponse, error)
}

type employeeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployeeServiceClient(cc grpc.ClientConnInterface) EmployeeServiceClient {
	return &employeeServiceClient{cc}
}

func (c *employeeServiceClient) CreateEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*EmployeeId, error) {
	out := new(EmployeeId)
	err := c.cc.Invoke(ctx, "/proto.EmployeeService/CreateEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) ReadEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*Employees, error) {
	out := new(Employees)
	err := c.cc.Invoke(ctx, "/proto.EmployeeService/ReadEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) ReadEmployeeById(ctx context.Context, in *EmployeeId, opts ...grpc.CallOption) (*Employees, error) {
	out := new(Employees)
	err := c.cc.Invoke(ctx, "/proto.EmployeeService/ReadEmployeeById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) ReadEmployeeByFirstName(ctx context.Context, in *EmployeeFirstName, opts ...grpc.CallOption) (*Employees, error) {
	out := new(Employees)
	err := c.cc.Invoke(ctx, "/proto.EmployeeService/ReadEmployeeByFirstName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) ReadEmployeeByLastName(ctx context.Context, in *EmployeeLastName, opts ...grpc.CallOption) (*Employees, error) {
	out := new(Employees)
	err := c.cc.Invoke(ctx, "/proto.EmployeeService/ReadEmployeeByLastName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) ReadEmployeeByDisplayName(ctx context.Context, in *EmployeeDisplayName, opts ...grpc.CallOption) (*Employees, error) {
	out := new(Employees)
	err := c.cc.Invoke(ctx, "/proto.EmployeeService/ReadEmployeeByDisplayName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) ReadEmployeeByAge(ctx context.Context, in *EmployeeAge, opts ...grpc.CallOption) (*Employees, error) {
	out := new(Employees)
	err := c.cc.Invoke(ctx, "/proto.EmployeeService/ReadEmployeeByAge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) ReadEmployeeBySalary(ctx context.Context, in *EmployeeSalary, opts ...grpc.CallOption) (*Employees, error) {
	out := new(Employees)
	err := c.cc.Invoke(ctx, "/proto.EmployeeService/ReadEmployeeBySalary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) ReadEmployeeByDesignation(ctx context.Context, in *EmployeeDesignation, opts ...grpc.CallOption) (*Employees, error) {
	out := new(Employees)
	err := c.cc.Invoke(ctx, "/proto.EmployeeService/ReadEmployeeByDesignation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) ReadEmployeeByDepartment(ctx context.Context, in *EmployeeDepartment, opts ...grpc.CallOption) (*Employees, error) {
	out := new(Employees)
	err := c.cc.Invoke(ctx, "/proto.EmployeeService/ReadEmployeeByDepartment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) UpdateEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*UpdateEmployeeResponse, error) {
	out := new(UpdateEmployeeResponse)
	err := c.cc.Invoke(ctx, "/proto.EmployeeService/UpdateEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) DeleteEmployee(ctx context.Context, in *EmployeeId, opts ...grpc.CallOption) (*DeleteEmployeeResponse, error) {
	out := new(DeleteEmployeeResponse)
	err := c.cc.Invoke(ctx, "/proto.EmployeeService/DeleteEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployeeServiceServer is the server API for EmployeeService service.
// All implementations must embed UnimplementedEmployeeServiceServer
// for forward compatibility
type EmployeeServiceServer interface {
	CreateEmployee(context.Context, *Employee) (*EmployeeId, error)
	ReadEmployee(context.Context, *Employee) (*Employees, error)
	ReadEmployeeById(context.Context, *EmployeeId) (*Employees, error)
	ReadEmployeeByFirstName(context.Context, *EmployeeFirstName) (*Employees, error)
	ReadEmployeeByLastName(context.Context, *EmployeeLastName) (*Employees, error)
	ReadEmployeeByDisplayName(context.Context, *EmployeeDisplayName) (*Employees, error)
	ReadEmployeeByAge(context.Context, *EmployeeAge) (*Employees, error)
	ReadEmployeeBySalary(context.Context, *EmployeeSalary) (*Employees, error)
	ReadEmployeeByDesignation(context.Context, *EmployeeDesignation) (*Employees, error)
	ReadEmployeeByDepartment(context.Context, *EmployeeDepartment) (*Employees, error)
	UpdateEmployee(context.Context, *Employee) (*UpdateEmployeeResponse, error)
	DeleteEmployee(context.Context, *EmployeeId) (*DeleteEmployeeResponse, error)
	mustEmbedUnimplementedEmployeeServiceServer()
}

// UnimplementedEmployeeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmployeeServiceServer struct {
}

func (UnimplementedEmployeeServiceServer) CreateEmployee(context.Context, *Employee) (*EmployeeId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmployee not implemented")
}
func (UnimplementedEmployeeServiceServer) ReadEmployee(context.Context, *Employee) (*Employees, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadEmployee not implemented")
}
func (UnimplementedEmployeeServiceServer) ReadEmployeeById(context.Context, *EmployeeId) (*Employees, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadEmployeeById not implemented")
}
func (UnimplementedEmployeeServiceServer) ReadEmployeeByFirstName(context.Context, *EmployeeFirstName) (*Employees, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadEmployeeByFirstName not implemented")
}
func (UnimplementedEmployeeServiceServer) ReadEmployeeByLastName(context.Context, *EmployeeLastName) (*Employees, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadEmployeeByLastName not implemented")
}
func (UnimplementedEmployeeServiceServer) ReadEmployeeByDisplayName(context.Context, *EmployeeDisplayName) (*Employees, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadEmployeeByDisplayName not implemented")
}
func (UnimplementedEmployeeServiceServer) ReadEmployeeByAge(context.Context, *EmployeeAge) (*Employees, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadEmployeeByAge not implemented")
}
func (UnimplementedEmployeeServiceServer) ReadEmployeeBySalary(context.Context, *EmployeeSalary) (*Employees, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadEmployeeBySalary not implemented")
}
func (UnimplementedEmployeeServiceServer) ReadEmployeeByDesignation(context.Context, *EmployeeDesignation) (*Employees, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadEmployeeByDesignation not implemented")
}
func (UnimplementedEmployeeServiceServer) ReadEmployeeByDepartment(context.Context, *EmployeeDepartment) (*Employees, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadEmployeeByDepartment not implemented")
}
func (UnimplementedEmployeeServiceServer) UpdateEmployee(context.Context, *Employee) (*UpdateEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmployee not implemented")
}
func (UnimplementedEmployeeServiceServer) DeleteEmployee(context.Context, *EmployeeId) (*DeleteEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmployee not implemented")
}
func (UnimplementedEmployeeServiceServer) mustEmbedUnimplementedEmployeeServiceServer() {}

// UnsafeEmployeeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmployeeServiceServer will
// result in compilation errors.
type UnsafeEmployeeServiceServer interface {
	mustEmbedUnimplementedEmployeeServiceServer()
}

func RegisterEmployeeServiceServer(s grpc.ServiceRegistrar, srv EmployeeServiceServer) {
	s.RegisterService(&EmployeeService_ServiceDesc, srv)
}

func _EmployeeService_CreateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Employee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).CreateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmployeeService/CreateEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).CreateEmployee(ctx, req.(*Employee))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_ReadEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Employee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).ReadEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmployeeService/ReadEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).ReadEmployee(ctx, req.(*Employee))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_ReadEmployeeById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).ReadEmployeeById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmployeeService/ReadEmployeeById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).ReadEmployeeById(ctx, req.(*EmployeeId))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_ReadEmployeeByFirstName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeFirstName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).ReadEmployeeByFirstName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmployeeService/ReadEmployeeByFirstName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).ReadEmployeeByFirstName(ctx, req.(*EmployeeFirstName))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_ReadEmployeeByLastName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeLastName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).ReadEmployeeByLastName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmployeeService/ReadEmployeeByLastName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).ReadEmployeeByLastName(ctx, req.(*EmployeeLastName))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_ReadEmployeeByDisplayName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeDisplayName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).ReadEmployeeByDisplayName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmployeeService/ReadEmployeeByDisplayName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).ReadEmployeeByDisplayName(ctx, req.(*EmployeeDisplayName))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_ReadEmployeeByAge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeAge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).ReadEmployeeByAge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmployeeService/ReadEmployeeByAge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).ReadEmployeeByAge(ctx, req.(*EmployeeAge))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_ReadEmployeeBySalary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeSalary)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).ReadEmployeeBySalary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmployeeService/ReadEmployeeBySalary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).ReadEmployeeBySalary(ctx, req.(*EmployeeSalary))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_ReadEmployeeByDesignation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeDesignation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).ReadEmployeeByDesignation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmployeeService/ReadEmployeeByDesignation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).ReadEmployeeByDesignation(ctx, req.(*EmployeeDesignation))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_ReadEmployeeByDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeDepartment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).ReadEmployeeByDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmployeeService/ReadEmployeeByDepartment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).ReadEmployeeByDepartment(ctx, req.(*EmployeeDepartment))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_UpdateEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Employee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).UpdateEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmployeeService/UpdateEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).UpdateEmployee(ctx, req.(*Employee))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_DeleteEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).DeleteEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EmployeeService/DeleteEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).DeleteEmployee(ctx, req.(*EmployeeId))
	}
	return interceptor(ctx, in, info, handler)
}

// EmployeeService_ServiceDesc is the grpc.ServiceDesc for EmployeeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmployeeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EmployeeService",
	HandlerType: (*EmployeeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEmployee",
			Handler:    _EmployeeService_CreateEmployee_Handler,
		},
		{
			MethodName: "ReadEmployee",
			Handler:    _EmployeeService_ReadEmployee_Handler,
		},
		{
			MethodName: "ReadEmployeeById",
			Handler:    _EmployeeService_ReadEmployeeById_Handler,
		},
		{
			MethodName: "ReadEmployeeByFirstName",
			Handler:    _EmployeeService_ReadEmployeeByFirstName_Handler,
		},
		{
			MethodName: "ReadEmployeeByLastName",
			Handler:    _EmployeeService_ReadEmployeeByLastName_Handler,
		},
		{
			MethodName: "ReadEmployeeByDisplayName",
			Handler:    _EmployeeService_ReadEmployeeByDisplayName_Handler,
		},
		{
			MethodName: "ReadEmployeeByAge",
			Handler:    _EmployeeService_ReadEmployeeByAge_Handler,
		},
		{
			MethodName: "ReadEmployeeBySalary",
			Handler:    _EmployeeService_ReadEmployeeBySalary_Handler,
		},
		{
			MethodName: "ReadEmployeeByDesignation",
			Handler:    _EmployeeService_ReadEmployeeByDesignation_Handler,
		},
		{
			MethodName: "ReadEmployeeByDepartment",
			Handler:    _EmployeeService_ReadEmployeeByDepartment_Handler,
		},
		{
			MethodName: "UpdateEmployee",
			Handler:    _EmployeeService_UpdateEmployee_Handler,
		},
		{
			MethodName: "DeleteEmployee",
			Handler:    _EmployeeService_DeleteEmployee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "employee.proto",
}
