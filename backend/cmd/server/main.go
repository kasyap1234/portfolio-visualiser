package main

import (
	"github.com/kasyap1234/portfolio-visualiser/backend/app"
	"github.com/kasyap1234/portfolio-visualiser/backend/configs"
)

func main() {
	config := configs.New()
	app := app.NewApp(config)
	app.Run()

}
