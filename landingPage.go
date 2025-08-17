package main

import (
	"net/http"
	"streamSite/db"

	"github.com/gin-gonic/gin"
)

func LANDINGPAGE(router *gin.Engine) {
	// Get all manPage records
	router.GET("/mainPage", func(c *gin.Context) {
		// Query the manPage table from the database
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

	// Get a specific manPage by ID
	// router.GET("/mainPage/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	// Convert string ID to int (you might want to add validation)
	// 	var intID int
	// 	if _, err := fmt.Sscanf(id, "%d", &intID); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Invalid ID format",
	// 		})
	// 		return
	// 	}

	// 	mainPage, err := db.GetManPageByID(intID)
	// 	if err != nil {
	// 		c.JSON(http.StatusNotFound, gin.H{
	// 			"error": "Record not found",
	// 		})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"data": mainPage,
	// 	})
	// })

	// // Create a new manPage record
	// router.POST("/mainPage", func(c *gin.Context) {
	// 	var request struct {
	// 		Name string `json:"name" binding:"required"`
	// 	}

	// 	if err := c.ShouldBindJSON(&request); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Invalid request data: " + err.Error(),
	// 		})
	// 		return
	// 	}

	// 	manPage, err := db.CreateManPage(request.Name)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"error": "Failed to create record: " + err.Error(),
	// 		})
	// 		return
	// 	}

	// 	c.JSON(http.StatusCreated, gin.H{
	// 		"data": manPage,
	// 	})
	// })

	// // Update an existing manPage record
	// router.PUT("/mainPage/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	var intID int
	// 	if _, err := fmt.Sscanf(id, "%d", &intID); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Invalid ID format",
	// 		})
	// 		return
	// 	}

	// 	var request struct {
	// 		Name string `json:"name" binding:"required"`
	// 	}

	// 	if err := c.ShouldBindJSON(&request); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Invalid request data: " + err.Error(),
	// 		})
	// 		return
	// 	}

	// 	if err := db.UpdateManPage(intID, request.Name); err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"error": "Failed to update record: " + err.Error(),
	// 		})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Record updated successfully",
	// 	})
	// })

	// // Delete a manPage record
	// router.DELETE("/mainPage/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	var intID int
	// 	if _, err := fmt.Sscanf(id, "%d", &intID); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "Invalid ID format",
	// 		})
	// 		return
	// 	}

	// 	if err := db.DeleteManPage(intID); err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"error": "Failed to delete record: " + err.Error(),
	// 		})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Record deleted successfully",
	// 	})
	// })
}