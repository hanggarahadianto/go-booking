package main

import (
	"fmt"
	"go-base/db"
	"go-base/utils"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func init() {
	// AuthRouteController = authroutes.NewAuthRouteController(AuthController)
	// PostRouteController = postroutes.NewRoutePostController(PostController)
	gin.SetMode(gin.ReleaseMode)
	server = gin.Default()
}

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables ", err)
	}

	db.InitializeDb(&config)

	fmt.Println("server running on port " + config.ServerPort)
	log.Fatal(server.Run(":" + config.ServerPort)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
