package controllers

import (
	"context"
	"go-base/apps/models"
	"go-base/config"
	"go-base/db"
	"log"
	"net/http"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BuildingController struct {
	DB *gorm.DB
}

func (bc *BuildingController) GetBuildings(c *gin.Context) {
	var buildingList []models.Building

	result := db.DB.Debug().Find(&buildingList)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   buildingList,
	})

}

func (bc *BuildingController) CreateBuilding(c *gin.Context) {

	// currentUser:= c.MustGet("currentUser").(models.Building)

	// var buildingData models.CreateBuildingInput

	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "no file uploaded",
			"error":  err.Error(),
		})
	}

	ctx := context.Background()
	cld, err := cloudinary.NewFromParams(config.EnvCloudName(), config.EnvCloudAPIKey(), config.EnvCloudAPISecret())
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"status": " cant get caludinary account",
			"error":  err.Error(),
		})
		return
	}

	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: config.EnvCloudUploadFolder()})
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"status": " cant upload",
			"error":  err.Error(),
		})
		return
	}
	log.Println(resp.SecureURL)

	now := time.Now()
	newBuilding := models.Building{
		Name:    c.Request.PostFormValue("name"),
		Content: c.Request.PostFormValue("content"),
		// User:      buildingtUser.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	newBuilding.Image = resp.SecureURL

	result := db.DB.Create(&newBuilding)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   newBuilding,
	})
}
