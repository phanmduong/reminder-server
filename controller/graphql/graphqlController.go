package graphql

import (
	"github.com/gin-gonic/gin"
	"reminder/graphql/util"
	"reminder/graphql/schema"
	"encoding/json"
	)

type graphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func GraphQL(c *gin.Context) {
	rawData, _ := c.GetRawData()
	data := string(rawData)
	var request graphQLRequest
	json.Unmarshal([]byte(data), &request)

	result := util.ExecuteQuery(request.Query, request.Variables, schema.DefaultSchema)
	c.JSON(200, result)
}
