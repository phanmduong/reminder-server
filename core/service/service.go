package service

import (
	"reminder/core/database"
	"sync"
)

type Service struct {
	DB *database.DatabaseFacade
}

var instance *Service
var once sync.Once

func NewService() *Service {
	return &Service{
		DB:      database.NewDatabase(),
	}
}
func GetService() *Service {
	once.Do(func() {
		instance = NewService()
	})
	return instance
}