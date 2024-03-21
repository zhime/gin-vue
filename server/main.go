package main

import (
	"fmt"
	"github.com/zhime/gin-vue/server/core"
	"github.com/zhime/gin-vue/server/global"
)

func main() {
	core.InitViper()
	core.InitZap()
	global.Zap.Info("test")
	core.InitGorm()
	fmt.Println(global.Db)
}
