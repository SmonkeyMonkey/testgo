package models_test

import (
	"github.com/stretchr/testify/assert"
	"test/internal/app/models"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	testcase := []struct {
		name     string
		u        func() *models.User
		is_valid bool
	}{
		{
			name: "valid",
			u: func() *models.User {
				return models.TestUser(t)
			},
			is_valid: true,
		},
		{
			name: "empty email",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Email = ""
				return u
			},
			is_valid: false,
		},
		{
			name: "incorrect email",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Email = "email"
				return u
			},
			is_valid: false,
		},
		{
			name: "empty last name",
			u: func() *models.User {
				u := models.TestUser(t)
				u.LastName = ""
				return u
			},
			is_valid: false,
		},
		{
			name: "empty country",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Country = ""
				return u
			},
			is_valid: false,
		},
		{
			name: "empty city",
			u: func() *models.User {
				u := models.TestUser(t)
				u.City = ""
				return u
			},
			is_valid: false,
		},
		{
			name: "empty gender",
			u: func() *models.User {
				u := models.TestUser(t)
				u.Gender = ""
				return u
			},
			is_valid: false,
		},
		{
			name: "empty birth date",
			u: func() *models.User {
				u := models.TestUser(t)
				u.BirthDate = ""
				return u
			},
			is_valid: false,
		},
	}
	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			if tc.is_valid {
				assert.NoError(t, tc.u().Validation())
			} else {
				assert.Error(t, tc.u().Validation())
			}
		})
	}
}
