package field

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"reminder/authorize"
	"reminder/core/service"
	"reminder/graphql/gqltype"
	"reminder/model"
)

//var FieldGroup = &graphql.Field{
//	Type:        gqltype.GroupType,
//	Description: "Get group by id",
//	Args: graphql.FieldConfigArgument{
//		"token": &graphql.ArgumentConfig{
//			Type: graphql.String,
//		},
//	},
//	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
//
//		token, ok := p.Args["token"].(string)
//		fmt.Println(token)
//		if ok {
//			isValid, _ := authorize.CheckAuthorization(token)
//			if isValid {
//				return nil, nil
//			}
//		}
//
//		return nil, nil
//	},
//}

var FieldGroups = &graphql.Field{
	Type:        graphql.NewList(gqltype.GroupType),
	Description: "Get group by id",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		db := service.GetService().DB.DB
		token, ok := p.Args["token"].(string)
		fmt.Println(token)
		var groups []model.Group
		if ok {
			isValid, user := authorize.CheckAuthorization(token)
			if isValid {
				db.Where("user_id = ?", user.ID).Order("created_at asc").Find(&groups)
				return groups, nil
			} else {
				return nil, nil;
			}
		}

		return nil, nil
	},
}

var MutationGroup = &graphql.Field{
	Type:        gqltype.GroupType,
	Description: "create group",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		db := service.GetService().DB.DB
		token, ok := p.Args["token"].(string)
		id, ok2 := p.Args["id"].(int)
		name, ok1 := p.Args["name"].(string)

		if ok && ok1 {
			isValid, user := authorize.CheckAuthorization(token)
			if isValid {
				var group = model.Group{Name: name, UserID: user.ID}
				if ok2 {
					group.ID = uint(id)
					db.Save(&group)
				} else {
					db.Create(&group)
				}
				return group, nil

			} else {
				return nil, nil
			}
		}

		return nil, nil
	},
}

var MutationDeleteGroup = &graphql.Field{
	Type:        gqltype.GroupType,
	Description: "delete group",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		db := service.GetService().DB.DB
		token, ok := p.Args["token"].(string)
		id, ok1 := p.Args["id"].(int)
		if ok && ok1 {
			isValid, user := authorize.CheckAuthorization(token)
			if isValid {
				db.Debug().Where("id = ? AND user_id = ?", id, user.ID).Delete(&model.Group{})
				return nil, nil
			} else {
				return nil, nil
			}
		}

		return nil, nil
	},
}
