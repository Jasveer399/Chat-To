package database

import "github.com/Jasveer399/web-service-gin/models"

func Migrate() {
	DB.AutoMigrate(&models.Message{})
	DB.AutoMigrate(&models.User{})
}
