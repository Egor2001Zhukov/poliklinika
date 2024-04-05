package main

import (
	"fmt"
	"net/http"
	"poliklinika_gateway/app/settings"
	"poliklinika_gateway/app/urls"
	"strings"
)

func mainHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		microService := strings.Split(r.URL.Path, "/")[1]
		f, ok := urls.MicroServices[microService]
		if !ok {
			http.Error(w, "ServiceNotFound", http.StatusNotFound)
			return
		}
		f(w, r)
	})
}

func main() {
	handler := mainHandler()
	for _, middleware := range settings.CommonMiddlewares {
		handler = middleware(handler)
	}
	http.Handle("/", handler)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
