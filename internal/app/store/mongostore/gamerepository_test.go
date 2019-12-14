package mongostore_test

import (
	"github.com/stretchr/testify/assert"
	"test/internal/app/models"
	"test/internal/app/store/mongostore"
	"testing"
)

func TestGameRepository_Create(t *testing.T) {
	db := mongostore.TestDB(t)
	s := mongostore.New(db)
	g := models.TestGame(t)
	assert.NoError(t, s.Game().Create(g))
}
func TestGameRepository_GetAll(t *testing.T) {
	db := mongostore.TestDB(t)
	s := mongostore.New(db)
	assert.NotEmpty(t, s.Game().GetAll())
}
