package main

import (
	"net/http"
	"streamSite/db"

	"github.com/gin-gonic/gin"
)

func Schedule(router *gin.Engine) {
	router.GET("/schedule", func(c *gin.Context) {
		mainPage, err := db.GetAllManPages()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch data from database: " + err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": mainPage,
		})
	})

}
