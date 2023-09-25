package main

import (
	"fmt"
	"go-base/apps/controllers"
	"go-base/db"
	"go-base/utils"
	"log"

	buildingroutes "go-base/routes"

	"github.com/gin-gonic/gin"
)

var (
	server                  *gin.Engine
	BuildingController      controllers.BuildingController
	BuildingRouteController buildingroutes.BuildingRouteController
)

func init() {
	// AuthRouteController = authroutes.NewAuthRouteController(AuthController)
	BuildingRouteController = buildingroutes.NewBuildingRouteController(BuildingController)
	gin.SetMode(gin.ReleaseMode)
	server = gin.Default()
}

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables ", err)
	}

	db.InitializeDb(&config)

	router := server.Group("/api")
	BuildingRouteController.BuildingRoute(router)

	fmt.Println("server running on port " + config.ServerPort)
	log.Fatal(server.Run(":" + config.ServerPort)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
