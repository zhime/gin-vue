package main

import (
	"github.com/zhime/gin-vue/core"
	"github.com/zhime/gin-vue/global"
	"github.com/zhime/gin-vue/initialize"
)

func main() {
	global.G_VIPER = core.Viper()
	r := initialize.Routers()
	_ = r.Run(":8080")
}
