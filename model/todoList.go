package model

import (
	"reminder/core/model"
	"time"
)

// Table name default is the pluralized version of struct
type TodoList struct {
	model.Model
	Name     string     `json:"name"`
	Note     string     `json:"note"`
	Status   int        `json:"status"`
	Deadline *time.Time `json:"deadline"`
	Image    *string    `json:"image"`
	Group    Group
	GroupID  *uint
}

func (b TodoList) TableName() string {
	return "to_do_list"
}
