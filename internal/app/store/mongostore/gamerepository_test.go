package mongostore_test

import (
	"github.com/stretchr/testify/assert"
	"test/internal/app/models"
	"test/internal/app/store/mongostore"
	"testing"
)

var (
	t   testing.T
	db  = mongostore.TestDB(&t)
	red = mongostore.TestRedis(&t)
	s   = mongostore.New(db, red)
)

func TestGameRepository_Create(t *testing.T) {
	g := models.TestGame(t)
	assert.NoError(t, s.Game().Create(g, "5defaf210d7a7fb756a4396c"))
}

func TestGameRepository_GetAll(t *testing.T) {
	assert.NotEmpty(t, s.Game().GetAll(1))
	assert.Len(t, s.Game().GetAll(1), 30)
}

func TestGetTopUsers(t *testing.T) {
	result := s.Game().GetTopUsers(1)
	assert.Len(t, result, 3)
}
func TestGetSortedGames(t *testing.T) {
	r := s.Game().GetSortedGames("game_type", 1)
	assert.Len(t, r, 30)
}

func BenchmarkGameRepository_GetAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.Game().GetAll(1)
	}
}

func BenchmarkGameRepository_GetTopUsers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.Game().GetTopUsers(1)
	}
}
func BenchmarkGameRepository_GetSortedGames(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.Game().GetSortedGames("game_type", 1)
	}
}
