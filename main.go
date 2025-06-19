package main

import(
	"github.com/gin-gonic/gin"
    "linebot/config"
    "linebot/models"
    "net/http"
    "os"
)


func main() {
	config.InitDB()
	config.DB.AutoMigrate(^models.Booking{})

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context)) {
		c.JSON(http.StatusOK, gin.H{"messga":"pong"})
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)

}