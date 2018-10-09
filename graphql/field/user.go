package field

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"reminder/core/service"
	"reminder/graphql/gqltype"
	"reminder/model"
)

var FieldUser = &graphql.Field{
	Type:        gqltype.UserType,
	Description: "Get user by id",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"fb_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"google_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		db := service.GetService().DB.DB

		var user model.User

		id, ok := p.Args["id"].(int)
		if ok {

			db.Debug().Find(&user, id)
			fmt.Println(user.ID)
			return user, nil
		}

		fb_id, ok := p.Args["fb_id"].(int)
		if ok {
			db.Where("fb_id = ?", fb_id).First(&user)
			return user, nil
		}

		google_id, ok := p.Args["google_id"].(int)
		if ok {
			db.Where("google_id = ?", google_id).First(&user)
			return user, nil
		}
		return nil, nil
	},
}

var FieldUsers = &graphql.Field{
	Type:        graphql.NewList(gqltype.UserType),
	Description: "Get user list",
	Args: graphql.FieldConfigArgument{
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		var users []model.User
		db := service.GetService().DB.DB
		db.Order("created_at asc").Find(&users)

		return users, nil
	},
}
