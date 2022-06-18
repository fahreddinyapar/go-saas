package dal

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/goxiaoy/go-eventbus"
	"github.com/goxiaoy/go-saas-kit/pkg/blob"
	kitconf "github.com/goxiaoy/go-saas-kit/pkg/conf"
	"github.com/goxiaoy/go-saas-kit/pkg/email"
	event2 "github.com/goxiaoy/go-saas-kit/pkg/event"
	"github.com/goxiaoy/go-saas-kit/pkg/event/event"
	kitgorm "github.com/goxiaoy/go-saas-kit/pkg/gorm"
	kitredis "github.com/goxiaoy/go-saas-kit/pkg/redis"
	kituow "github.com/goxiaoy/go-saas-kit/pkg/uow"
	"github.com/goxiaoy/go-saas/common"
	"github.com/goxiaoy/go-saas/data"
	sgorm "github.com/goxiaoy/go-saas/gorm"
	"github.com/goxiaoy/uow"
)

type (
	ConnName             string
	ConstConnStrResolver data.ConnStrResolver
	ConstDbProvider      sgorm.DbProvider
)

var (
	//DefaultProviderSet shared provider for all data layer
	DefaultProviderSet = wire.NewSet(
		NewConnStrResolver,
		NewConstantConnStrResolver,
		kitgorm.NewDbOpener,
		NewGormConfig,

		kitgorm.NewDbProvider,
		NewConstDbProvider,

		wire.Value(UowCfg),
		kituow.NewUowManager,

		NewBlobFactory,

		NewRedis,
		kitredis.NewCache,

		NewEmailer,
		NewEventSender,
		NewRemoteEventReceiver,
		wire.Value(eventbus.Default),
	)
)

var (
	UowCfg = &uow.Config{
		SupportNestedTransaction: false,
	}
)

func NewConnStrResolver(c *kitconf.Data, ts common.TenantStore) data.ConnStrResolver {
	return kitgorm.NewConnStrResolver(c.Endpoints, ts)
}

// NewConstantConnStrResolver ignore multi-tenancy
func NewConstantConnStrResolver(c *kitconf.Data) ConstConnStrResolver {
	// ignore multi-tenancy
	conn := make(data.ConnStrings, 1)
	for k, v := range c.Endpoints.Databases {
		conn[k] = v.Source
	}
	return data.NewDefaultConnStrResolver(data.NewConnStrOption(conn))
}

// NewConstDbProvider  ignore multi-tenancy
func NewConstDbProvider(cs ConstConnStrResolver, c *sgorm.Config, opener sgorm.DbOpener) ConstDbProvider {
	return sgorm.NewDefaultDbProvider(cs, c, opener)
}

func NewBlobFactory(c *kitconf.Data) blob.Factory {
	return blob.NewFactory(c.Blobs)
}

func NewGormConfig(d *kitconf.Data, name ConnName) *sgorm.Config {
	return kitgorm.NewGormConfig(d, string(name))
}

func NewRedis(c *kitconf.Data, name ConnName) *redis.Client {
	if c.Endpoints.Redis == nil {
		panic("redis endpoints required")
	}
	r := kitredis.ResolveRedisEndpointOrDefault(c.Endpoints.Redis, string(name))
	return kitredis.NewRedisClient(r)
}

func NewEmailer(c *kitconf.Data) email.LazyClient {
	return email.NewLazyClient(c.Endpoints)
}

func NewEventSender(c *kitconf.Data, logger log.Logger, name ConnName) (event.Sender, func(), error) {
	e := c.Endpoints.GetEventMergedDefault(string(name))
	return event2.NewEventSender(e, logger)
}

func NewRemoteEventReceiver(c *kitconf.Data, logger log.Logger, handler event.Handler, name ConnName) (event.Receiver, func(), error) {
	e := c.Endpoints.GetEventMergedDefault(string(name))
	return event2.NewEventReceiver(e, handler, logger)
}