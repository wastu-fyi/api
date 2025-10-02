package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"wastu/app/http/controllers"
	"wastu/pkg/resp"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users/{id}", userController.Show)

	facades.Route().Fallback(func(ctx http.Context) http.Response {
		return resp.NotFound(ctx, "The requested resource could not be found")
	})

}
