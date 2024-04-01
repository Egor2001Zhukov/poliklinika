package appointment

import (
	"common_go/dbs/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Appointment struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

func (a *Appointment) GetCollection() *mongo.Collection {
	collection := mongodb.MainDBGetCollection("Appointment")
	return collection
}

func (a *Appointment) Save(ctx context.Context) error {
	documentBSON, err := bson.Marshal(a)
	if err != nil {
		return err
	}
	opts := options.FindOneAndReplace().SetUpsert(true)
	filter := bson.D{{"_id", a.ID}}
	err = a.GetCollection().FindOneAndReplace(ctx, filter, documentBSON, opts).Decode(a)
	if err != nil {
		return err
	}
	return nil
}

func (a *Appointment) FindByID(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := primitive.M{"_id": objectID}
	err = a.GetCollection().FindOne(ctx, filter).Decode(a)
	return err
}
