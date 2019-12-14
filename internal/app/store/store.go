package store

type Store interface {
	User() UserRepository
	Game() GameRepository
}
