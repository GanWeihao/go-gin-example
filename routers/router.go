package routers

import (
	"github.com/GWH/go-gin-example/pkg/setting"
	v1 "github.com/GWH/go-gin-example/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	tagApiV1 := r.Group("/api/v1")
	{
		// 获取标签列表
		tagApiV1.GET("/tags", v1.GetTags)
		// 新建标签
		tagApiV1.POST("/tags", v1.AddTags)
		// 修改指定标签
		tagApiV1.PUT("/tags/:id", v1.EditTags)
		// 删除指定标签
		tagApiV1.DELETE("/tags/:id", v1.DeleteTags)
	}

	return r
}
