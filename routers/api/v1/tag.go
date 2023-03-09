package v1

import (
	"github.com/GWH/go-gin-example/models"
	"github.com/GWH/go-gin-example/pkg/e"
	"github.com/GWH/go-gin-example/pkg/setting"
	"github.com/GWH/go-gin-example/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["tatol"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddTags(c *gin.Context) {
	var tag models.Tag
	err := c.ShouldBindJSON(&tag)
	if err != nil {
		log.Fatalln(err)
	}

	valid := validation.Validation{}
	valid.Required(tag.Name, "name").Message("名称不能为空")
	valid.MaxSize(tag.Name, 100, "name").Message("名称最长为100字符")
	valid.Required(tag.CreatedBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(tag.CreatedBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(tag.State, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagName(tag.Name) {
			code = e.SUCCESS
			models.AddTag(tag.Name, tag.CreatedBy, tag.State)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	} else {
		for _, e2 := range valid.Errors {
			log.Println(e2)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditTags(c *gin.Context) {

}

func DeleteTags(c *gin.Context) {

}
