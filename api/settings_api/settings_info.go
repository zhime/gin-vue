package settings_api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s SettingApi) SettingsApiView(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":     "xxx",
		"success": "true",
	})
}
