package schema

import (
	"github.com/graphql-go/graphql"
	"reminder/graphql/query"
)

var DefaultSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:query.RootQuery,
	},
)