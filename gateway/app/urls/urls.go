package urls

import (
	"auth/app/handlers"
	"net/http"
)

var MicroServices = map[string]func(w http.ResponseWriter, r *http.Request){
	"gateway": handlers.HelloHandler,
	"auth":    handlers.HelloHandler,
}
