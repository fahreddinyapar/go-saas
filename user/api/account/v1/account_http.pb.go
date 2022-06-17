// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

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

const OperationAccountCreateAddresses = "/user.api.account.v1.Account/CreateAddresses"
const OperationAccountDeleteAddresses = "/user.api.account.v1.Account/DeleteAddresses"
const OperationAccountGetAddresses = "/user.api.account.v1.Account/GetAddresses"
const OperationAccountGetProfile = "/user.api.account.v1.Account/GetProfile"
const OperationAccountGetSettings = "/user.api.account.v1.Account/GetSettings"
const OperationAccountUpdateAddresses = "/user.api.account.v1.Account/UpdateAddresses"
const OperationAccountUpdateProfile = "/user.api.account.v1.Account/UpdateProfile"
const OperationAccountUpdateSettings = "/user.api.account.v1.Account/UpdateSettings"

type AccountHTTPServer interface {
	CreateAddresses(context.Context, *CreateAddressesRequest) (*CreateAddressReply, error)
	DeleteAddresses(context.Context, *DeleteAddressRequest) (*DeleteAddressesReply, error)
	GetAddresses(context.Context, *GetAddressesRequest) (*GetAddressesReply, error)
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error)
	GetSettings(context.Context, *GetSettingsRequest) (*GetSettingsResponse, error)
	UpdateAddresses(context.Context, *UpdateAddressesRequest) (*UpdateAddressesReply, error)
	UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error)
	UpdateSettings(context.Context, *UpdateSettingsRequest) (*UpdateSettingsResponse, error)
}

func RegisterAccountHTTPServer(s *http.Server, srv AccountHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/account/profile", _Account_GetProfile0_HTTP_Handler(srv))
	r.PUT("/v1/account/profile", _Account_UpdateProfile0_HTTP_Handler(srv))
	r.GET("/v1/account/settings", _Account_GetSettings0_HTTP_Handler(srv))
	r.PUT("/v1/account/settings", _Account_UpdateSettings0_HTTP_Handler(srv))
	r.GET("/v1/account/addresses", _Account_GetAddresses0_HTTP_Handler(srv))
	r.POST("/v1/account/addresses", _Account_CreateAddresses0_HTTP_Handler(srv))
	r.PUT("/v1/account/address/{address.id}", _Account_UpdateAddresses0_HTTP_Handler(srv))
	r.DELETE("/v1/account/addresses", _Account_DeleteAddresses0_HTTP_Handler(srv))
}

func _Account_GetProfile0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetProfileRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountGetProfile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProfile(ctx, req.(*GetProfileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetProfileResponse)
		return ctx.Result(200, reply)
	}
}

func _Account_UpdateProfile0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateProfileRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountUpdateProfile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateProfile(ctx, req.(*UpdateProfileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateProfileResponse)
		return ctx.Result(200, reply)
	}
}

func _Account_GetSettings0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSettingsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountGetSettings)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSettings(ctx, req.(*GetSettingsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSettingsResponse)
		return ctx.Result(200, reply)
	}
}

func _Account_UpdateSettings0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateSettingsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountUpdateSettings)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateSettings(ctx, req.(*UpdateSettingsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateSettingsResponse)
		return ctx.Result(200, reply)
	}
}

func _Account_GetAddresses0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetAddressesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountGetAddresses)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAddresses(ctx, req.(*GetAddressesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAddressesReply)
		return ctx.Result(200, reply)
	}
}

func _Account_CreateAddresses0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateAddressesRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountCreateAddresses)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAddresses(ctx, req.(*CreateAddressesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateAddressReply)
		return ctx.Result(200, reply)
	}
}

func _Account_UpdateAddresses0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateAddressesRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountUpdateAddresses)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateAddresses(ctx, req.(*UpdateAddressesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateAddressesReply)
		return ctx.Result(200, reply)
	}
}

func _Account_DeleteAddresses0_HTTP_Handler(srv AccountHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteAddressRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountDeleteAddresses)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteAddresses(ctx, req.(*DeleteAddressRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteAddressesReply)
		return ctx.Result(200, reply)
	}
}

type AccountHTTPClient interface {
	CreateAddresses(ctx context.Context, req *CreateAddressesRequest, opts ...http.CallOption) (rsp *CreateAddressReply, err error)
	DeleteAddresses(ctx context.Context, req *DeleteAddressRequest, opts ...http.CallOption) (rsp *DeleteAddressesReply, err error)
	GetAddresses(ctx context.Context, req *GetAddressesRequest, opts ...http.CallOption) (rsp *GetAddressesReply, err error)
	GetProfile(ctx context.Context, req *GetProfileRequest, opts ...http.CallOption) (rsp *GetProfileResponse, err error)
	GetSettings(ctx context.Context, req *GetSettingsRequest, opts ...http.CallOption) (rsp *GetSettingsResponse, err error)
	UpdateAddresses(ctx context.Context, req *UpdateAddressesRequest, opts ...http.CallOption) (rsp *UpdateAddressesReply, err error)
	UpdateProfile(ctx context.Context, req *UpdateProfileRequest, opts ...http.CallOption) (rsp *UpdateProfileResponse, err error)
	UpdateSettings(ctx context.Context, req *UpdateSettingsRequest, opts ...http.CallOption) (rsp *UpdateSettingsResponse, err error)
}

type AccountHTTPClientImpl struct {
	cc *http.Client
}

func NewAccountHTTPClient(client *http.Client) AccountHTTPClient {
	return &AccountHTTPClientImpl{client}
}

func (c *AccountHTTPClientImpl) CreateAddresses(ctx context.Context, in *CreateAddressesRequest, opts ...http.CallOption) (*CreateAddressReply, error) {
	var out CreateAddressReply
	pattern := "/v1/account/addresses"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountCreateAddresses))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) DeleteAddresses(ctx context.Context, in *DeleteAddressRequest, opts ...http.CallOption) (*DeleteAddressesReply, error) {
	var out DeleteAddressesReply
	pattern := "/v1/account/addresses"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAccountDeleteAddresses))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) GetAddresses(ctx context.Context, in *GetAddressesRequest, opts ...http.CallOption) (*GetAddressesReply, error) {
	var out GetAddressesReply
	pattern := "/v1/account/addresses"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAccountGetAddresses))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...http.CallOption) (*GetProfileResponse, error) {
	var out GetProfileResponse
	pattern := "/v1/account/profile"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAccountGetProfile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) GetSettings(ctx context.Context, in *GetSettingsRequest, opts ...http.CallOption) (*GetSettingsResponse, error) {
	var out GetSettingsResponse
	pattern := "/v1/account/settings"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAccountGetSettings))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) UpdateAddresses(ctx context.Context, in *UpdateAddressesRequest, opts ...http.CallOption) (*UpdateAddressesReply, error) {
	var out UpdateAddressesReply
	pattern := "/v1/account/address/{address.id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountUpdateAddresses))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...http.CallOption) (*UpdateProfileResponse, error) {
	var out UpdateProfileResponse
	pattern := "/v1/account/profile"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountUpdateProfile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AccountHTTPClientImpl) UpdateSettings(ctx context.Context, in *UpdateSettingsRequest, opts ...http.CallOption) (*UpdateSettingsResponse, error) {
	var out UpdateSettingsResponse
	pattern := "/v1/account/settings"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountUpdateSettings))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
