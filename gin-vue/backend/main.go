package main

import (
	core2 "github.com/zhime/gin-vue/gin-vue/backend/core"
	"github.com/zhime/gin-vue/gin-vue/backend/flag"
	"github.com/zhime/gin-vue/global"
	"github.com/zhime/gin-vue/router"
)

func main() {
	// 读取配置文件
	core2.InitConfig()

	// 连接数据库
	global.DB = core2.InitGorm()

	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	// 日志
	global.Log = core2.InitLogger()

	// 路由
	r := router.InitRouter()

	addr := global.Config.System.Addr()
	global.Log.Infof("程序运行： %s", addr)
	err := r.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
