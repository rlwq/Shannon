package app

import (
	"Shannon/bot"
	"Shannon/repository"
	"Shannon/service"
)

type App struct {
	Bot        *bot.Bot
	Service    *service.Service
	Repository *repository.Repository
}

func NewApp(token string, db_path string) *App {
	bot, _ := bot.NewBot(token)
	service := service.NewService()
	repository := repository.NewRepository(db_path)

	return &App{bot, service, repository}
}

func (app *App) Run() {
	app.Bot.Run()
}
