package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	//gorm.Model

	ID        uint            `json:"ID" gorm:"primary_key;auto_increment;not_null"`
	CreatedAt time.Time       `json:",omitempty"`
	UpdatedAt time.Time       `json:",omitempty"`
	DeletedAt *gorm.DeletedAt `json:",omitempty" gorm:"index"`

	ProductName   string  `json:"product_name" gorm:"uniqueIndex"`
	ProductDesc   string  `json:"product_desc"`
	ProductPrice  float64 `json:"product_price"`
	ProductImages []Image `json:"product_images"`

	//Category []Category `json:"category,omitempty" gorm:"many2many:product_category;"`
	Category         *Category         `json:"category,omitempty"`
	CategoryID       uint              `json:"category_id"`
	ProductInventory *ProductInventory `json:"inventory,omitempty" gorm:"OnDelete:SET NULL;"`
	//ProductInventoryID uint              `json:"product_inventory_id,omitempty"`
}

type Category struct {
	//gorm.Model

	ID        uint            `json:"ID" gorm:"primary_key;auto_increment;not_null"`
	CreatedAt time.Time       `json:",omitempty"`
	UpdatedAt time.Time       `json:",omitempty"`
	DeletedAt *gorm.DeletedAt `json:",omitempty" gorm:"index"`

	Name string `json:"name,omitempty" gorm:"uniqueIndex"`
	Desc string `json:"desc,omitempty"`
	//ProductID uint            `json:"-"`
	//Product *[]Product `json:"product,omitempty"`
}

type ProductInventory struct {
	//gorm.Model

	ID        uint            `json:"ID" gorm:"primary_key;auto_increment;not_null"`
	CreatedAt time.Time       `json:",omitempty"`
	UpdatedAt time.Time       `json:",omitempty"`
	DeletedAt *gorm.DeletedAt `json:",omitempty" gorm:"index"`

	Quantity  int  `json:"quantity"`
	ProductID uint `json:"product_id,omitempty" gorm:"uniqueIndex" `
}
