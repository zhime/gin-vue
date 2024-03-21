package model

import "github.com/zhime/gin-vue/server/global"

type BaseMenuBtn struct {
	global.Model
	Name       string `json:"name" gorm:"comment:按钮关键key"`
	Desc       string `json:"desc" gorm:"按钮备注"`
	BaseMenuID uint   `json:"BaseMenuID" gorm:"comment:菜单ID"`
}
