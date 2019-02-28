package tollbooth_gin

import (
	"gin-blog/pkg/e"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LimitHandler(lmt *limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			//c.Data(httpError.StatusCode, lmt.GetMessageContentType(), []byte(httpError.Message))
			c.JSON(http.StatusForbidden, gin.H{
				"code": e.SUCCESS,
				"msg":  e.GetMsg(e.FORBIDDEN),
				"data": map[string]string{},
			})

			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
