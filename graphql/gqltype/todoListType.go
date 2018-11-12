package gqltype

import (
	"github.com/graphql-go/graphql"
	"reminder/model"
	"time"
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
			"image": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: graphql.Int,
			},
			"deadline": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					deadline := p.Source.(model.TodoList).Deadline
					if deadline == nil {
						return nil, nil
					}
					return deadline.Format(time.RFC3339), nil
				},
			},
		},
	},
)
