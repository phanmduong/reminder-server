package router

import (
	"reminder/core"
	"reminder/controller/graphql"
)

func RegisterGraphQLRouter(context *core.Context) {
	server := context.Server
	server.POST("graphql", graphql.GraphQL)
	//server.GET("graphql", graphql.GraphQL)
}
