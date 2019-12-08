package models

import (

	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Email     string `json:"email"  validate:"required,email"`
	LastName  string `json:"last_name" validate:"required"`
	Country   string `json:"country" validate:"required"`
	City      string `json:"city" validate:"required"`
	Gender    string `json:"gender" validate:"required"`
	BirthDate string `json:"birth_date" validate:"required"`
}

func (u *User) Validation() error {
	validate := validator.New()
	err := validate.Struct(*u)
	if err != nil {
		return err
	}
	return nil
}
func (u *User)Create() error{

}