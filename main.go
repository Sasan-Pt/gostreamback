package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Configure CORS to allow all origins and headers
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Requested-With", "Accept"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowWildcard = true
	router.Use(cors.New(config))

	slider.sliderImages(router)

	// // Serve static image files
	// router.Static("/images", "./Image")

	router.Run(":8080")
}
