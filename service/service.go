package service

import (
	"Shannon/bot"
	"Shannon/repo"
)

type Service struct {
	bot  *bot.Bot
	repo *repo.Repo
}

func NewService() *Service {
	return &Service{}
}

func (service *Service) LinkBotRepo(bot *bot.Bot, repo *repo.Repo) {
	service.bot = bot
	service.repo = repo
}
