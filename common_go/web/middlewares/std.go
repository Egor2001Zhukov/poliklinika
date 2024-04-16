package middlewares

import (
	"common_go/web/handlers"
	"common_go/web/request"
	"common_go/web/response"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func LoggerMiddleware() MiddlewareFunc {
	return func(next handlers.Handler) handlers.Handler {
		return handlers.HandlerFunc(func(req *request.Request, res *response.Response) {
			next.ServeHTTP(req, res)
			log.Printf("Request: %s %s | Status: %d\n", req.Method, req.Path, res.StatusCode)
		})
	}
}

func ErrorHandlerMiddleware(isDev bool) MiddlewareFunc {
	return func(next handlers.Handler) handlers.Handler {
		{
			return handlers.HandlerFunc(func(req *request.Request, res *response.Response) {
				defer func() {
					if r := recover(); r != nil {
						log.Println("Возникла ошибка:", r)
						res.SetStatusCode(http.StatusInternalServerError)
						var errorMessage string
						if isDev {
							errorMessage = fmt.Sprintf("%v", r)
						} else {
							errorMessage = "Internal Server Error"
						}
						res.WriteBody([]byte(errorMessage))
					}
				}()
				// Вызываем следующий обработчик в цепочке
				next.ServeHTTP(req, res)
			})
		}
	}
}

func AuthenticationMiddleware(publicEndpoints []string, redirectEndpoint string, withNext bool) MiddlewareFunc {
	return func(next handlers.Handler) handlers.Handler {
		return handlers.HandlerFunc(func(req *request.Request, res *response.Response) {
			_, err := req.Cookie("session")
			if err != nil {
				var isPublic bool
				for _, service := range publicEndpoints {
					if strings.HasPrefix(req.Path, "/"+service) {
						isPublic = true
						break
					}
				}
				if !isPublic {
					var nextEndpoint string
					if withNext {
						currentURL := req.Path
						nextEndpoint = "?next=" + url.QueryEscape(currentURL)
					}

					redirectURL := "/" + redirectEndpoint + "/" + nextEndpoint
					res.Redirect(redirectURL, http.StatusFound)
					return
				}
			}
			next.ServeHTTP(req, res)
		})
	}
}
