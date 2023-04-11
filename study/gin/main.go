package main

import (
	"github.com/zhime/gin-vue/study/gin/initialize"
)

func main() {
	//r := router.Routers()
	//r := gin.Default()
	//apiGroup := r.Group("api")
	//router.InitRouterGroup(apiGroup)

	r := initialize.InitRouters()

	//zap.S().Infof("服务启动")

	if err := r.Run(":8080"); err != nil {
		//zap.S().Panic("启动失败:", err.Error())
		panic("启动失败")
	}
}
