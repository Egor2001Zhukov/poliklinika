package urls

import (
	"auth/app/handlers"
	"net/http"
)

var urls = map[string]func(http.ResponseWriter, *http.Request) interface{}{
	"/hello": handlers.HelloHandler,
}
