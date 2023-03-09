package routers

import (
	"github.com/GWH/go-gin-example/pkg/e"
	"github.com/GWH/go-gin-example/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(e.SUCCESS, gin.H{
			"message": "test",
		})
	})
	return r
}
