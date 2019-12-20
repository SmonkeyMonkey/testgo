package store

import "test/internal/app/models"

type UserRepository interface {
	Create(user *models.User) error
	GetAll(page int) []models.User
}
type GameRepository interface {
	Create(game *models.Game, userId string) error
	GetAll(page int) []models.Game
	GetTopUsers(page int) []models.User
}
