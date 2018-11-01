package model

import "reminder/core/model"

// Table name default is the pluralized version of struct
type Group struct {
	model.Model
	Name   string `json:"name"`
	User   User
	UserID uint
}

func (b Group) TableName() string {
	return "groups"
}
