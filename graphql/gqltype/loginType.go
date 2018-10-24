package gqltype

import (
	"github.com/graphql-go/graphql"
)

var LoginType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"token": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return p.Source.(string), nil
				},
			},
		},
	},
)
