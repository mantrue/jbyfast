package controller

import "github.com/gin-gonic/gin"

type Persion struct {
}

func (p *Persion) StartPage(c *gin.Context) {
	c.String(200, "chengpenghui")
}
