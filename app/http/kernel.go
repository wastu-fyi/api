package http

import (
	"github.com/goravel/framework/contracts/http"

	"wastu/app/http/middleware"
)

type Kernel struct {
}

func (kernel Kernel) Middleware() []http.Middleware {
	return []http.Middleware{
		middleware.RecoverJSON(),
	}
}
