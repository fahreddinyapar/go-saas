// Code generated by protoc-gen-go-grpc-proxy. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc-proxy v1.2.0
// - protoc             (unknown)
// source: user/api/auth/v1/auth.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

var _ AuthServer = (*authClientProxy)(nil)

// authClientProxy is the proxy to turn Auth client to server interface.
//
type authClientProxy struct {
	cc AuthClient
}

func NewAuthClientProxy(cc AuthClient) AuthServer {
	return &authClientProxy{cc}
}

func (c *authClientProxy) Register(ctx context.Context, in *RegisterAuthRequest) (*RegisterAuthReply, error) {
	return c.cc.Register(ctx, in)
}
func (c *authClientProxy) Login(ctx context.Context, in *LoginAuthRequest) (*LoginAuthReply, error) {
	return c.cc.Login(ctx, in)
}
func (c *authClientProxy) GetLogin(ctx context.Context, in *GetLoginRequest) (*GetLoginResponse, error) {
	return c.cc.GetLogin(ctx, in)
}
func (c *authClientProxy) Token(ctx context.Context, in *TokenRequest) (*TokenReply, error) {
	return c.cc.Token(ctx, in)
}
func (c *authClientProxy) Refresh(ctx context.Context, in *RefreshTokenAuthRequest) (*RefreshTokenAuthReply, error) {
	return c.cc.Refresh(ctx, in)
}
func (c *authClientProxy) SendPasswordlessToken(ctx context.Context, in *PasswordlessTokenAuthRequest) (*PasswordlessTokenAuthReply, error) {
	return c.cc.SendPasswordlessToken(ctx, in)
}
func (c *authClientProxy) LoginPasswordless(ctx context.Context, in *LoginPasswordlessRequest) (*LoginPasswordlessReply, error) {
	return c.cc.LoginPasswordless(ctx, in)
}
func (c *authClientProxy) SendForgetPasswordToken(ctx context.Context, in *ForgetPasswordTokenRequest) (*ForgetPasswordTokenReply, error) {
	return c.cc.SendForgetPasswordToken(ctx, in)
}
func (c *authClientProxy) ForgetPassword(ctx context.Context, in *ForgetPasswordRequest) (*ForgetPasswordReply, error) {
	return c.cc.ForgetPassword(ctx, in)
}
func (c *authClientProxy) ChangePasswordByForget(ctx context.Context, in *ChangePasswordByForgetRequest) (*ChangePasswordByForgetReply, error) {
	return c.cc.ChangePasswordByForget(ctx, in)
}
func (c *authClientProxy) ValidatePassword(ctx context.Context, in *ValidatePasswordRequest) (*ValidatePasswordReply, error) {
	return c.cc.ValidatePassword(ctx, in)
}
func (c *authClientProxy) ChangePasswordByPre(ctx context.Context, in *ChangePasswordByPreRequest) (*ChangePasswordByPreReply, error) {
	return c.cc.ChangePasswordByPre(ctx, in)
}
func (c *authClientProxy) GetCsrfToken(ctx context.Context, in *GetCsrfTokenRequest) (*GetCsrfTokenResponse, error) {
	return c.cc.GetCsrfToken(ctx, in)
}
func (c *authClientProxy) RefreshRememberToken(ctx context.Context, in *RefreshRememberTokenRequest) (*RefreshRememberTokenReply, error) {
	return c.cc.RefreshRememberToken(ctx, in)
}
