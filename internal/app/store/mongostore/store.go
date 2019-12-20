package mongostore

import (
	"github.com/go-redis/redis/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"test/internal/app/store"
)

type Store struct {
	db             *mongo.Client
	redis          *redis.Client
	userRepository *UserRepository
	gameRepository *GameRepository
}

func New(db *mongo.Client, redis *redis.Client) *Store {
	return &Store{
		db:    db,
		redis: redis,
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
