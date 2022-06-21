package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	sapi "github.com/goxiaoy/go-saas-kit/pkg/api"
	"github.com/goxiaoy/go-saas-kit/pkg/authn/jwt"
	"github.com/goxiaoy/go-saas-kit/pkg/conf"
	"github.com/goxiaoy/go-saas-kit/pkg/localize"
	"github.com/goxiaoy/go-saas-kit/pkg/server"
	"github.com/goxiaoy/go-saas-kit/pkg/uow"
	"github.com/goxiaoy/go-saas-kit/saas/api"
	"github.com/goxiaoy/go-saas-kit/saas/i18n"
	"github.com/goxiaoy/go-saas-kit/saas/private/service"
	uapi "github.com/goxiaoy/go-saas-kit/user/api"
	"github.com/goxiaoy/go-saas/common"
	http2 "github.com/goxiaoy/go-saas/common/http"
	uow2 "github.com/goxiaoy/uow"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Services,
	sCfg *conf.Security,
	tokenizer jwt.Tokenizer,
	ts common.TenantStore,
	uowMgr uow2.Manager,
	mOpt *http2.WebMultiTenancyOption,
	apiOpt *sapi.Option,
	reqDecoder http.DecodeRequestFunc,
	resEncoder http.EncodeResponseFunc,
	errEncoder http.EncodeErrorFunc,
	logger log.Logger,
	validator sapi.TrustedContextValidator,
	userTenant *uapi.UserTenantContributor,
	register service.HttpServerRegister) *http.Server {
	var opts []http.ServerOption
	opts = server.PatchHttpOpts(logger, opts, api.ServiceName, c, sCfg, reqDecoder, resEncoder, errEncoder)
	m := middleware.Chain(
		server.Recovery(),
		tracing.Server(),
		logging.Server(logger),
		server.Stack(),
		metrics.Server(),
		validate.Validator(),
		localize.I18N(i18n.Files...),
		jwt.ServerExtractAndAuth(tokenizer, logger),
		sapi.ServerPropagation(apiOpt, validator, logger),
		server.Saas(mOpt, ts, validator, func(o *common.TenantResolveOption) {
			o.AppendContributors(userTenant)
		}),
		uow.Uow(logger, uowMgr),
	)
	opts = append(opts, []http.ServerOption{
		http.Middleware(
			m,
		),
	}...)
	srv := http.NewServer(opts...)

	register.Register(srv, m)

	return srv
}
