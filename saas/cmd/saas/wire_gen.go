// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/goxiaoy/go-saas-kit/pkg/api"
	"github.com/goxiaoy/go-saas-kit/pkg/authn/jwt"
	"github.com/goxiaoy/go-saas-kit/pkg/authz/authorization"
	"github.com/goxiaoy/go-saas-kit/pkg/conf"
	server2 "github.com/goxiaoy/go-saas-kit/pkg/server"
	uow2 "github.com/goxiaoy/go-saas-kit/pkg/uow"
	"github.com/goxiaoy/go-saas-kit/saas/private/biz"
	conf2 "github.com/goxiaoy/go-saas-kit/saas/private/conf"
	"github.com/goxiaoy/go-saas-kit/saas/private/data"
	"github.com/goxiaoy/go-saas-kit/saas/private/server"
	"github.com/goxiaoy/go-saas-kit/saas/private/service"
	api2 "github.com/goxiaoy/go-saas-kit/user/api"
	"github.com/goxiaoy/go-saas-kit/user/remote"
	"github.com/goxiaoy/go-saas/common/http"
	"github.com/goxiaoy/go-saas/gorm"
	"github.com/goxiaoy/uow"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(services *conf.Services, security *conf.Security, confData *conf2.Data, logger log.Logger, config *uow.Config, gormConfig *gorm.Config, webMultiTenancyOption *http.WebMultiTenancyOption, arg ...grpc.ClientOption) (*kratos.App, func(), error) {
	tokenizerConfig := jwt.NewTokenizerConfig(security)
	tokenizer := jwt.NewTokenizer(tokenizerConfig)
	tenantRepo := data.NewTenantRepo()
	tenantStore := data.NewTenantStore(tenantRepo)
	dbOpener, cleanup := gorm.NewDbOpener()
	manager := uow2.NewUowManager(gormConfig, config, dbOpener)
	tenantUseCase := biz.NewTenantUserCase(tenantRepo)
	clientName := _wireClientNameValue
	saasContributor := api.NewSaasContributor(webMultiTenancyOption)
	userContributor := api.NewUserContributor()
	option := api.NewDefaultOption(saasContributor, userContributor)
	inMemoryTokenManager := api.NewInMemoryTokenManager(tokenizer)
	grpcConn, cleanup2 := api2.NewGrpcConn(clientName, services, option, inMemoryTokenManager, arg...)
	permissionServiceClient := api2.NewPermissionGrpcClient(grpcConn)
	permissionChecker := remote.NewRemotePermissionChecker(permissionServiceClient)
	authorizationOption := service.NewAuthorizationOption()
	subjectResolverImpl := authorization.NewSubjectResolver(authorizationOption)
	defaultAuthorizationService := authorization.NewDefaultAuthorizationService(permissionChecker, subjectResolverImpl, logger)
	factory := data.NewBlobFactory(confData)
	tenantService := service.NewTenantService(tenantUseCase, defaultAuthorizationService, factory)
	decodeRequestFunc := _wireDecodeRequestFuncValue
	encodeResponseFunc := _wireEncodeResponseFuncValue
	encodeErrorFunc := _wireEncodeErrorFuncValue
	httpServer := server.NewHTTPServer(services, security, tokenizer, tenantStore, manager, tenantService, webMultiTenancyOption, option, decodeRequestFunc, encodeResponseFunc, encodeErrorFunc, factory, confData, logger)
	grpcServer := server.NewGRPCServer(services, tokenizer, tenantStore, manager, tenantService, webMultiTenancyOption, option, logger)
	dbProvider := data.NewProvider(confData, gormConfig, dbOpener, tenantStore, logger)
	dataData, cleanup3, err := data.NewData(confData, dbProvider, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	migrate := data.NewMigrate(dataData)
	seeder := server.NewSeeder(confData, manager, migrate)
	app := newApp(logger, httpServer, grpcServer, seeder)
	return app, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

var (
	_wireClientNameValue         = server.ClientName
	_wireDecodeRequestFuncValue  = server2.ReqDecode
	_wireEncodeResponseFuncValue = server2.ResEncoder
	_wireEncodeErrorFuncValue    = server2.ErrEncoder
)
