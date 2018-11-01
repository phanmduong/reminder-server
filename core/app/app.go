package app

import (
	"github.com/gin-gonic/gin"
	"reminder/core"
	"reminder/core/service"
	"reminder/registry"
	"reminder/router"
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
	server.Static("/othello", "./public/othello")
	server.Static("/static", "./public/graphqlui/static")
	server.Static("/static1", "./public/othello/static1")
}

func (app *App) Init() {
	app.context.RegistryManager.RegisterControllerRegistry(registry.GetControllerRegistry())
	router.RegisterGraphQLRouter(app.context)
	server := core.GetContext().Server

	setupGraphQLUI(server)

}

func (app *App) Run() {
	server := app.context.Server
	server.Run(":9080")
}
