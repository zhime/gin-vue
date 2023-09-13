package api

import (
	"github.com/zhime/gin-vue/gin-vue/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingApi
}

var ApiGroupApp = new(ApiGroup)
