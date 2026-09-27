package services

import (
	"go_train/models"
	"go_train/repositories"
)

type PlayerService struct {
	repository *repositories.PlayerRepository
}

func (s *PlayerService) Add(player models.Player) models.Player {
	return s.repository.Add(player)
}

func (s *PlayerService) GetPlayers() []models.Player {
	return s.repository.GetAll()
}

func (s *PlayerService) FindPlayer(id string) (models.Player, error) {
	return s.repository.GetByID(id)
}

func (s *PlayerService) DeletePlayer(id string) error {
	return s.repository.Delete(id)
}

func NewPlayerService(repository *repositories.PlayerRepository) *PlayerService {
	return &PlayerService{
		repository: repository,
	}
}
