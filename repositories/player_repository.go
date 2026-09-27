package repositories

import (
	"errors"
	"go_train/models"
)

type PlayerRepository struct {
	players []models.Player
}

func (r *PlayerRepository) Add(player models.Player) models.Player {
	r.players = append(r.players, player)
	return player
}

func (r *PlayerRepository) GetAll() []models.Player {
	return r.players
}

func (r *PlayerRepository) GetByID(id string) (models.Player, error) {
	for _, player := range r.players {
		if player.Id == id {
			return player, nil
		}
	}

	return models.Player{}, errors.New("There is no player with this id")
}

func (r *PlayerRepository) Delete(id string) error {
	for i, player := range r.players {
		if player.Id == id {
			r.players = append(r.players[:i], r.players[i+1:]...)
			return nil
		}
	}

	return errors.New("Player wasn't found")
}
