package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Simple struct to represent a message
type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()

	// Sample data
	messages := []Message{
		{ID: 1, Content: "Hello World"},
		{ID: 2, Content: "Hello from Go!"},
	}

	// GET endpoint to fetch all messages
	r.GET("/messages", func(c *gin.Context) {
		c.JSON(http.StatusOK, messages)
	})

	// POST endpoint to add a new message
	r.POST("/messages", func(c *gin.Context) {
		var newMessage Message
		if err := c.ShouldBindJSON(&newMessage); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		messages = append(messages, newMessage)
		c.JSON(http.StatusCreated, newMessage)
	})

	r.Run(":8080") // Run the server on port 8080
}
