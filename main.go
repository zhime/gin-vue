package main

import (
	"fmt"
	"github.com/zhime/gin-vue/core"
	"github.com/zhime/gin-vue/global"
)

func main() {
	// 读取配置文件
	core.InitConfig()
	fmt.Println(global.Config)
	fmt.Println(global.Config.Mysql.Dsn())

	// 连接数据库
	global.DB = core.InitGorm()
	fmt.Println(global.DB)

	// 日志
	global.Log = core.InitLogger()
	global.Log.Warnln("warn")
	global.Log.Error("error")
	global.Log.Infof("info")
	//r := gin.Default()
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "success")
	//})
	//_ = r.Run()
}
