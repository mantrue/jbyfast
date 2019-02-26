package routers

import (
	"github.com/gin-gonic/gin"

	"gin-blog/middleware/jwt"
	"gin-blog/middleware/tollbooth_gin"
	"gin-blog/pkg/setting"
	"gin-blog/routers/api"
	"gin-blog/routers/api/v1"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"net/http"
	"time"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.StaticFS("/static", http.Dir("./static")) //加载静态资源

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	lmt := tollbooth.NewLimiter(setting.MaxRequestNum, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Minute})
	lmt.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"}) //初始化IP访问控制

	r.GET("/auth", tollbooth_gin.LimitHandler(lmt), api.GetAuth) //接入限速中间件

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT()) //接入jwt中间件
	//	apiv1.Use(tollbooth_gin.LimitHandler(lmt)) //接入限速中间件

	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
