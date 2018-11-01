package model

import "reminder/core/model"

// Table name default is the pluralized version of struct
type TodoList struct {
	model.Model
	Name    string `json:"name"`
	Note    string `json:"note"`
	Status  string `json:"status"`
	Group   Group
	GroupID int
}

func (b TodoList) TableName() string {
	return "to_do_list"
}
