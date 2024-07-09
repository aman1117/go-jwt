package initializers

import "github.com/joho/godotenv"

func LoadEnvVariables() {
	// Load environment variables here
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
