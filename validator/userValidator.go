package validator

import (
	"fiber-api/models"
)

func ValidateUser(user models.User) []*ErrorResponse {
	err := GenValidate(user)
	return err
}

type Logininput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=40"`
}
type RefreshToken struct {
	Refreshtoken string `json:"refresh_token" validate:"required,jwt"`
}

func ValidateLogin(login Logininput) []*ErrorResponse {
	err := GenValidate(login)
	return err
}

func ValidateRefreshToken(token RefreshToken) []*ErrorResponse {
	err := GenValidate(token)
	return err
}
