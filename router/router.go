package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {

	//Initialize router
	router := gin.Default()

	//Initialize Routes
	initializeRoutes(router)

	//Get the port from the enviroment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run("0.0.0.0:" + port) // listen and serve on 0.0.0.0:8080
}
