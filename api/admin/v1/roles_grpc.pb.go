// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: roles.proto

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

const (
	Roles_CreateRoles_FullMethodName      = "/api.admin.v1.Roles/CreateRoles"
	Roles_UpdateRoles_FullMethodName      = "/api.admin.v1.Roles/UpdateRoles"
	Roles_ListRoles_FullMethodName        = "/api.admin.v1.Roles/ListRoles"
	Roles_ChangeRoleStatus_FullMethodName = "/api.admin.v1.Roles/ChangeRoleStatus"
	Roles_DataScope_FullMethodName        = "/api.admin.v1.Roles/DataScope"
	Roles_DeleteRoles_FullMethodName      = "/api.admin.v1.Roles/DeleteRoles"
	Roles_GetRoles_FullMethodName         = "/api.admin.v1.Roles/GetRoles"
)

// RolesClient is the client API for Roles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RolesClient interface {
	// 创建角色
	CreateRoles(ctx context.Context, in *CreateRolesRequest, opts ...grpc.CallOption) (*CreateRolesReply, error)
	// 更新角色
	UpdateRoles(ctx context.Context, in *UpdateRolesRequest, opts ...grpc.CallOption) (*UpdateRolesReply, error)
	// 角色列表
	ListRoles(ctx context.Context, in *ListRolesRequest, opts ...grpc.CallOption) (*ListRolesReply, error)
	// 更改角色状态
	ChangeRoleStatus(ctx context.Context, in *ChangeRoleStatusRequest, opts ...grpc.CallOption) (*ChangeRoleStatusReply, error)
	// 更改角色数据范围
	DataScope(ctx context.Context, in *DataScopeRequest, opts ...grpc.CallOption) (*DataScopeReply, error)
	// 删除角色
	DeleteRoles(ctx context.Context, in *DeleteRolesRequest, opts ...grpc.CallOption) (*DeleteRolesReply, error)
	// 获取角色
	GetRoles(ctx context.Context, in *GetRolesRequest, opts ...grpc.CallOption) (*GetRolesReply, error)
}

type rolesClient struct {
	cc grpc.ClientConnInterface
}

func NewRolesClient(cc grpc.ClientConnInterface) RolesClient {
	return &rolesClient{cc}
}

func (c *rolesClient) CreateRoles(ctx context.Context, in *CreateRolesRequest, opts ...grpc.CallOption) (*CreateRolesReply, error) {
	out := new(CreateRolesReply)
	err := c.cc.Invoke(ctx, Roles_CreateRoles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) UpdateRoles(ctx context.Context, in *UpdateRolesRequest, opts ...grpc.CallOption) (*UpdateRolesReply, error) {
	out := new(UpdateRolesReply)
	err := c.cc.Invoke(ctx, Roles_UpdateRoles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) ListRoles(ctx context.Context, in *ListRolesRequest, opts ...grpc.CallOption) (*ListRolesReply, error) {
	out := new(ListRolesReply)
	err := c.cc.Invoke(ctx, Roles_ListRoles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) ChangeRoleStatus(ctx context.Context, in *ChangeRoleStatusRequest, opts ...grpc.CallOption) (*ChangeRoleStatusReply, error) {
	out := new(ChangeRoleStatusReply)
	err := c.cc.Invoke(ctx, Roles_ChangeRoleStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) DataScope(ctx context.Context, in *DataScopeRequest, opts ...grpc.CallOption) (*DataScopeReply, error) {
	out := new(DataScopeReply)
	err := c.cc.Invoke(ctx, Roles_DataScope_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) DeleteRoles(ctx context.Context, in *DeleteRolesRequest, opts ...grpc.CallOption) (*DeleteRolesReply, error) {
	out := new(DeleteRolesReply)
	err := c.cc.Invoke(ctx, Roles_DeleteRoles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) GetRoles(ctx context.Context, in *GetRolesRequest, opts ...grpc.CallOption) (*GetRolesReply, error) {
	out := new(GetRolesReply)
	err := c.cc.Invoke(ctx, Roles_GetRoles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RolesServer is the server API for Roles service.
// All implementations must embed UnimplementedRolesServer
// for forward compatibility
type RolesServer interface {
	// 创建角色
	CreateRoles(context.Context, *CreateRolesRequest) (*CreateRolesReply, error)
	// 更新角色
	UpdateRoles(context.Context, *UpdateRolesRequest) (*UpdateRolesReply, error)
	// 角色列表
	ListRoles(context.Context, *ListRolesRequest) (*ListRolesReply, error)
	// 更改角色状态
	ChangeRoleStatus(context.Context, *ChangeRoleStatusRequest) (*ChangeRoleStatusReply, error)
	// 更改角色数据范围
	DataScope(context.Context, *DataScopeRequest) (*DataScopeReply, error)
	// 删除角色
	DeleteRoles(context.Context, *DeleteRolesRequest) (*DeleteRolesReply, error)
	// 获取角色
	GetRoles(context.Context, *GetRolesRequest) (*GetRolesReply, error)
	mustEmbedUnimplementedRolesServer()
}

// UnimplementedRolesServer must be embedded to have forward compatible implementations.
type UnimplementedRolesServer struct {
}

func (UnimplementedRolesServer) CreateRoles(context.Context, *CreateRolesRequest) (*CreateRolesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoles not implemented")
}
func (UnimplementedRolesServer) UpdateRoles(context.Context, *UpdateRolesRequest) (*UpdateRolesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoles not implemented")
}
func (UnimplementedRolesServer) ListRoles(context.Context, *ListRolesRequest) (*ListRolesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoles not implemented")
}
func (UnimplementedRolesServer) ChangeRoleStatus(context.Context, *ChangeRoleStatusRequest) (*ChangeRoleStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeRoleStatus not implemented")
}
func (UnimplementedRolesServer) DataScope(context.Context, *DataScopeRequest) (*DataScopeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataScope not implemented")
}
func (UnimplementedRolesServer) DeleteRoles(context.Context, *DeleteRolesRequest) (*DeleteRolesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoles not implemented")
}
func (UnimplementedRolesServer) GetRoles(context.Context, *GetRolesRequest) (*GetRolesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoles not implemented")
}
func (UnimplementedRolesServer) mustEmbedUnimplementedRolesServer() {}

// UnsafeRolesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RolesServer will
// result in compilation errors.
type UnsafeRolesServer interface {
	mustEmbedUnimplementedRolesServer()
}

func RegisterRolesServer(s grpc.ServiceRegistrar, srv RolesServer) {
	s.RegisterService(&Roles_ServiceDesc, srv)
}

func _Roles_CreateRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).CreateRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Roles_CreateRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).CreateRoles(ctx, req.(*CreateRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_UpdateRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).UpdateRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Roles_UpdateRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).UpdateRoles(ctx, req.(*UpdateRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_ListRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).ListRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Roles_ListRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).ListRoles(ctx, req.(*ListRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_ChangeRoleStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeRoleStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).ChangeRoleStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Roles_ChangeRoleStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).ChangeRoleStatus(ctx, req.(*ChangeRoleStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_DataScope_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataScopeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).DataScope(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Roles_DataScope_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).DataScope(ctx, req.(*DataScopeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_DeleteRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).DeleteRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Roles_DeleteRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).DeleteRoles(ctx, req.(*DeleteRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_GetRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).GetRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Roles_GetRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).GetRoles(ctx, req.(*GetRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Roles_ServiceDesc is the grpc.ServiceDesc for Roles service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Roles_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.admin.v1.Roles",
	HandlerType: (*RolesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoles",
			Handler:    _Roles_CreateRoles_Handler,
		},
		{
			MethodName: "UpdateRoles",
			Handler:    _Roles_UpdateRoles_Handler,
		},
		{
			MethodName: "ListRoles",
			Handler:    _Roles_ListRoles_Handler,
		},
		{
			MethodName: "ChangeRoleStatus",
			Handler:    _Roles_ChangeRoleStatus_Handler,
		},
		{
			MethodName: "DataScope",
			Handler:    _Roles_DataScope_Handler,
		},
		{
			MethodName: "DeleteRoles",
			Handler:    _Roles_DeleteRoles_Handler,
		},
		{
			MethodName: "GetRoles",
			Handler:    _Roles_GetRoles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "roles.proto",
}