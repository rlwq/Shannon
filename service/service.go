package service

import (
	"Shannon/repo"
	"Shannon/shannon"
)

type Service struct {
	repo *repo.Repo
}

func NewService(repo *repo.Repo) *Service {
	return &Service{repo}
}

func (service *Service) CreateProfile(profile shannon.Profile) {
	service.CreateProfile(profile)
}
