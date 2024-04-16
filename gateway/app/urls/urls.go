package urls

import (
	"common_go/web/request"
	"common_go/web/response"
	"poliklinika_gateway/app/handlers"
)

var MicroServices = map[string]func(req *request.Request, res *response.Response){
	"server": handlers.ServerHandler,
	"auth":   handlers.AuthHandler,
}
