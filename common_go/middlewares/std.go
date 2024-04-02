package middlewares

import (
	"fmt"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Логирование запроса
		fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)

		// Передача управления следующему обработчику в цепочке
		next.ServeHTTP(w, r)
	})
}
