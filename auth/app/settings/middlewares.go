package settings

import (
	"auth/app/middlewares"
)

var commonMiddlewares = []middlewares.MiddlewareFunc{
	LoggingMiddleware,
}
