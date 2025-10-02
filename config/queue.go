package config

import (
	"github.com/goravel/framework/contracts/queue"
	"github.com/goravel/framework/facades"
	redisfacades "github.com/goravel/redis/facades"
)

func init() {
	config := facades.Config()
	config.Add("queue", map[string]any{
		"default": config.Env("QUEUE_CONNECTION", "redis"),
		"connections": map[string]any{
			"sync": map[string]any{
				"driver": "sync",
			},
			"database": map[string]any{
				"driver":     "database",
				"connection": "postgres",
				"queue":      "default",
				"concurrent": 1,
			},
			"redis": map[string]any{
				"driver":     "custom",
				"connection": "default",
				"queue":      "default",
				"via": func() (queue.Driver, error) {
					return redisfacades.Queue("redis")
				},
			},
		},
		"failed": map[string]any{
			"database": config.Env("DB_CONNECTION", "mysql"),
			"table":    "failed_jobs",
		},
	})
}
