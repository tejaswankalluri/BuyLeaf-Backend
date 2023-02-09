package util

import (
	"fiber-api/models"
)

func SanitizeUserModel(u *models.User) {
	//return models.User{gorm.Model{ID: u.ID, CreatedAt: u.CreatedAt, UpdatedAt: u.UpdatedAt, DeletedAt: u.DeletedAt}, u.Username, u.Email, ""}
	u.Password = ""
}
