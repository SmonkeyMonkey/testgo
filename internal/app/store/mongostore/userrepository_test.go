package mongostore_test

import (
	"github.com/stretchr/testify/assert"
	"test/internal/app/models"
	"testing"
)

var (
	u = models.TestUser(&t)
)

func TestUserRepository_Create(t *testing.T) {
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u.Email)
}
func TestUserRepository_GetAll(t *testing.T) {
	assert.NotEmpty(t, s.User().GetAll(1))
	assert.Len(t, s.User().GetAll(1), 30)
}
func BenchmarkUserRepository_Create(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.User().Create(u)
	}
}
func BenchmarkUserRepository_GetAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.User().GetAll(1)
	}
}
