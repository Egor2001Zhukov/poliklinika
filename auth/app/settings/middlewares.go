package settings

import (
	"common_go/web/middlewares"
)

var commonMiddlewares = []middlewares.MiddlewareFunc{
	middlewares.ErrorHandlerMiddleware,
	middlewares.ErrorHandlerMiddleware,
	middlewares.LoggerMiddleware,
}
