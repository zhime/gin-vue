package model

import "github.com/zhime/gin-vue/global"

type BaseMenuBtn struct {
	global.MODEL
	Name          string `json:"name" gorm_study:"comment:按钮关键key"`
	Desc          string `json:"desc" gorm_study:"按钮备注"`
	SysBaseMenuID uint   `json:"sysBaseMenuID" gorm_study:"comment:菜单ID"`
}
