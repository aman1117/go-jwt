package initializers

import "github.com/aman1117/go-jwt/models"

func SyncDatabase() {
	// Sync database here
	DB.AutoMigrate(&models.User{})
}
