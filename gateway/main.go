package main

import (
	"common_go/web/request"
	"fmt"
	"net/http"
	"poliklinika_gateway/app/urls"
	"strings"
)

func getResponse(w http.ResponseWriter, r *http.Request) {
	req := request.New()
	req.Body = r.Body
	req.Body = r.
	req.Headers = r.
	microService := strings.Split(r.URL.Path, "/")[1]
	f, ok := urls.MicroServices[microService]
	if !ok {
		http.Error(w, "ServiceNotFound", http.StatusNotFound)
		return
	}
	f(w, r)
}

func main() {
	http.HandleFunc("/", getResponse)

	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
