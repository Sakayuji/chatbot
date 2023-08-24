package main

import (
	"github.com/Sakayuji/chatbot/api"
	"github.com/Sakayuji/chatbot/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get messages
	r.POST("/SendMessage", api.SendMessage)

	// Create Template
	r.POST("/CreateTemplate", api.CreateTemplate)

	// Get messages
	r.GET("/GetMessages", api.GetMessage)

	// Get templates
	r.GET("/GetTemplates", api.GetTemplate)

	return r
}

func main() {
	db.DB = db.GetDb()
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")

}
