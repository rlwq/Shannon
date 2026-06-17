package app

import (
	"Shannon/bot"
	"Shannon/repo"
	"Shannon/service"
)

type App struct {
	Bot        *bot.Bot
	Service    *service.Service
	Repository *repo.Repo
}

func NewApp(token string, db_path string) *App {
	repo := repo.NewRepository(db_path)
	service := service.NewService(repo)
	bot, _ := bot.NewBot(token, service)

	return &App{bot, service, repo}
}

func (app *App) Run() {
	app.Bot.Run()
}
