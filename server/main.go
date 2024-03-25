package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhime/gin-vue/server/core"
	"github.com/zhime/gin-vue/server/global"
	"github.com/zhime/gin-vue/server/router"
)

func main() {
	core.InitViper()
	core.InitZap()
	global.Zap.Info("test")
	core.InitGorm()
	//fmt.Println(global.Db)

	r := gin.Default()
	//systemRouter := router.RouterGroupApp.System
	//
	//systemRouter.InitUserRouter(r)
	api := r.Group("api")
	router.RouterGroupApp.System.UserRouter.InitUserRouter(api)

	_ = r.Run()

}
