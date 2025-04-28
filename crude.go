package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var people = []Person{
	{ID: "1", Name: "Abbed"},
	{ID: "2", Name: "Ahmad"},
	{ID: "3", Name: "Ghassan"},
}

func CRUDRoutes(router *gin.Engine) {
	// Create - POST /people
	router.POST("/people", func(c *gin.Context) {
		var newPerson Person
		if err := c.BindJSON(&newPerson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		people = append(people, newPerson)
		c.JSON(http.StatusCreated, newPerson)
	})

	// Read - GET /people
	router.GET("/people", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": people,
		})
	})

	// Read Single - GET /people/:id
	router.GET("/people/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, person := range people {
			if person.ID == id {
				c.JSON(http.StatusOK, person)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
	})

	// Update - PUT /people/:id
	router.PUT("/people/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedPerson Person
		if err := c.BindJSON(&updatedPerson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i, person := range people {
			if person.ID == id {
				updatedPerson.ID = id // Ensure ID remains the same
				people[i] = updatedPerson
				c.JSON(http.StatusOK, updatedPerson)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
	})

	// Delete - DELETE /people/:id
	router.DELETE("/people/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, person := range people {
			if person.ID == id {
				people = append(people[:i], people[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Person deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
	})
}
