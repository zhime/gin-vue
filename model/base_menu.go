package model

import "github.com/zhime/gin-vue/global"

type BaseMenu struct {
	global.MODEL
	MenuLevel uint   `json:"-"`
	ParentId  string `json:"parentId" gorm_study:"comment:父菜单ID"`     // 父菜单ID
	Path      string `json:"path" gorm_study:"comment:路由path"`        // 路由path
	Name      string `json:"name" gorm_study:"comment:路由name"`        // 路由name
	Hidden    bool   `json:"hidden" gorm_study:"comment:是否在列表隐藏"`     // 是否在列表隐藏
	Component string `json:"component" gorm_study:"comment:对应前端文件路径"` // 对应前端文件路径
	Sort      int    `json:"sort" gorm_study:"comment:排序标记"`          // 排序标记
	//Meta       `json:"meta" gorm_study:"embedded;comment:附加属性"`                     // 附加属性
	//Authoritys []Authority         `json:"authoritys" gorm_study:"many2many:authority_menus;"`
	Children []BaseMenu `json:"children" gorm_study:"-"`
	//Parameters []BaseMenuParameter `json:"parameters"`
	//MenuBtn    []BaseMenuBtn       `json:"menuBtn"`
}

type Meta struct {
	ActiveName  string `json:"activeName" gorm_study:"comment:高亮菜单"`
	KeepAlive   bool   `json:"keepAlive" gorm_study:"comment:是否缓存"`           // 是否缓存
	DefaultMenu bool   `json:"defaultMenu" gorm_study:"comment:是否是基础路由（开发中）"` // 是否是基础路由（开发中）
	Title       string `json:"title" gorm_study:"comment:菜单名"`                // 菜单名
	Icon        string `json:"icon" gorm_study:"comment:菜单图标"`                // 菜单图标
	CloseTab    bool   `json:"closeTab" gorm_study:"comment:自动关闭tab"`         // 自动关闭tab
}

type BaseMenuParameter struct {
	global.MODEL
	SysBaseMenuID uint
	Type          string `json:"type" gorm_study:"comment:地址栏携带参数为params还是query"` // 地址栏携带参数为params还是query
	Key           string `json:"key" gorm_study:"comment:地址栏携带参数的key"`            // 地址栏携带参数的key
	Value         string `json:"value" gorm_study:"comment:地址栏携带参数的值"`            // 地址栏携带参数的值
}
