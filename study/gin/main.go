package main

import (
	"github.com/zhime/gin-vue/study/gin/initialize"
	"go.uber.org/zap"
)

func main() {
	//r := router.Routers()
	//r := gin.Default()
	//apiGroup := r.Group("api")
	//router.InitRouterGroup(apiGroup)

	r := initialize.InitRouters()

	logger, _ := zap.NewProduction()
	logger.Info("服务启动")

	//_ = r.Run()

	if err := r.Run(":8080"); err != nil {
		//logger.Panic("启动失败")
		panic("启动失败")
	}
}
