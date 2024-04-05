package settings

import "github.com/zhumorist/common_go/middlewares"

var CommonMiddlewares = []middlewares.MiddlewareFunc{
	middlewares.AuthenticationMiddleware(PublicServices, LoginEndpoint, true),
	middlewares.ErrorHandlerMiddleware(), // После должен быть обработчик ошибок
	middlewares.LoggerMiddleware(),       // Первым должен быть логер потому что он подменяет записывающую функцию рекордером
}
