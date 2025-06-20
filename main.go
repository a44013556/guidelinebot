package main

import (
	"guidelinebot/config"
	"guidelinebot/handlers"
	"guidelinebot/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	config.InitDB()
	config.DB.AutoMigrate(&models.Booking{}, &models.JapanArea{}, &models.AreaSpot{})

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.POST("/linewebhook", handlers.LineWebhookHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)

}
