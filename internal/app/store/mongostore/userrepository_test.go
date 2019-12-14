package mongostore_test

import (
	"github.com/stretchr/testify/assert"
	"test/internal/app/models"
	"test/internal/app/store/mongostore"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db := mongostore.TestDB(t)
	s := mongostore.New(db)
	u := models.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u.Email)
}
func TestUserRepository_GetAll(t *testing.T) {
	db := mongostore.TestDB(t)
	s := mongostore.New(db)
	assert.NotEmpty(t, s.User().GetAll())
	assert.Len(t, s.User().GetAll(), 25)
}
