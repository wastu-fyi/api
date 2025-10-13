package config

import (
	"github.com/goravel/framework/contracts/database/driver"
	"github.com/goravel/framework/facades"
	mysqlfacades "github.com/goravel/mysql/facades"
)

func init() {
	config := facades.Config()
	config.Add("database", map[string]any{
		"default": config.Env("DB_CONNECTION", "mysql"),
		"connections": map[string]any{
			"mysql": map[string]any{
				"host":     config.Env("DB_HOST", "127.0.0.1"),
				"port":     config.Env("DB_PORT", 3306),
				"database": config.Env("DB_DATABASE", "forge"),
				"username": config.Env("DB_USERNAME", ""),
				"password": config.Env("DB_PASSWORD", ""),
				"charset":  "utf8mb4",
				"prefix":   "",
				"singular": false,
				"via": func() (driver.Driver, error) {
					return mysqlfacades.Mysql("mysql")
				},
			},
		},
		"pool": map[string]any{
			"max_idle_conns":    10,
			"max_open_conns":    100,
			"conn_max_idletime": 3600,
			"conn_max_lifetime": 3600,
		},
		"slow_threshold": 200,
		"migrations": map[string]any{
			"table": "migrations",
		},
		"redis": map[string]any{
			"default": map[string]any{
				"host":     config.Env("REDIS_HOST", ""),
				"password": config.Env("REDIS_PASSWORD", ""),
				"port":     config.Env("REDIS_PORT", 6379),
				"database": config.Env("REDIS_DB", 0),
			},
		},
	})
}
