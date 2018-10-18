package mutation

import (
	"github.com/graphql-go/graphql"
	"reminder/graphql/field"
)

var RootMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"login": field.Login,
		},
	},
)
