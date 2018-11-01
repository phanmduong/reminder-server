package field

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"reminder/authorize"
	"reminder/graphql/gqltype"
)

var FieldUser = &graphql.Field{
	Type:        gqltype.UserType,
	Description: "Get user by id",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		token, ok := p.Args["token"].(string)
		fmt.Println(token)
		if ok {
			isValid, user := authorize.CheckAuthorization(token)
			if isValid {
				return user, nil
			}
		}

		return nil, nil
	},
}
