package settings

import "common_go/web/middlewares"

var CommonMiddlewares = []middlewares.MiddlewareFunc{
	middlewares.AuthenticationMiddleware(PublicServices, LoginEndpoint, true),
	middlewares.ErrorHandlerMiddleware(isDev), // После должен быть обработчик ошибок
	middlewares.LoggerMiddleware(),            // Первым должен быть логер потому что он подменяет записывающую функцию рекордером
}
