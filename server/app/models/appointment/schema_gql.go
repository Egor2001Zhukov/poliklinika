package appointment

import (
	"common_go/gql"
	"context"
	"fmt"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
)

var appointmentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "appointment",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: gql.ObjectID,
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
		"updateAppointment": &graphql.Field{
			Type: appointmentType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(gql.ObjectID),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id := params.Args["id"].(primitive.ObjectID)
				name := params.Args["name"].(string)
				description := params.Args["description"].(string)
				appointment := Appointment{
					ID:          id,
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
	schema graphql.Schema
	once   sync.Once

	schemaErr error
)

func GetAppointmentSchema() (*graphql.Schema, error) {
	once.Do(func() {
		schema, schemaErr = graphql.NewSchema(graphql.SchemaConfig{
			Query:    appointmentQuery,
			Mutation: appointmentMutation,
		})
	})
	if schemaErr != nil {
		return nil, schemaErr
	}
	return &schema, nil
}
