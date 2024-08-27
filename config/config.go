package config

import (
	"gudang/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:@tcp(localhost)/gudangappdb?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connect to database: %v", err)
	}
}

func AutoMigrate() {
	DB.AutoMigrate(&models.User{},
		&models.Product{},
		&models.Kategori{},
	)
}
