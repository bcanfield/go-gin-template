package main

import (
	"go-template/db"
	"go-template/handlers"
	"go-template/middleware"
	"go-template/models"

	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	// Attempt to load .env file
	err := godotenv.Load()
	if err != nil {
		log.Print("Could not load .env file")
	}

	// Open Database
	databaseUrl := os.Getenv("DATABASE_URL")
	database := db.Init(databaseUrl)

	// Migrate
	err = database.AutoMigrate(&models.Comment{}, &models.User{}, &models.Video{})
	if err != nil {
		fmt.Println("db err: (Migrate) ", err)
	}

	// Seed
	db.Seed(database)

	// Setup Router
	router := SetupRouter(database)

	// Start Server
	appPort := os.Getenv("APP_PORT")
	log.Printf("Running Analytics on port %s", appPort)
	router.Run(fmt.Sprintf(":%s", appPort))
}

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.SetDB(db))

	// User routes
	router.GET("/users", handlers.GetUsers)
	router.POST("/users", handlers.CreateUser)
	router.GET("/users/:id", handlers.GetUser)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)

	// Video routes
	router.GET("/videos", handlers.GetVideos)
	router.POST("/videos", handlers.CreateVideo)
	router.GET("/videos/:id", handlers.GetVideo)
	router.PUT("/videos/:id", handlers.UpdateVideo)
	router.DELETE("/videos/:id", handlers.DeleteVideo)

	// Comment routes
	router.GET("/comments", handlers.GetComments)
	router.POST("/comments", handlers.CreateComment)
	router.GET("/comments/:id", handlers.GetComment)
	router.PUT("/comments/:id", handlers.UpdateComment)
	router.DELETE("/comments/:id", handlers.DeleteComment)

	// Health Check route
	router.GET("/api/health", handlers.HealthCheck)

	return router
}
