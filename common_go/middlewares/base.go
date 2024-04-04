package middlewares

import "net/http"

type MiddlewareFunc func(handler http.Handler) http.Handler
