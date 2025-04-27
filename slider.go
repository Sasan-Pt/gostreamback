package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func SliderImages(router *gin.Engine) {
	router.GET("/images", func(c *gin.Context) {
		files, err := os.ReadDir("./Image")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image directory"})
			return
		}

		var imageURLs []string
		for _, file := range files {
			if !file.IsDir() && isImageFile(file.Name()) {
				imageURLs = append(imageURLs, "http://localhost:8080/images/"+file.Name())
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"images": imageURLs,
			"count":  len(imageURLs),
		})
	})
}

func isImageFile(filename string) bool {
	ext := filepath.Ext(filename)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".webp", ".avif":
		return true
	default:
		return false
	}
}
