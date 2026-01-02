package app

import (
	"github.com/kasyap1234/portfolio-visualiser/backend/configs"
)

type App struct {
	config *configs.Config
	
}

func NewApp(config *configs.Config) *App {
	return &App{
		config: config,
	}
}

func (a *App) Run() {
  
}
