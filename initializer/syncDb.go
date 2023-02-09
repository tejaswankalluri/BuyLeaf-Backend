package initializer

import (
	"fiber-api/models"
	"gorm.io/gorm"
	"log"
)

func SyncDB(DB *gorm.DB) {
	err := DB.AutoMigrate(&models.Book{},
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.ProductInventory{},
		&models.Image{},
	)
	if err != nil {
		log.Fatal("Error while Auto migration: ", err)
	}
}
