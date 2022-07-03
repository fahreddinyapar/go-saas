// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-saas/kit/pkg/api"
	"github.com/go-saas/kit/pkg/authn/jwt"
	"github.com/go-saas/kit/pkg/authz/authz"
	"github.com/go-saas/kit/pkg/conf"
	"github.com/go-saas/kit/pkg/dal"
	"github.com/go-saas/kit/pkg/gorm"
	"github.com/go-saas/kit/pkg/job"
	server2 "github.com/go-saas/kit/pkg/server"
	"github.com/go-saas/kit/pkg/uow"
	api3 "github.com/go-saas/kit/saas/api"
	"github.com/go-saas/kit/sys/private/biz"
	"github.com/go-saas/kit/sys/private/data"
	"github.com/go-saas/kit/sys/private/server"
	"github.com/go-saas/kit/sys/private/service"
	api2 "github.com/go-saas/kit/user/api"
	"github.com/go-saas/saas/http"
	"github.com/goxiaoy/go-eventbus"
)

import (
	_ "github.com/go-saas/kit/event/kafka"
	_ "github.com/go-saas/kit/event/pulsar"
	_ "github.com/go-saas/kit/pkg/registry/etcd"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(services *conf.Services, security *conf.Security, webMultiTenancyOption *http.WebMultiTenancyOption, confData *conf.Data, logger log.Logger, arg ...grpc.ClientOption) (*kratos.App, func(), error) {
	tokenizerConfig := jwt.NewTokenizerConfig(security)
	tokenizer := jwt.NewTokenizer(tokenizerConfig)
	dbCache, cleanup := gorm.NewDbCache(confData, logger)
	manager := uow.NewUowManager(dbCache)
	decodeRequestFunc := _wireDecodeRequestFuncValue
	encodeResponseFunc := _wireEncodeResponseFuncValue
	encodeErrorFunc := _wireEncodeErrorFuncValue
	option := api.NewDefaultOption(logger)
	trustedContextValidator := api.NewClientTrustedContextValidator()
	clientName := _wireClientNameValue
	discovery, err := api.NewDiscovery(services)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	inMemoryTokenManager := api.NewInMemoryTokenManager(tokenizer, logger)
	grpcConn, cleanup2 := api2.NewGrpcConn(clientName, services, discovery, option, inMemoryTokenManager, logger, arg...)
	permissionServiceServer := api2.NewPermissionGrpcClient(grpcConn)
	permissionChecker := api2.NewRemotePermissionChecker(permissionServiceServer)
	authzOption := server.NewAuthorizationOption()
	subjectResolverImpl := authz.NewSubjectResolver(authzOption)
	defaultAuthorizationService := authz.NewDefaultAuthorizationService(permissionChecker, subjectResolverImpl, logger)
	apiGrpcConn, cleanup3 := api3.NewGrpcConn(clientName, services, discovery, option, inMemoryTokenManager, logger, arg...)
	tenantInternalServiceServer := api3.NewTenantInternalGrpcClient(apiGrpcConn)
	tenantStore := api3.NewTenantStore(tenantInternalServiceServer)
	connStrResolver := dal.NewConnStrResolver(confData, tenantStore)
	dbProvider := gorm.NewDbProvider(dbCache, connStrResolver, confData)
	eventBus := _wireEventBusValue
	menuRepo := data.NewMenuRepo(dbProvider, eventBus)
	menuService := service.NewMenuService(defaultAuthorizationService, menuRepo, logger)
	factory := dal.NewBlobFactory(confData)
	connName := _wireConnNameValue
	client, err := dal.NewRedis(confData, connName)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	redisConnOpt := job.NewAsynqClientOpt(client)
	httpServerRegister := service.NewHttpServerRegister(menuService, defaultAuthorizationService, encodeErrorFunc, factory, confData, redisConnOpt)
	httpServer := server.NewHTTPServer(services, security, tokenizer, manager, decodeRequestFunc, encodeResponseFunc, encodeErrorFunc, option, logger, trustedContextValidator, httpServerRegister)
	grpcServerRegister := service.NewGrpcServerRegister(menuService)
	grpcServer := server.NewGRPCServer(services, tokenizer, manager, option, logger, trustedContextValidator, grpcServerRegister)
	jobServer := server.NewJobServer(redisConnOpt, logger)
	dataData, cleanup4, err := data.NewData(confData, dbProvider, logger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	migrate := data.NewMigrate(dataData)
	menuSeed := biz.NewMenuSeed(dbProvider, menuRepo)
	seeding := server.NewSeeding(manager, migrate, menuSeed)
	seeder := server.NewSeeder(tenantStore, seeding)
	app := newApp(logger, httpServer, grpcServer, jobServer, seeder)
	return app, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

var (
	_wireDecodeRequestFuncValue  = server2.ReqDecode
	_wireEncodeResponseFuncValue = server2.ResEncoder
	_wireEncodeErrorFuncValue    = server2.ErrEncoder
	_wireClientNameValue         = server.ClientName
	_wireEventBusValue           = eventbus.Default
	_wireConnNameValue           = biz.ConnName
)
