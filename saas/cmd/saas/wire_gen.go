// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/goxiaoy/go-saas-kit/auth/jwt"
	"github.com/goxiaoy/go-saas-kit/authorization/authorization"
	"github.com/goxiaoy/go-saas-kit/pkg/api"
	"github.com/goxiaoy/go-saas-kit/pkg/conf"
	uow2 "github.com/goxiaoy/go-saas-kit/pkg/uow"
	"github.com/goxiaoy/go-saas-kit/saas/internal_/biz"
	conf2 "github.com/goxiaoy/go-saas-kit/saas/internal_/conf"
	"github.com/goxiaoy/go-saas-kit/saas/internal_/data"
	"github.com/goxiaoy/go-saas-kit/saas/internal_/server"
	"github.com/goxiaoy/go-saas-kit/saas/internal_/service"
	api2 "github.com/goxiaoy/go-saas-kit/user/api"
	"github.com/goxiaoy/go-saas/common/http"
	"github.com/goxiaoy/go-saas/gorm"
	"github.com/goxiaoy/uow"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(services *conf.Services, confData *conf2.Data, logger log.Logger, tokenizerConfig *jwt.TokenizerConfig, config *uow.Config, gormConfig *gorm.Config, webMultiTenancyOption *http.WebMultiTenancyOption, arg ...grpc.ClientOption) (*kratos.App, func(), error) {
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
	userServiceClient := api2.NewUserGrpcClient(grpcConn)
	remoteRoleContributor := api2.NewRemoteRoleContributor(userServiceClient)
	authorizationOption := service.NewAuthorizationOption(remoteRoleContributor)
	permissionService := authorization.NewPermissionService(logger)
	defaultAuthorizationService := authorization.NewDefaultAuthorizationService(authorizationOption, permissionService, logger)
	tenantService := service.NewTenantService(tenantUseCase, defaultAuthorizationService)
	httpServer := server.NewHTTPServer(services, tokenizer, tenantStore, manager, tenantService, webMultiTenancyOption, option, logger)
	grpcServer := server.NewGRPCServer(services, tokenizer, tenantStore, manager, tenantService, webMultiTenancyOption, option, logger)
	dbProvider := data.NewProvider(confData, gormConfig, dbOpener, manager, tenantStore, logger)
	dataData, cleanup3, err := data.NewData(confData, dbProvider, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	migrate := data.NewMigrate(dataData)
	roleServiceClient := api2.NewRoleGrpcClient(grpcConn)
	permissionSeeder := biz.NewPermissionSeeder(permissionService, roleServiceClient)
	seeder := server.NewSeeder(confData, manager, migrate, permissionSeeder)
	app := newApp(logger, httpServer, grpcServer, seeder)
	return app, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

var (
	_wireClientNameValue = server.ClientName
)
