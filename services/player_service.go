package services

import (
	"errors"
	"fmt"
	"go_train/models"
)

type PlayerService struct {
	players []models.Player
}

func (s *PlayerService) AddPlayer(player models.Player) models.Player {
	s.players = append(s.players, player)
	return player
}

func (s *PlayerService) GetPlayers() []models.Player {
	return s.players
}

func (s *PlayerService) FindPlayer(id string) (models.Player, error) {
	for _, player := range s.players {
		if player.Id == id {
			return player, nil
		}
	}

	return models.Player{}, errors.New("player with this id does not exist")
}

func (s *PlayerService) DeletePlayer(id string) error {
	for i, player := range s.players {
		if player.Id == id {
			s.players = append(s.players[:i], s.players[i+1:]...)
			return nil
		}
	}
	return errors.New("player not found")
}
