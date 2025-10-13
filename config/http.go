package config

import (
	"github.com/gin-gonic/gin/render"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"github.com/goravel/gin"
	ginfacades "github.com/goravel/gin/facades"
)

func init() {
	config := facades.Config()
	config.Add("http", map[string]any{
		"default": "gin",
		"drivers": map[string]any{
			"gin": map[string]any{
				"body_limit":   4096,
				"header_limit": 4096,
				"route": func() (route.Route, error) {
					return ginfacades.Route("gin"), nil
				},
				"template": func() (render.HTMLRender, error) {
					return gin.DefaultTemplate()
				},
			},
		},
		"url":             config.Env("APP_URL", "http://localhost"),
		"host":            config.Env("APP_HOST", "127.0.0.1"),
		"port":            config.Env("APP_PORT", "3000"),
		"request_timeout": 3,
		"tls": map[string]any{
			"host": config.Env("APP_HOST", "127.0.0.1"),
			"port": config.Env("APP_PORT", "3000"),
			"ssl": map[string]any{
				"cert": "",
				"key":  "",
			},
		},
		"client": map[string]any{
			"base_url":                config.GetString("HTTP_CLIENT_BASE_URL"),
			"timeout":                 config.GetDuration("HTTP_CLIENT_TIMEOUT"),
			"max_idle_conns":          config.GetInt("HTTP_CLIENT_MAX_IDLE_CONNS"),
			"max_idle_conns_per_host": config.GetInt("HTTP_CLIENT_MAX_IDLE_CONNS_PER_HOST"),
			"max_conns_per_host":      config.GetInt("HTTP_CLIENT_MAX_CONN_PER_HOST"),
			"idle_conn_timeout":       config.GetDuration("HTTP_CLIENT_IDLE_CONN_TIMEOUT"),
		},
	})
}
