package models

import "testing"

func TestUser(t *testing.T) *User {
	t.Helper()
	return &User{
		Email:     "ivan@example.com",
		LastName:  "Ivanov",
		Country:   "Ukraine",
		City:      "Kyiv",
		Gender:    "male",
		BirthDate: "Friday, April 4, 8527 8:45 AM",
	}
}

func TestGame(t *testing.T) *Game {
	t.Helper()
	return &Game{
		PointsGained: "677",
		WinStatus:    "0",
		GameType:     "11",
		Created:      "8/17/2019 8:54 PM",
	}
}
