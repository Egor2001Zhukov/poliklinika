package handlers

import (
	"context"
	"encoding/json"
	"fmt"
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
	queryParams := r.URL.Query()
	filter := primitive.M{}
	for key, values := range queryParams {
		// Проверяем, есть ли у параметра несколько значений
		if len(values) == 1 {
			// Если значение параметра одно, добавляем его в фильтр
			filter[key] = values[0]
		} else {
			// Если у параметра несколько значений, добавляем их как массив
			filter[key] = values
		}
	}
	fmt.Println(filter)

	ctx := context.Background()
	findAppointment, err := models.FindAppointments(ctx, filter)
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
	filter := primitive.M{"_id": objectID}
	ctx := context.Background()
	findAppointment, err := models.FindAppointment(ctx, filter)
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
	var appointment models.Appointment

	err := json.NewDecoder(r.Body).Decode(&appointment)
	fmt.Println(&appointment)
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
