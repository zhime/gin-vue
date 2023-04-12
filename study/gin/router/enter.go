package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhime/gin-vue/study/gin/api"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := RouterGroupApp.System.SysRouter
	PrivateGroup := Router.Group("")
	{
		systemRouter.InitSysRouter(PrivateGroup)
	}
	v1 := Router.Group("/v1")
	{
		v1.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "success",
				"data": "login",
			})
		})

		v1.GET("/list/:id/:action", func(c *gin.Context) {
			id := c.Param("id")
			action := c.Param("action")
			c.JSON(http.StatusOK, gin.H{
				"msg":    "success",
				"data":   "list",
				"id":     id,
				"action": action,
			})
		})

		v1.GET("/sendsms", api.ALiYunSms)

		//v1.GET("/list/:id/*action", func(c *gin.Context) {
		//	id := c.Param("id")
		//	action := c.Param("action")
		//	c.JSON(http.StatusOK, gin.H{
		//		"msg":    "success",
		//		"data":   "list",
		//		"id":     id,
		//		"action": action,
		//	})
		//})
	}
	return Router
}
