package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type userHandler struct {
}

func NewUserHandler() Handler {
	return &userHandler{}
}
func (h *userHandler) Register(router *httprouter.Router) {
	router.GET("/users", h.GetList)
	router.GET("/users/:id", h.Get)
	router.POST("/users", h.Post)
	router.PUT("/users/:id", h.Put)
	router.DELETE("/users/:id", h.Delete)
}

func (h *userHandler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_, err := w.Write([]byte(fmt.Sprint("GetList User")))
	if err != nil {
		panic(err)
	}
}

func (h *userHandler) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	_, err := w.Write([]byte(fmt.Sprintf("Get %s User", id)))
	if err != nil {
		panic(err)
	}
}

func (h *userHandler) Post(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_, err := w.Write([]byte(fmt.Sprint("Post User")))
	if err != nil {
		panic(err)
	}
}

func (h *userHandler) Put(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	_, err := w.Write([]byte(fmt.Sprintf("Put %s User", id)))
	if err != nil {
		panic(err)
	}
}

func (h *userHandler) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	_, err := w.Write([]byte(fmt.Sprintf("Delete %s User", id)))
	if err != nil {
		panic(err)
	}
}
