package controllers

import (
	"wastu/pkg/resp"

	"github.com/goravel/framework/contracts/http"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (r *UserController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Input("id")

	if id == "" {
		return resp.BadRequest(ctx, "id is required",
			map[string]any{"field": "id"},
			resp.WithCode("INVALID_PAYLOAD"),
		)
	}

	user := map[string]any{"id": id, "name": "Suluh"}
	return resp.OK(ctx, user, "User fetched")
}
