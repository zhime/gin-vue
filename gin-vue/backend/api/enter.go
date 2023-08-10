package api

import (
	"github.com/zhime/gin-vue/gin-vue/backend/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingApi
}

var ApiGroupApp = new(ApiGroup)
