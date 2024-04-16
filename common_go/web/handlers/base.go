package handlers

import (
	"common_go/web/request"
	"common_go/web/response"
)

type Handler interface {
	ServeHTTP(*request.Request, *response.Response)
}
type HandlerWrapper struct {
	f func(*request.Request, *response.Response)
}

func (handler *HandlerWrapper) ServeHTTP(req *request.Request, res *response.Response) {
	handler.f(req, res)
}

func HandlerFunc(f func(*request.Request, *response.Response)) Handler {
	return &HandlerWrapper{f}
}
