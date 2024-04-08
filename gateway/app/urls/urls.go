package urls

import (
	"net/http"
	"poliklinika_gateway/app/handlers"
)

var MicroServices = map[string]func(w http.ResponseWriter, r *http.Request){
	"gateway": handlers.HelloHandler,
	"auth":    handlers.HelloHandler,
}
