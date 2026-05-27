package controllers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"wastu/pkg/resp"
)

type UptimeController struct{}

func NewUptimeController() *UptimeController {
	return &UptimeController{}
}

func (c *UptimeController) Index(ctx http.Context) http.Response {
	token := facades.Config().Env("UPTIME_TOKEN", "")
	if token == nil || token == "" {
		return resp.InternalServerError(ctx, "UPTIME_TOKEN is not configured.")
	}

	id := facades.Config().Env("UPTIME_ID", "")
	if id == nil || id == "" {
		return resp.InternalServerError(ctx, "UPTIME_ID is not configured.")
	}

	url := fmt.Sprintf("https://uptime.betterstack.com/api/v2/status-pages/%v", id)
	response, err := facades.Http().WithToken(token.(string)).Get(url)
	if err != nil {
		return resp.InternalServerError(ctx, "Gagal menghubungi layanan BetterStack.", resp.WithMessage(err.Error()))
	}

	body, bodyErr := response.Body()
	if bodyErr != nil {
		return resp.InternalServerError(ctx, "Gagal membaca respon dari BetterStack.", resp.WithMessage(bodyErr.Error()))
	}

	if !response.Successful() {
		return resp.InternalServerError(ctx, "Layanan BetterStack memberikan respon kesalahan.", resp.WithMessage(body))
	}

	var result struct {
		Data struct {
			Attributes struct {
				CompanyName    string `json:"company_name"`
				AggregateState string `json:"aggregate_state"`
			} `json:"attributes"`
		} `json:"data"`
	}

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return resp.InternalServerError(ctx, "Gagal memproses data dari BetterStack.", resp.WithMessage(err.Error()))
	}

	data := map[string]string{
		"name":   result.Data.Attributes.CompanyName,
		"status": strings.ToLower(result.Data.Attributes.AggregateState),
	}

	ctx.Response().Header("Cache-Control", "public, max-age=60")

	return resp.OK(ctx, data, "Status uptime berhasil dimuat.")
}
