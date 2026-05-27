package http

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/gin"

	localmiddleware "wastu/app/http/middleware"
)

type Kernel struct {
}

func (kernel Kernel) Middleware() []http.Middleware {
	return []http.Middleware{
		localmiddleware.RecoverJSON(),
		gin.Cors(),
	}
}
