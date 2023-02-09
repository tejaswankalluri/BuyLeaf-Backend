package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	//gorm.Model

	ID        uint            `json:"ID" gorm:"primary_key;auto_increment;not_null"`
	CreatedAt time.Time       `json:",omitempty"`
	UpdatedAt time.Time       `json:",omitempty"`
	DeletedAt *gorm.DeletedAt `json:",omitempty" gorm:"index"`

	Username string `json:"username,omitempty" gorm:"uniqueIndex" validate:"required,min=3,max=40"`
	Email    string `json:"email,omitempty" gorm:"uniqueIndex" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=8,max=40"`
	Role     string `json:"role" gorm:"default:USER"`
}
