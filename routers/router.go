package routers

import (
	"github.com/gin-gonic/gin"
	"hgin/pkg/setting"
	"hgin/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	appv1 := r.Group("/api/v1")
	{
		appv1.GET("/tags", v1.GetTags)
		appv1.POST("/tags", v1.AddTag)
		appv1.PUT("/tags/:id", v1.EditTag)
		appv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
