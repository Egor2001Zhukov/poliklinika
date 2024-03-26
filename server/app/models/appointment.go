package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	db "site/databases/mongodb"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func FindAppointment(ctx context.Context, filter primitive.M) (*bson.M, error) {
	database, err := db.NewDatabase()
	if err != nil {
		return nil, err
	}

	var appointment bson.M
	err = database.Collection(AppointmentCollection).FindOne(ctx, filter).Decode(&appointment)
	if err != nil {
		return nil, err
	}

	return &appointment, nil
}

func FindAppointments(ctx context.Context, filter interface{}) (*[]bson.M, error) {
	database, err := db.NewDatabase()
	if err != nil {
		return nil, err
	}
	if filter == nil {
		filter = primitive.M{}
	}
	cursor, err := database.Collection(AppointmentCollection).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var appointments []bson.M
	if err = cursor.All(context.TODO(), &appointments); err != nil {
		log.Fatal(err)
	}

	return &appointments, nil
}
