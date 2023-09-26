package utils

// import (
// 	"context"
// 	"go-base/config"
// 	"time"

// 	"github.com/cloudinary/cloudinary-go/v2"
// 	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
// )

// // func UploadToCloudinary(file multipart.File, filePath string) (string, error) {
// func UploadToCloudinary(file interface{}) (string, error) {
// 	// file := middlewares.FileUploadMiddleware()
// 	ctx := context.Background()
// 	cld, err := cloudinary.NewFromParams(config.EnvCloudName(), config.EnvCloudAPIKey(), config.EnvCloudAPISecret())
// 	if err != nil {
// 		return "", err
// 	}
// 	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: config.EnvCloudUploadFolder()})

// 	if err != nil {
// 		return "", err
// 	}

// 	image := resp.SecureURL
// 	return image, nil
// }

// func ImageUploadHelper(input interface{}) (string, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	//create cloudinary instance
// 	cld, err := cloudinary.NewFromParams(config.EnvCloudName(), config.EnvCloudAPIKey(), config.EnvCloudAPISecret())
// 	if err != nil {
// 		return "", err
// 	}

// 	//upload file
// 	resp, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: config.EnvCloudUploadFolder()})
// 	if err != nil {
// 		return "", err
// 	}

// 	return resp.SecureURL, nil
// }
