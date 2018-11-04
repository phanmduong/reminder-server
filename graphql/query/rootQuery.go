package query

import (
	"github.com/graphql-go/graphql"
	"reminder/graphql/field"
)

var RootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user":   field.FieldUser,
			"groups": field.FieldGroups,
			"todoLists": field.FieldTodoLists,
		},
	},
)
