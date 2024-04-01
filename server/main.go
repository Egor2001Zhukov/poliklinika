package main

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"server/app/handlers"
	mongodb "server/databases/mongodb"
)

func main() {
	// create router
	router := httprouter.New()

	// register handlers
	handlers.NewUserHandler().Register(router)
	handlers.NewAppointmentHandler().Register(router)

	// start server
	fmt.Println("Сервер запущен на 127.0.0.1:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	} // Стартуем веб-сервер на порту 8080

	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(mongodb.Client(), context.Background())
}
