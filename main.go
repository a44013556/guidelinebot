package main

import (
	"linebot-gin/config"
	"linebot-gin/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	config.DB.AutoMigrate(&models.Booking{})

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)

}
