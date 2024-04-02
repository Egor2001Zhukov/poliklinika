package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ObjectID = graphql.NewScalar(
	graphql.ScalarConfig{
		Name:        "ObjectID",
		Description: "MongoDB ObjectID",
		Serialize: func(value interface{}) interface{} {
			// Сериализуем значение в строку
			return value.(primitive.ObjectID).Hex()
		},
		ParseValue: func(value interface{}) interface{} {
			// Парсим значение из строки
			objectID, err := primitive.ObjectIDFromHex(value.(string))
			if err != nil {
				return err
			}
			return objectID
		},
		ParseLiteral: func(valueAST ast.Value) interface{} {
			if strValue, ok := valueAST.(*ast.StringValue); ok {
				// Если литерал - строка, преобразуем её в ObjectID
				objectID, err := primitive.ObjectIDFromHex(strValue.Value)
				if err != nil {
					return err
				}
				return objectID
			}
			return nil
		},
	},
)
