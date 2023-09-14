package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhime/gin-vue/stu/gin/api"
)

func InitRouterGroup(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", api.GerUserList)
		UserRouter.GET("sendsms", api.ALiYunSms)
	}
}
