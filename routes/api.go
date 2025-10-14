package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"wastu/app/http/controllers"
	"wastu/app/http/middleware"
	"wastu/pkg/resp"
)

func Api() {
	studentController := controllers.NewStudentController()

	facades.Route().Middleware(middleware.SanctumAuth()).Group(func(router route.Router) {
		router.Get("/students", studentController.Index)
		router.Get("/students/studies", studentController.Studies)
		router.Get("/students/years", studentController.Years)
		router.Get("/students/detail", studentController.Show)
	})

	facades.Route().Fallback(func(ctx http.Context) http.Response {
		return resp.NotFound(ctx, "The requested resource could not be found")
	})
}
