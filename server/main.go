package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"site/app/handlers"
)

func main() {
	// create router
	router := httprouter.New()

	// register handlers
	handlers.NewUserHandler().Register(router)
	handlers.NewAppointmentHandler().Register(router)

	// start server
	fmt.Println("Сервер запущен на 127.0.0.1:8080")
	http.ListenAndServe(":8080", router) // Стартуем веб-сервер на порту 8080
}
