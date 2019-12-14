package mongostore

import (
	"go.mongodb.org/mongo-driver/mongo"
	"test/internal/app/store"
)

type Store struct {
	db             *mongo.Client
	userRepository *UserRepository
	gameRepository *GameRepository
}

func New(db *mongo.Client) *Store {
	return &Store{
		db: db,
	}
}
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}
func (s *Store) Game() store.GameRepository {
	if s.gameRepository != nil {
		return s.gameRepository
	}
	s.gameRepository = &GameRepository{
		store: s,
	}
	return s.gameRepository
}
