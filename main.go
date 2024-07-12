package main

import (
	"fmt"
	"github.com/zhime/gin-vue/core"
	"github.com/zhime/gin-vue/global"
)

func main() {
	global.G_VIPER = core.Viper()
	fmt.Println(global.G_CONFIG)
	//r := initialize.Routers()
	//_ = r.Run(":8080")
}
