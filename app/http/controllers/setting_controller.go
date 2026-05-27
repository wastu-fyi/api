package controllers

import (
	"encoding/json"
	"strings"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"wastu/app/models"
	"wastu/pkg/resp"
)

type SettingController struct{}

func NewSettingController() *SettingController {
	return &SettingController{}
}

func (c *SettingController) Index(ctx http.Context) http.Response {
	var settings []models.Setting
	if err := facades.Orm().Query().Model(&models.Setting{}).Where("`group` = ?", "app").Get(&settings); err != nil {
		return resp.InternalServerError(ctx, "Gagal mengambil pengaturan aplikasi.", resp.WithMessage("select query error: "+err.Error()))
	}

	data := make(map[string]interface{})
	for _, s := range settings {
		var val interface{}
		if len(s.Payload) > 0 {
			if err := json.Unmarshal(s.Payload, &val); err == nil {
				// Handle double-encoded JSON (strings that contain JSON)
				if str, ok := val.(string); ok {
					trimmed := strings.TrimSpace(str)
					if (strings.HasPrefix(trimmed, "{") && strings.HasSuffix(trimmed, "}")) ||
						(strings.HasPrefix(trimmed, "[") && strings.HasSuffix(trimmed, "]")) {
						var nested interface{}
						if err := json.Unmarshal([]byte(str), &nested); err == nil {
							val = nested
						}
					}
				}
			} else {
				val = string(s.Payload)
			}
		}
		data[s.Name] = val
	}

	return resp.OK(ctx, data, "Pengaturan aplikasi berhasil dimuat.")
}
