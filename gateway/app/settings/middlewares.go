package settings

import (
	"common_go/middlewares"
)

var commonMiddlewares = []middlewares.MiddlewareFunc{
	middlewares.ErrorHandlerMiddleware,
	middlewares.ErrorHandlerMiddleware,
	middlewares.LoggerMiddleware,
}
