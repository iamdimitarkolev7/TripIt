package main

import (
	"os"
	"tripit/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.SetupRoutes(router)

	router.Run(":" + port)
}
