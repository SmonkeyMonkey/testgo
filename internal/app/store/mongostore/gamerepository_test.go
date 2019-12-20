package mongostore_test

import (
	"github.com/stretchr/testify/assert"
	"test/internal/app/models"
	"test/internal/app/store/mongostore"
	"testing"
)

func TestGameRepository_Create(t *testing.T) {
	db := mongostore.TestDB(t)
	red := mongostore.TestRedis(t)
	s := mongostore.New(db, red)
	g := models.TestGame(t)
	assert.NoError(t, s.Game().Create(g, "123"))
}
func TestGameRepository_GetAll(t *testing.T) {
	db := mongostore.TestDB(t)
	red := mongostore.TestRedis(t)
	s := mongostore.New(db, red)
	assert.NotEmpty(t, s.Game().GetAll(1))
	assert.Len(t, s.Game().GetAll(1), 20)
}
func TestGameRepository_GetGames(t *testing.T) {
	db := mongostore.TestDB(t)
	red := mongostore.TestRedis(t)
	s := mongostore.New(db, red)
	assert.NotEmpty(t, s.Game().GetGames())
}
func TestGameCount(t *testing.T) {
	db := mongostore.TestDB(t)
	red := mongostore.TestRedis(t)
	s := mongostore.New(db, red)
	s.Game().GetTopUsers(1)
}
