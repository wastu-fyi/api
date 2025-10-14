package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"wastu/app/models"
	"wastu/pkg/resp"
)

func SanctumAuth() http.Middleware {
	return func(ctx http.Context) {
		facades.Log().Debug("[SanctumAuth] enter")

		getHeader := func(k string) string { return strings.TrimSpace(ctx.Request().Header(k)) }
		authz := getHeader("Authorization")
		if authz == "" {
			authz = getHeader("HTTP_AUTHORIZATION")
		}
		if authz == "" {
			authz = getHeader("X-Forwarded-Authorization")
		}

		partsAuth := strings.Fields(authz)
		if len(partsAuth) != 2 || !strings.EqualFold(partsAuth[0], "Bearer") {
			ctx.Response().Header("WWW-Authenticate", `Bearer realm="sanctum"`)
			ctx.Response().Json(401, resp.Body{StatusCode: 401, Code: "UNAUTHORIZED", Message: "Missing bearer token"}).Abort()
			return
		}

		raw := partsAuth[1]
		parts := strings.SplitN(raw, "|", 2)
		if len(parts) != 2 {
			ctx.Response().Header("WWW-Authenticate", `Bearer realm="sanctum"`)
			ctx.Response().Json(401, resp.Body{StatusCode: 401, Code: "UNAUTHORIZED", Message: "Bad token format"}).Abort()
			return
		}

		idPart := parts[0]
		secret := parts[1]

		idNum, err := strconv.ParseUint(idPart, 10, 64)
		if err != nil {
			ctx.Response().Header("WWW-Authenticate", `Bearer realm="sanctum"`)
			ctx.Response().Json(401, resp.Body{StatusCode: 401, Code: "UNAUTHORIZED", Message: "Bad token format"}).Abort()
			return
		}

		h := sha256.Sum256([]byte(secret))
		hexHash := hex.EncodeToString(h[:])

		db := facades.Orm().Query()
		var pat models.PersonalAccessToken
		if err := db.Model(&models.PersonalAccessToken{}).
			Where("id = ? AND token = ?", idNum, hexHash).
			Where("expires_at IS NULL OR expires_at > ?", time.Now()).
			First(&pat); err != nil {
			ctx.Response().Header("WWW-Authenticate", `Bearer realm="sanctum"`)
			ctx.Response().Json(401, resp.Body{StatusCode: 401, Code: "UNAUTHORIZED", Message: "Invalid or expired token"}).Abort()
			return
		}

		ctx.WithValue("auth.user_id", pat.TokenableID)
		ctx.WithValue("auth.token_id", pat.ID)

		go func(id uint64) {
			_, _ = facades.Orm().Query().
				Model(&models.PersonalAccessToken{}).
				Where("id = ?", id).
				Update("last_used_at", time.Now())
		}(pat.ID)

		ctx.Request().Next()
	}
}
