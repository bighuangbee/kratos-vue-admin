// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: api.proto

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
	Api_ListApi_FullMethodName                = "/api.admin.v1.Api/ListApi"
	Api_AllApi_FullMethodName                 = "/api.admin.v1.Api/AllApi"
	Api_CreateApi_FullMethodName              = "/api.admin.v1.Api/CreateApi"
	Api_UpdateApi_FullMethodName              = "/api.admin.v1.Api/UpdateApi"
	Api_GetPolicyPathByRoleKey_FullMethodName = "/api.admin.v1.Api/GetPolicyPathByRoleKey"
	Api_GetApi_FullMethodName                 = "/api.admin.v1.Api/GetApi"
	Api_DeleteApi_FullMethodName              = "/api.admin.v1.Api/DeleteApi"
)

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	// api列表
	ListApi(ctx context.Context, in *ListApiRequest, opts ...grpc.CallOption) (*ListApiReply, error)
	// 获取所有api
	AllApi(ctx context.Context, in *AllApiRequest, opts ...grpc.CallOption) (*AllApiReply, error)
	// 创建api
	CreateApi(ctx context.Context, in *CreateApiRequest, opts ...grpc.CallOption) (*CreateApiReply, error)
	// 更新api
	UpdateApi(ctx context.Context, in *UpdateApiRequest, opts ...grpc.CallOption) (*UpdateApiReply, error)
	// 获取角色拥有的api权限
	GetPolicyPathByRoleKey(ctx context.Context, in *GetPolicyPathByRoleKeyRequest, opts ...grpc.CallOption) (*GetPolicyPathByRoleKeyReply, error)
	// 获取api
	GetApi(ctx context.Context, in *GetApiRequest, opts ...grpc.CallOption) (*GetApiReply, error)
	// 删除api
	DeleteApi(ctx context.Context, in *DeleteApiRequest, opts ...grpc.CallOption) (*DeleteApiReply, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) ListApi(ctx context.Context, in *ListApiRequest, opts ...grpc.CallOption) (*ListApiReply, error) {
	out := new(ListApiReply)
	err := c.cc.Invoke(ctx, Api_ListApi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) AllApi(ctx context.Context, in *AllApiRequest, opts ...grpc.CallOption) (*AllApiReply, error) {
	out := new(AllApiReply)
	err := c.cc.Invoke(ctx, Api_AllApi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) CreateApi(ctx context.Context, in *CreateApiRequest, opts ...grpc.CallOption) (*CreateApiReply, error) {
	out := new(CreateApiReply)
	err := c.cc.Invoke(ctx, Api_CreateApi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) UpdateApi(ctx context.Context, in *UpdateApiRequest, opts ...grpc.CallOption) (*UpdateApiReply, error) {
	out := new(UpdateApiReply)
	err := c.cc.Invoke(ctx, Api_UpdateApi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetPolicyPathByRoleKey(ctx context.Context, in *GetPolicyPathByRoleKeyRequest, opts ...grpc.CallOption) (*GetPolicyPathByRoleKeyReply, error) {
	out := new(GetPolicyPathByRoleKeyReply)
	err := c.cc.Invoke(ctx, Api_GetPolicyPathByRoleKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetApi(ctx context.Context, in *GetApiRequest, opts ...grpc.CallOption) (*GetApiReply, error) {
	out := new(GetApiReply)
	err := c.cc.Invoke(ctx, Api_GetApi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) DeleteApi(ctx context.Context, in *DeleteApiRequest, opts ...grpc.CallOption) (*DeleteApiReply, error) {
	out := new(DeleteApiReply)
	err := c.cc.Invoke(ctx, Api_DeleteApi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations must embed UnimplementedApiServer
// for forward compatibility
type ApiServer interface {
	// api列表
	ListApi(context.Context, *ListApiRequest) (*ListApiReply, error)
	// 获取所有api
	AllApi(context.Context, *AllApiRequest) (*AllApiReply, error)
	// 创建api
	CreateApi(context.Context, *CreateApiRequest) (*CreateApiReply, error)
	// 更新api
	UpdateApi(context.Context, *UpdateApiRequest) (*UpdateApiReply, error)
	// 获取角色拥有的api权限
	GetPolicyPathByRoleKey(context.Context, *GetPolicyPathByRoleKeyRequest) (*GetPolicyPathByRoleKeyReply, error)
	// 获取api
	GetApi(context.Context, *GetApiRequest) (*GetApiReply, error)
	// 删除api
	DeleteApi(context.Context, *DeleteApiRequest) (*DeleteApiReply, error)
	mustEmbedUnimplementedApiServer()
}

// UnimplementedApiServer must be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (UnimplementedApiServer) ListApi(context.Context, *ListApiRequest) (*ListApiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApi not implemented")
}
func (UnimplementedApiServer) AllApi(context.Context, *AllApiRequest) (*AllApiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllApi not implemented")
}
func (UnimplementedApiServer) CreateApi(context.Context, *CreateApiRequest) (*CreateApiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApi not implemented")
}
func (UnimplementedApiServer) UpdateApi(context.Context, *UpdateApiRequest) (*UpdateApiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApi not implemented")
}
func (UnimplementedApiServer) GetPolicyPathByRoleKey(context.Context, *GetPolicyPathByRoleKeyRequest) (*GetPolicyPathByRoleKeyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicyPathByRoleKey not implemented")
}
func (UnimplementedApiServer) GetApi(context.Context, *GetApiRequest) (*GetApiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApi not implemented")
}
func (UnimplementedApiServer) DeleteApi(context.Context, *DeleteApiRequest) (*DeleteApiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApi not implemented")
}
func (UnimplementedApiServer) mustEmbedUnimplementedApiServer() {}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_ListApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).ListApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_ListApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).ListApi(ctx, req.(*ListApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_AllApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).AllApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_AllApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).AllApi(ctx, req.(*AllApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_CreateApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).CreateApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_CreateApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).CreateApi(ctx, req.(*CreateApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_UpdateApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).UpdateApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_UpdateApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).UpdateApi(ctx, req.(*UpdateApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetPolicyPathByRoleKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPolicyPathByRoleKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetPolicyPathByRoleKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_GetPolicyPathByRoleKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetPolicyPathByRoleKey(ctx, req.(*GetPolicyPathByRoleKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_GetApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetApi(ctx, req.(*GetApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_DeleteApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).DeleteApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_DeleteApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).DeleteApi(ctx, req.(*DeleteApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.admin.v1.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListApi",
			Handler:    _Api_ListApi_Handler,
		},
		{
			MethodName: "AllApi",
			Handler:    _Api_AllApi_Handler,
		},
		{
			MethodName: "CreateApi",
			Handler:    _Api_CreateApi_Handler,
		},
		{
			MethodName: "UpdateApi",
			Handler:    _Api_UpdateApi_Handler,
		},
		{
			MethodName: "GetPolicyPathByRoleKey",
			Handler:    _Api_GetPolicyPathByRoleKey_Handler,
		},
		{
			MethodName: "GetApi",
			Handler:    _Api_GetApi_Handler,
		},
		{
			MethodName: "DeleteApi",
			Handler:    _Api_DeleteApi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}