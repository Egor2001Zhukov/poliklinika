package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"site/app/models"

	"github.com/julienschmidt/httprouter"
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
	findAppointment, err := models.FindAppointments(context.Background(), r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	appointmentJSON, err := json.Marshal(findAppointment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(appointmentJSON)
}

func (h *appointmentHandler) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	findAppointment, err := models.FindAppointment(ctx, objectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	appointmentJSON, err := json.Marshal(findAppointment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(appointmentJSON)
}

func (h *appointmentHandler) Post(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var appointment bson.M
	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	insertAppointment, err := models.InsertAppointment(ctx, &appointment)
	if err != nil {
		return
	}
	appointmentJSON, err := json.Marshal(insertAppointment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(appointmentJSON)
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
