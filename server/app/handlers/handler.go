package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Handler interface {
	Register(router *httprouter.Router)
	GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Get(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Post(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Put(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
