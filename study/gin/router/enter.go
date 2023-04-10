package router

import (
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := RouterGroupApp.System
	PrivateGroup := Router.Group("")
	{
		systemRouter.SysRouter.InitSysRouter(PrivateGroup)
	}
	return Router
}
