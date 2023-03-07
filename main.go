package main

import (
	"github.com/zhime/gin-vue/core"
	"github.com/zhime/gin-vue/global"
	"github.com/zhime/gin-vue/router"
)

func main() {
	// 读取配置文件
	core.InitConfig()

	// 连接数据库
	global.DB = core.InitGorm()

	// 日志
	global.Log = core.InitLogger()

	// 路由
	r := router.InitRouter()

	addr := global.Config.System.Addr()
	global.Log.Infof("程序运行： %s", addr)
	_ = r.Run(addr)
}
