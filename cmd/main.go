package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

	r := gin.Default()
	// Ping test
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "hello world",
			"data": gin.H{
				"content": "pong",
			},
		})
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":5000")
}
