package store

import "test/internal/app/models"

type UserRepository interface {
	Create(user *models.User) error
	GetAll() []models.User
}
type GameRepository interface {
	Create(game *models.Game) error
	GetAll() []models.Game
}
