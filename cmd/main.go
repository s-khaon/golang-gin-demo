package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	r.Run()
}
