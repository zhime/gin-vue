package main

import (
	"fmt"
	"github.com/zhime/gin-vue/core"
	"github.com/zhime/gin-vue/global"
)

func main() {
	global.G_VIPER = core.Viper()
	fmt.Println(global.G_CONFIG)
	logger := core.Zap()
	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
	logger.Fatal("fatal")
	//r := initialize.Routers()
	//_ = r.Run(":8080")
}
