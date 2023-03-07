package main

import (
	"fmt"
	"github.com/zhime/gin-vue/core"
	"github.com/zhime/gin-vue/global"
)

func main() {

	core.InitConfig()
	fmt.Println(global.Config)
	//r := gin.Default()
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "success")
	//})
	//_ = r.Run()
}
