package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhime/gin-vue/gin-vue/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	//SettingsRouter(router)
	apiRouterGroup := router.Group("api")
	apiRouterGroupApp := RouterGroup{apiRouterGroup}
	apiRouterGroupApp.SettingsRouter()

	return router
}
