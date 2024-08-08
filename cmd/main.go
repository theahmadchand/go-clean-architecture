package main

import (
	"github.com/gin-gonic/gin"
	"github.com/theahmadchand/go-clean-architecture/internal/adapters/http"
	"github.com/theahmadchand/go-clean-architecture/internal/adapters/repositories"
	"github.com/theahmadchand/go-clean-architecture/internal/config"
	"github.com/theahmadchand/go-clean-architecture/internal/infrastructure"
	"github.com/theahmadchand/go-clean-architecture/internal/usecases"
	"log"
)

func main() {
	engine := gin.Default()

	// Initialize database connection
	db, err := infrastructure.NewDatabaseConnection()
	if err != nil {
		log.Fatalf("Error initializing database connection: %v", err)
	}

	// Set up repositories
	postRepo := repositories.NewPostRepository(db)

	// Set up use cases
	postUseCase := usecases.NewPostUseCase(postRepo)

	// Set up HTTP handlers
	postHandler := http.NewPostHandler(postUseCase)

	// Define routes
	config.SetupRoutes(engine, postHandler)

	// Start the server
	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}