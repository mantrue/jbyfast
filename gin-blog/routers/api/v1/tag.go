package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/astaxie/beego/validation"
	"github.com/Unknwon/com"

	"gin-blog/gredis"
	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	"gin-blog/service/cache_service"
	"gin-blog/service/mysql_service"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	data["lists"] = mysql_service.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = mysql_service.GetTagTotal(maps)

	//redis操作逻辑
	tg := cache_service.Tag{100, "penghui", 1, 1, 1}
	gredis.Set(tg.GetTagsKey(), tg.GetTagsKey(), 60)
	data["rediskey"] = tg.GetTagsKey()

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": data,
	})
}

//新增文章标签
func AddTag(c *gin.Context) {
	var t models.Tag
	t.Name = c.Query("name")
	id, err := mysql_service.AddTag(&t)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"msg":  e.GetMsg(e.SUCCESS),
			"data": map[string]int{"id": id},
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": map[string]int{"id": id},
	})
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}
