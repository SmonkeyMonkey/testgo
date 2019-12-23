package models

import (
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Email     string `json:"email"  validate:"required,email" bson:"email"`
	LastName  string `json:"last_name"  validate:"required" bson:"last_name"`
	Country   string `json:"country" validate:"required" bson:"country"`
	City      string `json:"city" validate:"required" bson:"city"`
	Gender    string `json:"gender" validate:"required" bson:"gender"`
	BirthDate string `json:"birth_date" validate:"required" bson:"birth_date"`
}

func (u *User) Validation() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return err
	}
	return nil
}
