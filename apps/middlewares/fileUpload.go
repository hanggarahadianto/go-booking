package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FileUploadMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad request to upload",
			})
			return
		}
		defer file.Close() // close file properly

		c.Set("filePath", header.Filename)
		c.Set("file", file)

		c.Next()
	}
}

// var (
// 	validate = validator.New()
// )

// type mediaUpload interface {
// 	FileUpload(building models.Building) (string, error)
// }

// type media struct{}

// func NewMediaUpload() mediaUpload {
// 	return &media{}
// }

// func (*media) FileUpload(building models.Building) (string, error) {
// 	//validate
// 	// err := validate.Struct(building)
// 	// if err != nil {
// 	// 	return "", err
// 	// }

// 	//upload
// 	// uploadUrl, err := utils.ImageUploadHelper(building.Image)
// 	uploadUrl, err := utils.UploadToCloudinary(building.Image)
// 	if err != nil {
// 		return "", err
// 	}

// 	return uploadUrl, nil
// }
