package gqltype

import (
	"github.com/graphql-go/graphql"
	"reminder/model"
)

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return p.Source.(model.User).ID, nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"fb_id": &graphql.Field{
				Type: graphql.String,
			},
			"google_id": &graphql.Field{
				Type: graphql.String,
			},
			"avatar_url": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
