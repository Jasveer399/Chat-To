package database

import "github.com/Jasveer399/Chat-To/models"

func Migrate() {
	DB.AutoMigrate(&models.Message{})
	DB.AutoMigrate(&models.User{})
}
