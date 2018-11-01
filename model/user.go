package model

import "reminder/core/model"

// Table name default is the pluralized version of struct
type User struct {
	model.Model
	Name      string `json:"name"`
	Email     string `json:"email"`
	FbID      string `json:"fb_id"`
	GoogleID  string `json:"google_id"`
	AvatarURL string `json:"avatar_url"`
}

func (b User) TableName() string {
	return "users"
}
