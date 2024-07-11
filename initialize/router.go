package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	PublicGroup := Router.Group("/public")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	PrivateGroup := Router.Group("/private")
	{
		PrivateGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	return Router
}
