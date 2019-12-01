package main

import (
	"fmt"
	"path/filepath"
	"strings"

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
	//router.Use(Cors())
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

	router.POST("/api/user/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		fmt.Println("xxxxxxxxxxxx", username, password)
		if username == "admin" && password == "123" {
			c.JSON(http.StatusOK, gin.H{
				"code":    "0000",
				"message": "登录成功",
				"token":   "admin",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    "0001",
				"message": "登录失败",
				"token":   "",
			})
		}
	})

	router.POST("/api/user/info/:token", func(c *gin.Context) {
		token := c.Param("token")
		type UserInfo struct {
			ID    uint
			Name  string
			Roles []string
		}
		if token == "admin" {
			c.JSON(http.StatusOK, gin.H{
				"code":    "0000",
				"message": "获取用户信息成功",
				"data":    UserInfo{1, "李雷", []string{"manager", "dev"}},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    "0001",
				"message": "获取用户信息失败",
				"data":    nil,
			})
		}
	})

	router.Run(":8887")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求

	}
}
