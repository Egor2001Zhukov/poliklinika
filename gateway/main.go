package main

import (
	"common_go/web/handlers"
	"common_go/web/request"
	"common_go/web/response"
	"fmt"
	"log"
	"net/http"
	"poliklinika_gateway/app/settings"
	"poliklinika_gateway/app/urls"
	"strings"
)

func getResponse(w http.ResponseWriter, r *http.Request) {
	req := request.New()
	req.ParseHTTPRequest(r)
	microService := strings.Split(req.Path, "/")[1]
	h, ok := urls.MicroServices[microService]
	if !ok {
		log.Printf("Request: %s %s | Status: %d\n", r.Method, r.URL.Path, http.StatusNotFound)
		http.Error(w, "ServiceNotFound", http.StatusNotFound)
		return
	}
	res := response.New()
	handler := handlers.HandlerFunc(h)
	for _, middleware := range settings.CommonMiddlewares {
		handler = middleware(handler)
	}
	handler.ServeHTTP(req, res)

	if res.RedirectURL != "" {
		http.Redirect(w, r, res.RedirectURL, res.StatusCode)
	} else {
		w.Write(res.Body)
	}
}

func main() {
	http.HandleFunc("/", getResponse)

	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
