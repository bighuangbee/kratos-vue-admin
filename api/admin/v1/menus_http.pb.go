// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.2
// - protoc             v4.23.2
// source: menus.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationMenusCreateMenus = "/api.admin.v1.Menus/CreateMenus"
const OperationMenusDeleteMenus = "/api.admin.v1.Menus/DeleteMenus"
const OperationMenusGetMenus = "/api.admin.v1.Menus/GetMenus"
const OperationMenusGetMenusTree = "/api.admin.v1.Menus/GetMenusTree"
const OperationMenusListMenus = "/api.admin.v1.Menus/ListMenus"
const OperationMenusRoleMenuTreeSelect = "/api.admin.v1.Menus/RoleMenuTreeSelect"
const OperationMenusUpdateMenus = "/api.admin.v1.Menus/UpdateMenus"

type MenusHTTPServer interface {
	// CreateMenus 创建菜单
	CreateMenus(context.Context, *CreateMenusRequest) (*CreateMenusReply, error)
	// DeleteMenus 删除菜单
	DeleteMenus(context.Context, *DeleteMenusRequest) (*DeleteMenusReply, error)
	// GetMenus 获取菜单
	GetMenus(context.Context, *GetMenusRequest) (*GetMenusReply, error)
	// GetMenusTree 获取菜单关系结构
	GetMenusTree(context.Context, *GetMenusTreeRequest) (*GetMenusTreeReply, error)
	// ListMenus 菜单列表
	ListMenus(context.Context, *ListMenusRequest) (*ListMenusReply, error)
	// RoleMenuTreeSelect 获取角色菜单树
	RoleMenuTreeSelect(context.Context, *RoleMenuTreeSelectRequest) (*RoleMenuTreeSelectReply, error)
	// UpdateMenus 更新菜单
	UpdateMenus(context.Context, *UpdateMenusRequest) (*UpdateMenusReply, error)
}

func RegisterMenusHTTPServer(s *http.Server, srv MenusHTTPServer) {
	r := s.Route("/")
	r.POST("/system/menu", _Menus_CreateMenus0_HTTP_Handler(srv))
	r.GET("/system/menu/list", _Menus_ListMenus0_HTTP_Handler(srv))
	r.GET("/system/menu/menuTreeSelect", _Menus_GetMenusTree0_HTTP_Handler(srv))
	r.PUT("/system/menu", _Menus_UpdateMenus0_HTTP_Handler(srv))
	r.DELETE("/system/menu/{id}", _Menus_DeleteMenus0_HTTP_Handler(srv))
	r.GET("/system/menu/{id}", _Menus_GetMenus0_HTTP_Handler(srv))
	r.GET("/system/menu/roleMenuTreeSelect/{roleId}", _Menus_RoleMenuTreeSelect0_HTTP_Handler(srv))
}

func _Menus_CreateMenus0_HTTP_Handler(srv MenusHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateMenusRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenusCreateMenus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateMenus(ctx, req.(*CreateMenusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateMenusReply)
		return ctx.Result(200, reply)
	}
}

func _Menus_ListMenus0_HTTP_Handler(srv MenusHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMenusRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenusListMenus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMenus(ctx, req.(*ListMenusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMenusReply)
		return ctx.Result(200, reply)
	}
}

func _Menus_GetMenusTree0_HTTP_Handler(srv MenusHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMenusTreeRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenusGetMenusTree)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMenusTree(ctx, req.(*GetMenusTreeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetMenusTreeReply)
		return ctx.Result(200, reply)
	}
}

func _Menus_UpdateMenus0_HTTP_Handler(srv MenusHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateMenusRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenusUpdateMenus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateMenus(ctx, req.(*UpdateMenusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateMenusReply)
		return ctx.Result(200, reply)
	}
}

func _Menus_DeleteMenus0_HTTP_Handler(srv MenusHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteMenusRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenusDeleteMenus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteMenus(ctx, req.(*DeleteMenusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteMenusReply)
		return ctx.Result(200, reply)
	}
}

func _Menus_GetMenus0_HTTP_Handler(srv MenusHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMenusRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenusGetMenus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMenus(ctx, req.(*GetMenusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetMenusReply)
		return ctx.Result(200, reply)
	}
}

func _Menus_RoleMenuTreeSelect0_HTTP_Handler(srv MenusHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RoleMenuTreeSelectRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenusRoleMenuTreeSelect)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoleMenuTreeSelect(ctx, req.(*RoleMenuTreeSelectRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoleMenuTreeSelectReply)
		return ctx.Result(200, reply)
	}
}

type MenusHTTPClient interface {
	CreateMenus(ctx context.Context, req *CreateMenusRequest, opts ...http.CallOption) (rsp *CreateMenusReply, err error)
	DeleteMenus(ctx context.Context, req *DeleteMenusRequest, opts ...http.CallOption) (rsp *DeleteMenusReply, err error)
	GetMenus(ctx context.Context, req *GetMenusRequest, opts ...http.CallOption) (rsp *GetMenusReply, err error)
	GetMenusTree(ctx context.Context, req *GetMenusTreeRequest, opts ...http.CallOption) (rsp *GetMenusTreeReply, err error)
	ListMenus(ctx context.Context, req *ListMenusRequest, opts ...http.CallOption) (rsp *ListMenusReply, err error)
	RoleMenuTreeSelect(ctx context.Context, req *RoleMenuTreeSelectRequest, opts ...http.CallOption) (rsp *RoleMenuTreeSelectReply, err error)
	UpdateMenus(ctx context.Context, req *UpdateMenusRequest, opts ...http.CallOption) (rsp *UpdateMenusReply, err error)
}

type MenusHTTPClientImpl struct {
	cc *http.Client
}

func NewMenusHTTPClient(client *http.Client) MenusHTTPClient {
	return &MenusHTTPClientImpl{client}
}

func (c *MenusHTTPClientImpl) CreateMenus(ctx context.Context, in *CreateMenusRequest, opts ...http.CallOption) (*CreateMenusReply, error) {
	var out CreateMenusReply
	pattern := "/system/menu"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMenusCreateMenus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenusHTTPClientImpl) DeleteMenus(ctx context.Context, in *DeleteMenusRequest, opts ...http.CallOption) (*DeleteMenusReply, error) {
	var out DeleteMenusReply
	pattern := "/system/menu/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMenusDeleteMenus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenusHTTPClientImpl) GetMenus(ctx context.Context, in *GetMenusRequest, opts ...http.CallOption) (*GetMenusReply, error) {
	var out GetMenusReply
	pattern := "/system/menu/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMenusGetMenus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenusHTTPClientImpl) GetMenusTree(ctx context.Context, in *GetMenusTreeRequest, opts ...http.CallOption) (*GetMenusTreeReply, error) {
	var out GetMenusTreeReply
	pattern := "/system/menu/menuTreeSelect"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMenusGetMenusTree))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenusHTTPClientImpl) ListMenus(ctx context.Context, in *ListMenusRequest, opts ...http.CallOption) (*ListMenusReply, error) {
	var out ListMenusReply
	pattern := "/system/menu/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMenusListMenus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenusHTTPClientImpl) RoleMenuTreeSelect(ctx context.Context, in *RoleMenuTreeSelectRequest, opts ...http.CallOption) (*RoleMenuTreeSelectReply, error) {
	var out RoleMenuTreeSelectReply
	pattern := "/system/menu/roleMenuTreeSelect/{roleId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMenusRoleMenuTreeSelect))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenusHTTPClientImpl) UpdateMenus(ctx context.Context, in *UpdateMenusRequest, opts ...http.CallOption) (*UpdateMenusReply, error) {
	var out UpdateMenusReply
	pattern := "/system/menu"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMenusUpdateMenus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}