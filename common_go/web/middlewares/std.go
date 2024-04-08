package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func LoggerMiddleware() MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			fmt.Println(w.Header())
			log.Printf("Request: %s %s | Status: ", r.Method, r.URL.RequestURI())
		})
	}
}

func ErrorHandlerMiddleware() MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		{
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
	}
}

func AuthenticationMiddleware(publicEndpoints []string, redirectEndpoint string, withNext bool) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := r.Cookie("session")
			if err != nil {
				var isPublic bool
				for _, service := range publicEndpoints {
					if strings.HasPrefix(r.URL.Path, "/"+service) {
						isPublic = true
						break
					}
				}
				if !isPublic {
					var nextEndpoint string
					if withNext {
						currentURL := r.URL.Path
						nextEndpoint = "?next=" + url.QueryEscape(currentURL)
					}

					redirectURL := "/" + redirectEndpoint + "/" + nextEndpoint
					http.Redirect(w, r, redirectURL, http.StatusFound)
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}
