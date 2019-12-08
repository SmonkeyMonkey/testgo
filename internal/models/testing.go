package models

import "testing"

func TestUser(t *testing.T) *User {
	t.Helper()
	return &User{
		Email:     "test@example.com",
		LastName:  "Ivanov",
		Country:   "Ukraine",
		City:      "Kyiv",
		Gender:    "male",
		BirthDate: "Friday, April 4, 8527 8:45 AM",
	}
}