package settings

import "auth/middlewares"

var commonMiddlewares = []middlewares.MiddlewareFunc{
	LoggingMiddleware,
}
