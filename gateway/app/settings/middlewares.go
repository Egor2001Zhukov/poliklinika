package settings

import (
	"common_go/middlewares"
)

var CommonMiddlewares = []middlewares.MiddlewareFunc{
	middlewares.AuthenticationMiddleware,
	middlewares.ErrorHandlerMiddleware, // После должен быть обработчик ошибок
	//middlewares.LoggerMiddleware,       // Первым должен быть логер потому что он подменяет записывающую функцию рекордером
}
