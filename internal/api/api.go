package api

import (
	"log"

	"github.com/bskcorona-github/Bread_Cli/internal/database"
	"github.com/bskcorona-github/Bread_Cli/pkg/graphql"
)

type App struct {
	DB      *database.DB
	GraphQL *graphql.Server
}

func NewApp() *App {
	db := database.NewDB("user=postgres dbname=postgres sslmode=disable password=tkz2001r")
	return &App{
		DB:      db,
		GraphQL: graphql.NewServer(db),
	}
}

func (app *App) Run() error {
	app.GraphQL.SetupRoutes()
	log.Println("GraphQL サーバーが http://localhost:8080/graphql で稼働しています")
	return app.GraphQL.StartServer()
}
