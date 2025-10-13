package config

import (
	"github.com/goravel/framework/auth"
	"github.com/goravel/framework/cache"
	"github.com/goravel/framework/console"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/crypt"
	"github.com/goravel/framework/database"
	"github.com/goravel/framework/event"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/filesystem"
	"github.com/goravel/framework/grpc"
	"github.com/goravel/framework/hash"
	"github.com/goravel/framework/http"
	"github.com/goravel/framework/log"
	"github.com/goravel/framework/mail"
	"github.com/goravel/framework/queue"
	"github.com/goravel/framework/route"
	"github.com/goravel/framework/schedule"
	"github.com/goravel/framework/session"
	"github.com/goravel/framework/support/carbon"
	"github.com/goravel/framework/testing"
	"github.com/goravel/framework/translation"
	"github.com/goravel/framework/validation"
	"github.com/goravel/gin"
	"github.com/goravel/mysql"
	"github.com/goravel/redis"

	"wastu/app/providers"
)

func Boot() {}

func init() {
	config := facades.Config()
	config.Add("app", map[string]any{
		"name":            config.Env("APP_NAME", "Wastu.FYI"),
		"env":             config.Env("APP_ENV", "production"),
		"debug":           config.Env("APP_DEBUG", false),
		"timezone":        carbon.Singapore,
		"locale":          "id",
		"fallback_locale": "en",
		"lang_path":       "lang",
		"key":             config.Env("APP_KEY", ""),
		"providers": []foundation.ServiceProvider{
			&log.ServiceProvider{},
			&console.ServiceProvider{},
			&mysql.ServiceProvider{},
			&database.ServiceProvider{},
			&redis.ServiceProvider{},
			&cache.ServiceProvider{},
			&http.ServiceProvider{},
			&route.ServiceProvider{},
			&schedule.ServiceProvider{},
			&event.ServiceProvider{},
			&queue.ServiceProvider{},
			&grpc.ServiceProvider{},
			&mail.ServiceProvider{},
			&auth.ServiceProvider{},
			&hash.ServiceProvider{},
			&crypt.ServiceProvider{},
			&filesystem.ServiceProvider{},
			&validation.ServiceProvider{},
			&session.ServiceProvider{},
			&translation.ServiceProvider{},
			&testing.ServiceProvider{},
			&providers.AppServiceProvider{},
			&providers.AuthServiceProvider{},
			&providers.RouteServiceProvider{},
			&providers.GrpcServiceProvider{},
			&providers.ConsoleServiceProvider{},
			&providers.QueueServiceProvider{},
			&providers.EventServiceProvider{},
			&providers.ValidationServiceProvider{},
			&providers.DatabaseServiceProvider{},
			&gin.ServiceProvider{},
		},
	})
}
