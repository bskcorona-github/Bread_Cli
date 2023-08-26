package main

import (
	"log"

	"github.com/bskcorona-github/Bread_Cli/internal/api"
)

func main() {
	app := api.NewApp()
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
