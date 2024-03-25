package db

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

// Database структура для работы с базой данных MongoDB
type Database struct {
	client *mongo.Client
}

// NewDatabase создает новый экземпляр базы данных
func NewDatabase() (*Database, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MAIN_MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MAIN_MONGODB_URI' environment variable. " +
			"See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	// Установите параметры подключения
	clientOptions := options.Client().ApplyURI(uri)

	// Подключение к MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Проверка подключения
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	// Возвращаем экземпляр базы данных
	return &Database{
		client: client,
	}, nil
}

// Disconnect отключает клиент от базы данных
func (db *Database) Disconnect() error {
	return db.client.Disconnect(context.Background())
}

// Collection возвращает коллекцию по ее имени
func (db *Database) Collection(name string) *mongo.Collection {
	return db.client.Database("main").Collection(name)
}
