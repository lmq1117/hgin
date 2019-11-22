package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(favicon.New("./favicon.ico"))
	r.LoadHTMLGlob("views/**/*")
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

	r.GET("/demo/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "html demo ~~~~~~~~~~~~~~~",
		})
	})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "html demo ~~~~~~~~~~~~~~~",
		})
	})

	r.Run(":8080")
}
