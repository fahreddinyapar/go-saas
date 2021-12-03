//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/goxiaoy/go-saas-kit/auth/jwt"
	"github.com/goxiaoy/go-saas-kit/authorization/authorization"
	conf2 "github.com/goxiaoy/go-saas-kit/pkg/conf"
	"github.com/goxiaoy/go-saas-kit/saas/internal_/biz"
	"github.com/goxiaoy/go-saas-kit/saas/internal_/conf"
	"github.com/goxiaoy/go-saas-kit/saas/internal_/data"
	"github.com/goxiaoy/go-saas-kit/saas/internal_/server"
	"github.com/goxiaoy/go-saas-kit/saas/internal_/service"
	"github.com/goxiaoy/go-saas-kit/user/api"
	"github.com/goxiaoy/go-saas/common/http"
	"github.com/goxiaoy/go-saas/gorm"
	"github.com/goxiaoy/uow"
)

// initApp init kratos application.
func initApp(*conf2.Services, *conf.Data, log.Logger, *jwt.TokenizerConfig, *uow.Config, *gorm.Config, *http.WebMultiTenancyOption, ...grpc.ClientOption) (*kratos.App, func(), error) {
	panic(wire.Build(authorization.ProviderSet, server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, api.GrpcProviderSet, newApp))
}
