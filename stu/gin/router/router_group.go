package router

import "github.com/zhime/gin-vue/stu/gin/router/system"

type RouteGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouteGroup)
