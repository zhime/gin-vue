package router

import (
	"github.com/zhime/gin-vue/gin-vue/api"
)

//func SettingsRouter(router *gin.Engine) {
//	settingsApi := api.ApiGroupApp.SettingsApi
//	router.GET("", settingsApi.SettingsApiView)
//}

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("settings", settingsApi.SettingsApiView)
}
