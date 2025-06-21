package main

import (
	"guidelinebot/config"
	"guidelinebot/handlers/linebot"
	"guidelinebot/models"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.Init()
	if err != nil {
		log.Fatal("Init Fail:", err)
	}
	if err := cfg.DB.AutoMigrate(&models.Booking{}, &models.JapanArea{}, &models.AreaSpot{}); err != nil {
		log.Fatalf("AutoMigrate error: %v", err)
	}
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	lineHandler := linebot.NewHandler(cfg.DB, cfg.RDB)
	r.POST("/linewebhook", lineHandler.LineWebhookHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)

}
