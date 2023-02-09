package models

import (
	"gorm.io/gorm"
	"time"
)

type Image struct {
	//gorm.Model

	ID        uint            `json:"ID" gorm:"primary_key;auto_increment;not_null"`
	CreatedAt time.Time       `json:",omitempty"`
	UpdatedAt time.Time       `json:",omitempty"`
	DeletedAt *gorm.DeletedAt `json:",omitempty" gorm:"index"`

	ImageUrl  string `json:"image_url,omitempty" gorm:"not_null"`
	ProductID uint   `json:"-"`
}
