package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Логирование запроса
		fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)

		// Передача управления следующему обработчику в цепочке
		next.ServeHTTP(w, r)
	}
}

func ErrorHandlerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Возникла ошибка:", r)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		// Вызываем следующий обработчик в цепочке
		next.ServeHTTP(w, r)
	}
}
