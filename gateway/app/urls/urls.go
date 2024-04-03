package urls

import (
	"auth/app/handlers"
	"net/http"
)

var Urls = map[string]func(http.ResponseWriter, *http.Request){
	"/hello": handlers.HelloHandler,
}
