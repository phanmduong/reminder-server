package gqltype

import (
	"github.com/graphql-go/graphql"
	"reminder/model"
)

var TodoListType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "TodoList",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return p.Source.(model.TodoList).ID, nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"note": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: graphql.Int,
			},
			"deadline": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)
