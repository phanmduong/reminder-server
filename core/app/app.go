package app

import (
	"reminder/core"
	"reminder/core/service"
	"reminder/registry"
	"reminder/router"
	"github.com/gin-gonic/gin"
)

type App struct {
	service *service.Service
	context *core.Context
}

func NewApp() *App {
	app := &App{
		context: core.GetContext(),
		service: service.NewService(),
	}
	return app
}

func setupGraphQLUI(server *gin.Engine) {
	server.Static("/graphqlui", "./public/graphqlui")
	server.Static("/static", "./public/graphqlui/static")
}

func (app *App) Init() {
	app.context.RegistryManager.RegisterControllerRegistry(registry.GetControllerRegistry())
	router.RegisterGraphQLRouter(app.context)
	server := core.GetContext().Server


	setupGraphQLUI(server)

}

func (app *App) Run() {
	server := app.context.Server
	server.Run(":8080")
}
