package config

import (
	// "go-base/utils"

	"log"
	"os"

	"github.com/joho/godotenv"
	// "github.com/joho/godotenv"
)

func EnvCloudName() string {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatal("Error loading app.env file")
	}
	return os.Getenv("CLOUDINARY_CLOUD_NAME")
}

func EnvCloudAPIKey() string {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatal("Error loading app.env file")
	}
	return os.Getenv("CLOUDINARY_API_KEY")
}

func EnvCloudAPISecret() string {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatal("Error loading app.env file")
	}
	return os.Getenv("CLOUDINARY_API_SECRET")
}

func EnvCloudUploadFolder() string {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatal("Error loading app.env file")
	}
	return os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
}
