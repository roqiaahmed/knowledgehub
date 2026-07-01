package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/roqiaahmed/knowledgehub/internal/database"
	"github.com/roqiaahmed/knowledgehub/internal/handlers"
	"github.com/roqiaahmed/knowledgehub/internal/models"
	"github.com/roqiaahmed/knowledgehub/internal/repositories"
	"github.com/roqiaahmed/knowledgehub/internal/services"
)

func main() {
	envPath := ".env"
	db, err := database.Init(envPath)
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	if err := models.MigrateAll(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	authRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	sqlDB, err := db.DB()
	router := gin.Default()
	router.POST("/auth/register", authHandler.Register)
	router.GET("/health", func(c *gin.Context) {
		if err != nil || sqlDB.Ping() != nil {
			c.JSON(500, gin.H{"status": "error", "database": "disconnected"})
			return
		}
		c.JSON(200, gin.H{
			"status":   "ok",
			"database": "connected",
		})
	})
	router.Run()
}
