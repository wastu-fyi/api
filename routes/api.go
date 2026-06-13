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
	settingController := controllers.NewSettingController()
	uptimeController := controllers.NewUptimeController()

	facades.Route().Get("/settings", settingController.Index)
	facades.Route().Get("/uptime", uptimeController.Index)
	facades.Route().Get("/ping", uptimeController.Ping)

	facades.Route().Middleware(middleware.SanctumAuth()).Group(func(router route.Router) {
		router.Get("/students", studentController.Index)
		router.Get("/students/studies", studentController.Studies)
		router.Get("/students/years", studentController.Years)
		router.Get("/students/detail", studentController.Show)
	})

	facades.Route().Fallback(func(ctx http.Context) http.Response {
		return resp.NotFound(ctx, "Rute atau sumber daya yang Anda minta tidak ditemukan. Periksa URL dan coba lagi.")
	})
}
