package routes

import (
	"go-base/apps/controllers"

	"github.com/gin-gonic/gin"
)

type BuildingRouteController struct {
	buildingController controllers.BuildingController
}

func NewBuildingRouteController(buildingController controllers.BuildingController) BuildingRouteController {
	return BuildingRouteController{buildingController}
}

func (bc *BuildingRouteController) BuildingRoute(rg *gin.RouterGroup) {

	r := rg.Group("building")

	r.POST("/", bc.buildingController.CreateBuilding)
	r.GET("/", bc.buildingController.GetBuildings)
	// r.GET("/:postId", pc.postController.PostById)
	// r.DELETE("/:postId", pc.postController.DeletePost)
	// r.PUT("/:postId", pc.postController.UpatePost)

}
