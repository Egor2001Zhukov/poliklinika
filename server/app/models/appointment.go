package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/url"
	db "site/databases/mongodb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var AppointmentCollection = "Appointment"

type Appointment struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
}

func InsertAppointment(ctx context.Context, a *bson.M) (*bson.M, error) {
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
	(*a)["_id"] = insertedID
	return a, nil
}

func FindAppointment(ctx context.Context, id primitive.ObjectID) (*bson.M, error) {

	database, err := db.NewDatabase()
	if err != nil {
		return nil, err
	}

	filter := primitive.M{"_id": id}
	var appointment bson.M
	err = database.Collection(AppointmentCollection).FindOne(ctx, filter).Decode(&appointment)
	if err != nil {
		return nil, err
	}

	return &appointment, nil
}

func FindAppointments(ctx context.Context, queryParams url.Values) (*[]bson.M, error) {
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

func UpdateAppointment(ctx context.Context, id primitive.ObjectID, a *bson.M) (*bson.M, error) {
	database, err := db.NewDatabase()
	if err != nil {
		return nil, err
	}
	filter := primitive.M{"_id": id}
	var updateAppointments bson.M
	err = database.Collection(AppointmentCollection).FindOneAndUpdate(ctx, filter, a).Decode(&updateAppointments)
	if err != nil {
		return nil, err
	}
	return &updateAppointments, nil
}
