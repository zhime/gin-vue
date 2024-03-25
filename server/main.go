package main

import (
	"github.com/zhime/gin-vue/server/core"
	"github.com/zhime/gin-vue/server/global"
)

func main() {
	core.InitViper()
	core.InitZap()
	global.Zap.Info("test")
	core.InitGorm()
	//fmt.Println(global.Db)

	//r := gin.Default()
	//systemRouter := router.RouterGroupApp.System
	//
	//systemRouter.InitUserRouter(r)

}
