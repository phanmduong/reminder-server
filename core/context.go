package core

import (
	"sync"
	"github.com/gin-gonic/gin"
	"reminder/core/registry"
	"github.com/gin-contrib/cors"
	"time"
)

type Context struct {
	Server          *gin.Engine
	RegistryManager *registry.RegistryManager
}

var instance *Context
var once sync.Once

func GetContext() *Context {
	once.Do(func() {
		instance = NewContext()
	})
	return instance
}

func NewContext() *Context {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		//AllowOrigins:     []string{""},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	context := &Context{
		Server:          server,
		RegistryManager: registry.NewRegistryManager(),
	}
	return context
}
