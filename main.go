package main

import (
	"fmt"
	"path/filepath"

	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BindFile struct {
	Name  string                `form:"name" binding:"required"`
	Email string                `form:"email" binding:"required"`
	File  *multipart.FileHeader `form:"file" binding:"required"`
}

func main() {
	//currentPath, _ := os.Getwd()
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.Static("/file", "./public")
	router.LoadHTMLGlob("views/*")
	router.GET("/upload", func(c *gin.Context) {
		//fmt.Println(8 << 20)
		c.HTML(http.StatusOK, "index.html", gin.H{
			//"title": "html demo ~~~~~~~~~~~~~~~",
		})
	})

	router.POST("/upload", func(c *gin.Context) {
		var bindFile BindFile

		// Bind file
		if err := c.ShouldBind(&bindFile); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
			return
		}

		// Save uploaded file
		file := bindFile.File
		dst := filepath.Base(file.Filename)
		//fmt.Println(file.Filename)
		//fmt.Println(dst)
		//fmt.Println(filepath.Dir(file.Filename))
		if err := c.SaveUploadedFile(file, "public/"+dst); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, bindFile.Name, bindFile.Email))
	})

	router.POST("/api/sport", func(c *gin.Context) {
		type Sport struct {
			ID      uint64
			Title   string
			Content string
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    "0000",
			"message": "success",
			"data": []Sport{
				{1, "世界杯开赛啦", "世界杯于明晚8点举行开幕式..."},
				{2, "NBA开赛倒计时5天", "NBA开赛倒计时5天, 万分期待..."},
				{3, "更多精彩...", "更多科技新闻，请持续关注..."},
			},
		})
	})

	router.POST("/api/tech", func(c *gin.Context) {
		type Tech struct {
			ID      uint64
			Title   string
			Content string
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    "0000",
			"message": "success",
			"data": []Tech{
				{1, "5G时代...", "5G时代的到来，让人工智能飞起来..."},
				{2, "互联网大洗牌...", "互联网大洗牌, 一个新时代的到来"},
				{3, "更多精彩...", "更多科技新闻，请持续关注..."},
			},
		})
	})

	router.Run(":8080")
}
