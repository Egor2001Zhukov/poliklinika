package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Document interface {
	GetCollection() *mongo.Collection
	Save(context.Context) error
}

//func FindDocument(ctx context.Context, document Document, id primitive.ObjectID) (*bson.M, error) {
//	filter := primitive.M{"_id": id}
//	var doc bson.M
//	err := document.GetCollection().FindOne(ctx, filter).Decode(&doc)
//	if err != nil {
//		return nil, err
//	}
//	return &doc, nil
//}

//func FindDocuments(ctx context.Context, document Document, query interface{}) (*[]bson.M, error) {
//	cursor, err := document.GetCollection().Find(ctx, filter, options.Find().SetProjection(projection))
//	if err != nil {
//		return nil, err
//	}
//	var docs []bson.M
//	if err = cursor.All(context.TODO(), &docs); err != nil {
//		return nil, err
//	}
//
//	return &docs, nil
//}
//
//func UpdateDocument(ctx context.Context, id primitive.ObjectID, a *bson.M) (*bson.M, error) {
//	database, err := db.Client()
//	if err != nil {
//		return nil, err
//	}
//	filter := primitive.M{"_id": id}
//	var updateAppointments bson.M
//	err = database.Collection(AppointmentCollection).FindOneAndUpdate(ctx, filter, a).Decode(&updateAppointments)
//	if err != nil {
//		return nil, err
//	}
//	return &updateAppointments, nil
//}
