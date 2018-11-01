package gqltype

import (
	"github.com/graphql-go/graphql"
	"reminder/model"
)

var GroupType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Group",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return p.Source.(model.Group).ID, nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
