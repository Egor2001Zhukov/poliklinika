package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	mongodb "server/databases/mongodb"
)

type Appointment struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
}

func (a *Appointment) GetCollection() *mongo.Collection {
	collection := mongodb.MainDBGetCollection("appointments")
	return collection
}
