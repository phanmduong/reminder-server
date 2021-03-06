package field

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"reminder/authorize"
	"reminder/core/service"
	"reminder/graphql/gqltype"
	"reminder/model"
	"time"
)

//var FieldTodoList = &graphql.Field{
//	Type:        gqltype.TodoListType,
//	Description: "Get todo list by id",
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

var FieldTodoLists = &graphql.Field{
	Type:        graphql.NewList(gqltype.TodoListType),
	Description: "Get all todo list",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		db := service.GetService().DB.DB
		token, ok := p.Args["token"].(string)
		groupID, ok1 := p.Args["group_id"].(int)
		fmt.Println(groupID)
		var todoLists []model.TodoList
		if ok {
			isValid, _ := authorize.CheckAuthorization(token)
			if isValid && ok1 {
				if groupID == -1 {
					db.Debug().Where("DATE(deadline) = DATE(NOW())").Order("created_at asc").Find(&todoLists)
					return todoLists, nil
				}

				db.Where("group_id = ?", groupID).Order("created_at asc").Find(&todoLists)
				return todoLists, nil
			} else {
				return nil, nil
			}
		}

		return nil, nil
	},
}

var MutationTodoList = &graphql.Field{
	Type:        gqltype.TodoListType,
	Description: "create todo list",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"note": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"deadline": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"image": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"group_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		db := service.GetService().DB.DB
		token, ok := p.Args["token"].(string)
		name, ok1 := p.Args["name"].(string)
		note, ok1 := p.Args["note"].(string)
		image, _ := p.Args["image"].(string)
		deadline, okDeadline := p.Args["deadline"].(string)
		id, ok2 := p.Args["id"].(int)
		group_id, _ := p.Args["group_id"].(int)
		groupID := uint(group_id)
		if ok && ok1 {
			isValid, _ := authorize.CheckAuthorization(token)
			if isValid {
				var todoList = model.TodoList{Name: name, Note: note, GroupID: &groupID, Image: &image}
				if (okDeadline) {
					timeDeadline, _ := time.Parse(time.RFC3339, deadline)
					todoList.Deadline = &timeDeadline
				}
				if ok2 {
					todoList.ID = uint(id)
					db.Debug().Save(&todoList)
				} else {
					db.Debug().Create(&todoList)
				}

				return todoList, nil

			} else {
				return nil, nil
			}
		}

		return nil, nil
	},
}

var MutationChangeStatusTodoList = &graphql.Field{
	Type:        gqltype.TodoListType,
	Description: "create group",
	Args: graphql.FieldConfigArgument{
		"token": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"status": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		db := service.GetService().DB.DB
		token, ok := p.Args["token"].(string)
		id, ok1 := p.Args["id"].(int)
		status, ok2 := p.Args["status"].(int)
		if ok && ok1 {
			isValid, _ := authorize.CheckAuthorization(token)
			if isValid {
				var todoList = model.TodoList{}
				db.First(&todoList, id)
				if ok2 {
					todoList.Status = status
				}
				db.Save(&todoList)

				return todoList, nil

			} else {
				return nil, nil
			}
		}

		return nil, nil
	},
}

var MutationDeleteTodoList = &graphql.Field{
	Type:        gqltype.TodoListType,
	Description: "delete todo list",
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
			isValid, _ := authorize.CheckAuthorization(token)
			if isValid {
				db.Debug().Where("id = ?", id).Delete(&model.TodoList{})
				return nil, nil
			} else {
				return nil, nil
			}
		}

		return nil, nil
	},
}
