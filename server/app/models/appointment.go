package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	db "site/databases/mongodb"
)

var AppointmentCollection = "Appointment"

type Appointment struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
}

func InsertAppointment(ctx context.Context, a *Appointment) (*Appointment, error) {
	database, err := db.NewDatabase()
	if err != nil {
		return nil, err
	}

	result, err := database.Collection(AppointmentCollection).InsertOne(ctx, a)
	if err != nil {
		return nil, err
	}
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, err
	}
	a.ID = insertedID
	return a, nil
}

func FindAppointment(ctx context.Context, filter primitive.M) (*Appointment, error) {
	database, err := db.NewDatabase()
	if err != nil {
		return nil, err
	}

	var appointment Appointment
	err = database.Collection(AppointmentCollection).FindOne(ctx, filter).Decode(&appointment)
	if err != nil {
		return nil, err
	}

	return &appointment, nil
}
