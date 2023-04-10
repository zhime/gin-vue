package main

import (
	"github.com/zhime/gin-vue/study/gin/router"
)

func main() {
	r := router.Routers()

	_ = r.Run(":8080")
}
