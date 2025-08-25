package main

import (
	"streamSite/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	if err := db.InitDB(); err != nil {
		panic("Failed to initialize database: " + err.Error())
	}

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

	// Add routes
	SliderImages(router)
	CRUDRoutes(router) // Add the CRUD routes
	LANDINGPAGE(router)

	// Serve static image files
	router.Static("/images", "./Image")

	router.Run(":8080")
}

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"

// 	_ "github.com/mattn/go-sqlite3"
// )

// func main() {

// 	db, err := sql.Open("sqlite3", "./api.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()


// 	imageFolder := "./image"


// 	files, err := os.ReadDir(imageFolder)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

	
// 	baseURL := "http://localhost:8080/images/"


// 	for _, file := range files {
// 		if !file.IsDir() {
			
// 			imageURL := baseURL + file.Name()

		
// 			name := file.Name()                     
// 			summary := "auto imported image"       
// 			link := imageURL                       

			
// 			_, err := db.Exec(`
// 				INSERT INTO RecentUploads (name, summary, link) 
// 				VALUES (?, ?, ?)
// 			`, name, summary, link)
// 			if err != nil {
// 				log.Println("❌ Failed to insert:", file.Name(), err)
// 			} else {
// 				fmt.Println("✅ Inserted:", imageURL)
// 			}
// 		}
// 	}
// }