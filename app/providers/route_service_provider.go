package providers

import (
	"fmt"
	"strconv"

	"github.com/goravel/framework/contracts/foundation"
	contractshttp "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/http/limit"
	"github.com/goravel/framework/http/middleware"

	"wastu/app/http"
	"wastu/routes"
)

type RouteServiceProvider struct {
}

func (receiver *RouteServiceProvider) Register(app foundation.Application) {
}

func (receiver *RouteServiceProvider) Boot(app foundation.Application) {
	facades.Route().GlobalMiddleware(append(http.Kernel{}.Middleware(), middleware.Throttle("global"))...)
	receiver.configureRateLimiting()

	routes.Api()
}

func (receiver *RouteServiceProvider) configureRateLimiting() {
	facades.RateLimiter().For("global", func(ctx contractshttp.Context) contractshttp.Limit {
		maxAny := facades.Config().Env("RATE_LIMIT_MAX", 60)
		var perMinute int
		switch v := maxAny.(type) {
		case int:
			perMinute = v
		case int64:
			perMinute = int(v)
		case string:
			if n, err := strconv.Atoi(v); err == nil {
				perMinute = n
			}
		}
		if perMinute <= 0 {
			perMinute = 60
		}

		if v := ctx.Value("auth.user_id"); v != nil {
			return limit.PerMinute(perMinute).By(fmt.Sprint(v))
		}
		return limit.PerMinute(perMinute).By(ctx.Request().Ip())
	})
}
