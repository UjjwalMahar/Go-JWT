package initializers

import "github.com/UjjwalMahar/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}