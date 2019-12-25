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
	assert.NoError(t, s.Game().Create(g, "5defaf210d7a7fb756a4396c"))
}

func TestGameRepository_GetAll(t *testing.T) {
	db := mongostore.TestDB(t)
	red := mongostore.TestRedis(t)
	s := mongostore.New(db, red)
	assert.NotEmpty(t, s.Game().GetAll(1))
	assert.Len(t, s.Game().GetAll(1), 30)
}

func TestGetTopUsers(t *testing.T) {
	db := mongostore.TestDB(t)
	red := mongostore.TestRedis(t)
	s := mongostore.New(db, red)
	result := s.Game().GetTopUsers(1)
	assert.Len(t, result, 15)
}
func TestGetSortedGames(t *testing.T){
	db := mongostore.TestDB(t)
	red := mongostore.TestRedis(t)
	s := mongostore.New(db, red)
	r := s.Game().GetSortedGames("game_type",1)
	assert.Len(t,r,30)
}