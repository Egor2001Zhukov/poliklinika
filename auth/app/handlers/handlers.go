package handlers

import "net/http"

func HelloHandler(http.ResponseWriter, *http.Request) interface{} {
	return "hello"
}
