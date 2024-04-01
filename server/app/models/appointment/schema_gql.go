package appointment

import (
	"common_go/dbs/mongodb"
	"context"
	"fmt"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"sync"
)

var appointmentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Appointment",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var appointmentMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createAppointment": &graphql.Field{
			Type: appointmentType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				defer func(client *mongo.Client, ctx context.Context) {
					err := client.Disconnect(ctx)
					if err != nil {
						log.Fatal(err)
					}
				}(mongodb.Client(), context.Background())
				name := params.Args["name"].(string)
				description := params.Args["description"].(string)
				appointment := Appointment{
					Name:        name,
					Description: description,
				}
				err := appointment.Save(context.Background())
				if err != nil {
					return nil, err
				}
				return appointment, nil

			},
		},
	},
})

var appointmentQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"appointment": &graphql.Field{
			Type: appointmentType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				id, ok := p.Args["id"].(string)
				if !ok {
					return nil, fmt.Errorf("ID пользователя не указан")
				}
				appointment := &Appointment{}
				err := appointment.FindByID(context.Background(), id)
				if err != nil {
					return nil, err
				}
				return appointment, nil
			},
		},
	},
})
var (
	Schema graphql.Schema
	once   sync.Once
)

func GetAppointmentSchema() (Schema graphql.Schema, err error) {
	once.Do(func() {
		Schema, err = graphql.NewSchema(graphql.SchemaConfig{
			Query:    appointmentQuery,
			Mutation: appointmentMutation,
		})
	})
	return
}
