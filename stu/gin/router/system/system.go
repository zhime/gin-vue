package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SysRouter struct {
}

func (s SysRouter) InitSysRouter(router *gin.RouterGroup) {
	sysRouter := router.Group("system")
	{
		sysRouter.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "success",
				"data": "system",
			})
		})
	}
}
