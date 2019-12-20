package mongostore_test

import (
	"github.com/stretchr/testify/assert"
	"test/internal/app/models"
	"test/internal/app/store/mongostore"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db := mongostore.TestDB(t)
	red := mongostore.TestRedis(t)
	s := mongostore.New(db, red)
	u := models.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u.Email)
}
func TestUserRepository_GetAll(t *testing.T) {
	db := mongostore.TestDB(t)
	red := mongostore.TestRedis(t)
	s := mongostore.New(db, red)
	assert.NotEmpty(t, s.User().GetAll(20))
	assert.Len(t, s.User().GetAll(1), 20)
}
