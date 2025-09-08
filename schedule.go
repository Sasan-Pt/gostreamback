package main

import (
	"net/http"
	"streamSite/db"

	"github.com/gin-gonic/gin"
)

func Schedule(router *gin.Engine) {
	router.GET("/schedule", func(c *gin.Context) {
		startDate := c.Query("startDate")
		endDate := c.Query("endDate")
		schedule, err := db.GetEpisodes(startDate, endDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch data from database: " + err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": schedule,
		})
	})

}
