package model

import (
	"time"
	)

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time

	//ignore field in gorm
	Args map[string]interface{} `gorm:"-"` // Args	in graphql
}
