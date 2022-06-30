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
	"github.com/go-saas/kit/pkg/registry"
	"github.com/go-saas/kit/pkg/server"
	"github.com/go-saas/kit/pkg/uow"
	api2 "github.com/go-saas/kit/saas/api"
	"github.com/go-saas/kit/saas/private/biz"
	conf2 "github.com/go-saas/kit/saas/private/conf"
	"github.com/go-saas/kit/saas/private/data"
	server2 "github.com/go-saas/kit/saas/private/server"
	"github.com/go-saas/kit/saas/private/service"
	api3 "github.com/go-saas/kit/user/api"
	"github.com/goxiaoy/go-eventbus"
)

import (
	_ "github.com/go-saas/kit/event/kafka"
	_ "github.com/go-saas/kit/event/pulsar"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(services *conf.Services, security *conf.Security, confData *conf.Data, saasConf *conf2.SaasConf, logger log.Logger, appConfig *conf.AppConfig, arg ...grpc.ClientOption) (*kratos.App, func(), error) {
	tokenizerConfig := jwt.NewTokenizerConfig(security)
	tokenizer := jwt.NewTokenizer(tokenizerConfig)
	trustedContextValidator := api.NewClientTrustedContextValidator()
	eventBus := _wireEventBusValue
	dbCache, cleanup := gorm.NewDbCache(confData, logger)
	connStrings := dal.NewConstantConnStrResolver(confData)
	constDbProvider := dal.NewConstDbProvider(dbCache, connStrings, confData)
	dataData, cleanup2, err := data.NewData(confData, constDbProvider, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	tenantRepo := data.NewTenantRepo(eventBus, dataData)
	connStrGenerator := biz.NewConfigConnStrGenerator(saasConf)
	connName := _wireConnNameValue
	producer, cleanup3, err := dal.NewEventSender(confData, connName)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tenantUseCase := biz.NewTenantUserCase(tenantRepo, connStrGenerator, producer)
	factory := dal.NewBlobFactory(confData)
	tenantInternalService := &service.TenantInternalService{
		Trusted: trustedContextValidator,
		UseCase: tenantUseCase,
		App:     appConfig,
		Blob:    factory,
	}
	tenantStore := api2.NewTenantStore(tenantInternalService)
	manager := uow.NewUowManager(dbCache)
	webMultiTenancyOption := server.NewWebMultiTenancyOption(appConfig)
	option := api.NewDefaultOption(logger)
	decodeRequestFunc := _wireDecodeRequestFuncValue
	encodeResponseFunc := _wireEncodeResponseFuncValue
	encodeErrorFunc := _wireEncodeErrorFuncValue
	clientName := _wireClientNameValue
	registryConf := registry.NewConf(services)
	inMemoryTokenManager := api.NewInMemoryTokenManager(tokenizer, logger)
	grpcConn, cleanup4 := api3.NewGrpcConn(clientName, services, registryConf, option, inMemoryTokenManager, logger, arg...)
	userServiceServer := api3.NewUserGrpcClient(grpcConn)
	userTenantContrib := api3.NewUserTenantContrib(userServiceServer)
	permissionServiceServer := api3.NewPermissionGrpcClient(grpcConn)
	permissionChecker := api3.NewRemotePermissionChecker(permissionServiceServer)
	authzOption := server2.NewAuthorizationOption()
	subjectResolverImpl := authz.NewSubjectResolver(authzOption)
	defaultAuthorizationService := authz.NewDefaultAuthorizationService(permissionChecker, subjectResolverImpl, logger)
	tenantService := service.NewTenantService(tenantUseCase, defaultAuthorizationService, trustedContextValidator, factory, appConfig)
	httpServerRegister := service.NewHttpServerRegister(tenantService, factory, defaultAuthorizationService, encodeErrorFunc, tenantInternalService, confData)
	httpServer := server2.NewHTTPServer(services, security, tokenizer, tenantStore, manager, webMultiTenancyOption, option, decodeRequestFunc, encodeResponseFunc, encodeErrorFunc, logger, trustedContextValidator, userTenantContrib, httpServerRegister)
	grpcServerRegister := service.NewGrpcServerRegister(tenantService, tenantInternalService)
	grpcServer := server2.NewGRPCServer(services, tokenizer, tenantStore, manager, webMultiTenancyOption, option, userTenantContrib, trustedContextValidator, grpcServerRegister, logger)
	client, err := dal.NewRedis(confData, connName)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	redisConnOpt := job.NewAsynqClientOpt(client)
	jobServer := server2.NewJobServer(redisConnOpt, logger)
	tenantReadyEventHandler := biz.NewTenantReadyEventHandler(tenantUseCase)
	consumerFactoryServer := server2.NewEventServer(confData, connName, logger, manager, tenantReadyEventHandler)
	migrate := data.NewMigrate(dataData)
	seeding := server2.NewSeeding(manager, migrate)
	seeder := server2.NewSeeder(tenantStore, seeding)
	app := newApp(logger, httpServer, grpcServer, jobServer, consumerFactoryServer, seeder)
	return app, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

var (
	_wireEventBusValue           = eventbus.Default
	_wireConnNameValue           = biz.ConnName
	_wireDecodeRequestFuncValue  = server.ReqDecode
	_wireEncodeResponseFuncValue = server.ResEncoder
	_wireEncodeErrorFuncValue    = server.ErrEncoder
	_wireClientNameValue         = server2.ClientName
)
