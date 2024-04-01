package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Handler interface {
	Register(router *httprouter.Router)
	Post(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
