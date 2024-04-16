package handlers

import (
	"common_go/web/request"
	"common_go/web/response"
)

func HelloHandler(req *request.Request, res *response.Response) {
	res.SetHeader("Content-Type", []string{"text/plain; charset=utf-16"})
	res.WriteBody([]byte("hello"))
	res.SetStatusCode(201)
}
