package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhime/gin-vue/study/gin/api"
)

func InitRouterGroup(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", api.GerUserList)
	}
}