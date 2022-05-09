// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.2

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

type PermissionServiceHTTPServer interface {
	AddSubjectPermission(context.Context, *AddSubjectPermissionRequest) (*AddSubjectPermissionResponse, error)
	CheckCurrent(context.Context, *CheckPermissionRequest) (*CheckPermissionReply, error)
	GetCurrent(context.Context, *GetCurrentPermissionRequest) (*GetCurrentPermissionReply, error)
	ListSubjectPermission(context.Context, *ListSubjectPermissionRequest) (*ListSubjectPermissionResponse, error)
	RemoveSubjectPermission(context.Context, *RemoveSubjectPermissionRequest) (*RemoveSubjectPermissionReply, error)
	UpdateSubjectPermission(context.Context, *UpdateSubjectPermissionRequest) (*UpdateSubjectPermissionResponse, error)
}

func RegisterPermissionServiceHTTPServer(s *http.Server, srv PermissionServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/permission/current", _PermissionService_GetCurrent0_HTTP_Handler(srv))
	r.POST("/v1/permission/check", _PermissionService_CheckCurrent0_HTTP_Handler(srv))
	r.POST("/v1/permission/subject", _PermissionService_AddSubjectPermission0_HTTP_Handler(srv))
	r.POST("/v1/permission/subject/list", _PermissionService_ListSubjectPermission0_HTTP_Handler(srv))
	r.GET("/v1/permission/subject", _PermissionService_ListSubjectPermission1_HTTP_Handler(srv))
	r.PUT("/v1/permission/subject", _PermissionService_UpdateSubjectPermission0_HTTP_Handler(srv))
	r.POST("/v1/permission/subject/rm", _PermissionService_RemoveSubjectPermission0_HTTP_Handler(srv))
}

func _PermissionService_GetCurrent0_HTTP_Handler(srv PermissionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCurrentPermissionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.permission.v1.PermissionService/GetCurrent")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCurrent(ctx, req.(*GetCurrentPermissionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCurrentPermissionReply)
		return ctx.Result(200, reply)
	}
}

func _PermissionService_CheckCurrent0_HTTP_Handler(srv PermissionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CheckPermissionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.permission.v1.PermissionService/CheckCurrent")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CheckCurrent(ctx, req.(*CheckPermissionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CheckPermissionReply)
		return ctx.Result(200, reply)
	}
}

func _PermissionService_AddSubjectPermission0_HTTP_Handler(srv PermissionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddSubjectPermissionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.permission.v1.PermissionService/AddSubjectPermission")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddSubjectPermission(ctx, req.(*AddSubjectPermissionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddSubjectPermissionResponse)
		return ctx.Result(200, reply)
	}
}

func _PermissionService_ListSubjectPermission0_HTTP_Handler(srv PermissionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListSubjectPermissionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.permission.v1.PermissionService/ListSubjectPermission")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListSubjectPermission(ctx, req.(*ListSubjectPermissionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListSubjectPermissionResponse)
		return ctx.Result(200, reply)
	}
}

func _PermissionService_ListSubjectPermission1_HTTP_Handler(srv PermissionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListSubjectPermissionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.permission.v1.PermissionService/ListSubjectPermission")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListSubjectPermission(ctx, req.(*ListSubjectPermissionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListSubjectPermissionResponse)
		return ctx.Result(200, reply)
	}
}

func _PermissionService_UpdateSubjectPermission0_HTTP_Handler(srv PermissionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateSubjectPermissionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.permission.v1.PermissionService/UpdateSubjectPermission")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateSubjectPermission(ctx, req.(*UpdateSubjectPermissionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateSubjectPermissionResponse)
		return ctx.Result(200, reply)
	}
}

func _PermissionService_RemoveSubjectPermission0_HTTP_Handler(srv PermissionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RemoveSubjectPermissionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.permission.v1.PermissionService/RemoveSubjectPermission")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RemoveSubjectPermission(ctx, req.(*RemoveSubjectPermissionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RemoveSubjectPermissionReply)
		return ctx.Result(200, reply)
	}
}

type PermissionServiceHTTPClient interface {
	AddSubjectPermission(ctx context.Context, req *AddSubjectPermissionRequest, opts ...http.CallOption) (rsp *AddSubjectPermissionResponse, err error)
	CheckCurrent(ctx context.Context, req *CheckPermissionRequest, opts ...http.CallOption) (rsp *CheckPermissionReply, err error)
	GetCurrent(ctx context.Context, req *GetCurrentPermissionRequest, opts ...http.CallOption) (rsp *GetCurrentPermissionReply, err error)
	ListSubjectPermission(ctx context.Context, req *ListSubjectPermissionRequest, opts ...http.CallOption) (rsp *ListSubjectPermissionResponse, err error)
	RemoveSubjectPermission(ctx context.Context, req *RemoveSubjectPermissionRequest, opts ...http.CallOption) (rsp *RemoveSubjectPermissionReply, err error)
	UpdateSubjectPermission(ctx context.Context, req *UpdateSubjectPermissionRequest, opts ...http.CallOption) (rsp *UpdateSubjectPermissionResponse, err error)
}

type PermissionServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewPermissionServiceHTTPClient(client *http.Client) PermissionServiceHTTPClient {
	return &PermissionServiceHTTPClientImpl{client}
}

func (c *PermissionServiceHTTPClientImpl) AddSubjectPermission(ctx context.Context, in *AddSubjectPermissionRequest, opts ...http.CallOption) (*AddSubjectPermissionResponse, error) {
	var out AddSubjectPermissionResponse
	pattern := "/v1/permission/subject"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.api.permission.v1.PermissionService/AddSubjectPermission"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PermissionServiceHTTPClientImpl) CheckCurrent(ctx context.Context, in *CheckPermissionRequest, opts ...http.CallOption) (*CheckPermissionReply, error) {
	var out CheckPermissionReply
	pattern := "/v1/permission/check"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.api.permission.v1.PermissionService/CheckCurrent"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PermissionServiceHTTPClientImpl) GetCurrent(ctx context.Context, in *GetCurrentPermissionRequest, opts ...http.CallOption) (*GetCurrentPermissionReply, error) {
	var out GetCurrentPermissionReply
	pattern := "/v1/permission/current"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.api.permission.v1.PermissionService/GetCurrent"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PermissionServiceHTTPClientImpl) ListSubjectPermission(ctx context.Context, in *ListSubjectPermissionRequest, opts ...http.CallOption) (*ListSubjectPermissionResponse, error) {
	var out ListSubjectPermissionResponse
	pattern := "/v1/permission/subject"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.api.permission.v1.PermissionService/ListSubjectPermission"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PermissionServiceHTTPClientImpl) RemoveSubjectPermission(ctx context.Context, in *RemoveSubjectPermissionRequest, opts ...http.CallOption) (*RemoveSubjectPermissionReply, error) {
	var out RemoveSubjectPermissionReply
	pattern := "/v1/permission/subject/rm"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.api.permission.v1.PermissionService/RemoveSubjectPermission"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PermissionServiceHTTPClientImpl) UpdateSubjectPermission(ctx context.Context, in *UpdateSubjectPermissionRequest, opts ...http.CallOption) (*UpdateSubjectPermissionResponse, error) {
	var out UpdateSubjectPermissionResponse
	pattern := "/v1/permission/subject"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.api.permission.v1.PermissionService/UpdateSubjectPermission"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
