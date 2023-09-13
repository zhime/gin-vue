package router

import "github.com/zhime/gin-vue/study/gin/router/system"

type RouteGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouteGroup)
