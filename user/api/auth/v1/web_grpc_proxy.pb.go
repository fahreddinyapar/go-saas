// Code generated by protoc-gen-go-grpc-proxy. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc-proxy v1.2.0
// - protoc             (unknown)
// source: user/api/auth/v1/web.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

var _ AuthWebServer = (*authWebClientProxy)(nil)

// authWebClientProxy is the proxy to turn AuthWeb client to server interface.
//
type authWebClientProxy struct {
	cc AuthWebClient
}

func NewAuthWebClientProxy(cc AuthWebClient) AuthWebServer {
	return &authWebClientProxy{cc}
}

func (c *authWebClientProxy) GetWebLogin(ctx context.Context, in *GetLoginRequest) (*GetLoginResponse, error) {
	return c.cc.GetWebLogin(ctx, in)
}
func (c *authWebClientProxy) WebLogin(ctx context.Context, in *WebLoginAuthRequest) (*WebLoginAuthReply, error) {
	return c.cc.WebLogin(ctx, in)
}
func (c *authWebClientProxy) GetWebLogout(ctx context.Context, in *GetLogoutRequest) (*GetLogoutResponse, error) {
	return c.cc.GetWebLogout(ctx, in)
}
func (c *authWebClientProxy) WebLogout(ctx context.Context, in *LogoutRequest) (*LogoutResponse, error) {
	return c.cc.WebLogout(ctx, in)
}
func (c *authWebClientProxy) GetConsent(ctx context.Context, in *GetConsentRequest) (*GetConsentResponse, error) {
	return c.cc.GetConsent(ctx, in)
}
func (c *authWebClientProxy) GrantConsent(ctx context.Context, in *GrantConsentRequest) (*GrantConsentResponse, error) {
	return c.cc.GrantConsent(ctx, in)
}
