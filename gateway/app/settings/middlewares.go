package settings

import (
	"common_go/middlewares"
)

var CommonMiddlewares = []middlewares.MiddlewareFunc{
	middlewares.LoggerMiddleware,
	middlewares.ErrorHandlerMiddleware,
}
