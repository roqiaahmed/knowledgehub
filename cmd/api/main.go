package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/roqiaahmed/knowledgehub/internal/database"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/hi", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	envPath := ".env"
	router.GET("/health", func(c *gin.Context) {
		_, err := database.Init(envPath)
		if err != nil {
			log.Fatalf("Database initialization failed: %v", err)
		}
		c.JSON(200, gin.H{
			"status":   "ok",
			"database": "connected",
		})
	})
	router.Run()
}

// router.GET("/health", func(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"status": "ok",
// 	})
