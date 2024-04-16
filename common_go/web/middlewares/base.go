package middlewares

import (
	"common_go/web/handlers"
)

type MiddlewareFunc func(handlers.Handler) handlers.Handler
