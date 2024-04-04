package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := httptest.NewRecorder()
		next.ServeHTTP(recorder, r)
		statusCode := recorder.Result().StatusCode
		log.Printf("Request: %s %s | Status: %d\n", r.Method, r.URL.RequestURI(), statusCode)
		fmt.Println(recorder.Header())
		err := recorder.Result().Write(w)
		if err != nil {
			return
		}
	})
}

func ErrorHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Возникла ошибка:", r)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		// Вызываем следующий обработчик в цепочке
		next.ServeHTTP(w, r)
	})
}

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("session")
		if err != nil {
			if !strings.HasPrefix(r.URL.Path, "/auth/") && !strings.HasPrefix(r.URL.Path, "/handbook/") {
				currentURL := r.URL.Path
				redirectURL := "/auth/login?next=" + url.QueryEscape(currentURL)
				http.Redirect(w, r, redirectURL, http.StatusFound)
				return
			}
		}
		// Вызываем следующий обработчик в цепочке
		next.ServeHTTP(w, r)
	})
}
