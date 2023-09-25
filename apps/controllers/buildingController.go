package controllers

import (
	"go-base/apps/models"
	"go-base/db"
	"net/http"
	"time"

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

	var buildingData models.CreateBuildingInput

	err := c.ShouldBindJSON(&buildingData)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newBuilding := models.Building{
		Name:     buildingData.Name,
		Content:  buildingData.Content,
		ImageURL: buildingData.ImageURL,
		// User:      buildingtUser.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}

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
