package handlers

import "net/http"

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-16")
	w.Write([]byte("hello"))
}
