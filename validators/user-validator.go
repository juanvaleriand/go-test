package validators

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

type UserInput struct {
	Email      string `json:"email" validate:"required"`
	First_name string `json:"first_name" validate:"required"`
	Last_name  string `json:"last_name" validate:"required"`
	Avatar     string `json:"avatar" validate:"required"`
}
