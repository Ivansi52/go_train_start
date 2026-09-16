package services

import (
	"errors"
	"fmt"
	"go_train/models"
)

type PlayerService struct {
	players []models.Player
}

func (s *PlayerService) AddPlayer(player models.Player) {
	s.players = append(s.players, player)
}

func (s *PlayerService) PrintPlayers() {
	for _, player := range s.players {
		fmt.Println(
			player.Nickname,
			player.Level,
			player.Gold,
			player.Online,
		)
	}
}

func (s *PlayerService) FindPlayer(nickname string) (models.Player, error) {
	for _, player := range s.players {
		if player.Nickname == nickname {
			return player, nil
		}
	}

	return models.Player{}, errors.New("player with this nickname does not exist")
}

func (s *PlayerService) GetFirstPlayer() (models.Player, error) {
	if len(s.players) == 0 {
		return models.Player{}, errors.New("not found")
	}

	return s.players[0], nil
}
