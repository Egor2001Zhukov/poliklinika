package mongodb

import (
	"context"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	once   sync.Once
)

func Client() *mongo.Client {
	once.Do(func() {
		uri := os.Getenv("MONGODB_URI")
		if uri == "" {
			log.Fatal("You must set your 'MONGODB_URI' environment variable.")
		}

		// Установите параметры подключения
		clientOptions := options.Client().ApplyURI(uri)

		// Подключение к MongoDB
		var err error
		client, err = mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatalf("Error connecting to MongoDB: %v", err)
		}
		// Проверка подключения
		err = client.Ping(context.Background(), nil)
		if err != nil {
			log.Fatalf("Error pinging MongoDB: %v", err)
		}
	})

	return client
}

func MainDBGetCollection(name string) *mongo.Collection {
	return Client().Database("main").Collection(name)
}
