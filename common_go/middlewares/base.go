package middlewares

import "net/http"

type MiddlewareFunc func(handlerFunc http.HandlerFunc) http.HandlerFunc
