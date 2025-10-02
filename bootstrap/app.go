package bootstrap

import (
	"github.com/goravel/framework/foundation"

	"wastu/config"
)

func Boot() {
	app := foundation.NewApplication()

	app.Boot()
	config.Boot()
}
