package util

import (
	"github.com/graphql-go/graphql"
	"log"
)

func ExecuteQuery(query string, variable map[string]interface{}, schema graphql.Schema, ) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  query,
		VariableValues: variable,
	})

	if len(result.Errors) > 0 {
		log.Printf("errors: %v", result.Errors)
	}
	return result
}
