package service

import (
	"Shannon/repo"
	"Shannon/shannon"
	"math/rand/v2"
)

type Service struct {
	repo *repo.Repo
}

func NewService(repo *repo.Repo) *Service {
	return &Service{repo}
}

func (service *Service) CreateProfile(profile shannon.Profile) {
	service.repo.CreateProfile(profile)
}

func (service *Service) NextProfileFor(user int64) shannon.Profile {
	profiles := service.repo.GetProfiles()
	randomProfile := profiles[rand.IntN(len(profiles))]
	return randomProfile
}

func (service *Service) DoesProfileExist(user int64) bool {
	return service.repo.DoesProfileExist(user)
}
