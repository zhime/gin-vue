package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/zhime/gin-vue/study/gin/router"
)

func InitRouters() *gin.Engine {
	Router := gin.Default()
	apiGroup := Router.Group("api")
	router.InitRouterGroup(apiGroup)

	return Router
}
