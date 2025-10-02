package middleware

import (
	"fmt"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"wastu/pkg/resp"
)

func RecoverJSON() http.Middleware {
	return func(ctx http.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				facades.Log().Error(fmt.Sprintf("panic recovered: %v", rec))

				resp.InternalServerError(ctx, "Unexpected server error",
					resp.WithMessage(fmt.Sprint(rec)),
				)

				return
			}
		}()

		ctx.Request().Next()
	}
}
