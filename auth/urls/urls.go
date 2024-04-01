package urls

import "net/http"
import "auth/handlers"

var urls = map[string]func(http.ResponseWriter, *http.Request) interface{}{
	"/hello": handlers.HelloHandler,
}
