package main

import (
	"github.com/zhime/gin-vue/initialize"
)

func main() {
	r := initialize.Routers()
	_ = r.Run(":8080")
}
