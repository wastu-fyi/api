package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"wastu/app/http/controllers"
	"wastu/pkg/resp"
)

func Api() {
	studentController := controllers.NewStudentController()

	facades.Route().Get("/students", studentController.Index)
	facades.Route().Get("/students/studies", studentController.Studies)
	facades.Route().Get("/students/years", studentController.Years)
	facades.Route().Get("/students/:student_id", studentController.Show)

	facades.Route().Fallback(func(ctx http.Context) http.Response {
		return resp.NotFound(ctx, "The requested resource could not be found")
	})
}
