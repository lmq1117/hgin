package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/some/json", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "Go语言",
			"tag":  "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.Run()
}
