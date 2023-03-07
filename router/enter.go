package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhime/gin-vue/global"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})
	return router
}
