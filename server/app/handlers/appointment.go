package handlers

import (
	"context"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"server/app/models/appointment"
)

type appointmentHandler struct {
}

func NewAppointmentHandler() Handler {
	return &appointmentHandler{}
}

func (h *appointmentHandler) Register(router *httprouter.Router) {
	router.POST("/graphql", h.Post)
}

func (h *appointmentHandler) Post(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()
	// Создаем новый контекст для выполнения запроса GraphQL
	ctx := context.Background()

	// Распарсим тело запроса
	var requestBody map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	schema, err := appointment.GetAppointmentSchema()
	if err != nil {
		http.Error(w, "Invalid schema", http.StatusInternalServerError)
		return
	}
	// Выполним запрос GraphQL
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: requestBody["query"].(string),
		Context:       ctx,
	})
	// Преобразуем результат выполнения запроса в JSON и отправим клиенту
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		return
	}
}
