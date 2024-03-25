package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type appointmentHandler struct {
}

func NewAppointmentHandler() Handler {
	return &appointmentHandler{}
}

func (h *appointmentHandler) Register(router *httprouter.Router) {
	router.GET("/appointments", h.GetList)
	router.GET("/appointments/:id", h.Get)
	router.POST("/appointments", h.Post)
	router.PUT("/appointments/:id", h.Put)
	router.DELETE("/appointments/:id", h.Delete)
}

func (h *appointmentHandler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_, err := w.Write([]byte(fmt.Sprint("GetList Appointment")))
	if err != nil {
		panic(err)
	}
}

func (h *appointmentHandler) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	_, err := w.Write([]byte(fmt.Sprintf("Get %s Appointment", id)))
	if err != nil {
		panic(err)
	}
}

func (h *appointmentHandler) Post(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_, err := w.Write([]byte(fmt.Sprint("Post Appointment")))
	if err != nil {
		panic(err)
	}
}

func (h *appointmentHandler) Put(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	_, err := w.Write([]byte(fmt.Sprintf("Put %s Appointment", id)))
	if err != nil {
		panic(err)
	}
}

func (h *appointmentHandler) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	_, err := w.Write([]byte(fmt.Sprintf("Delete %s Appointment", id)))
	if err != nil {
		panic(err)
	}
}
