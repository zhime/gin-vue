package model

import "time"

type Authority struct {
	CreatedAt       time.Time   // 创建时间
	UpdatedAt       time.Time   // 更新时间
	DeletedAt       *time.Time  `sql:"index"`
	AuthorityId     uint        `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	AuthorityName   string      `json:"authorityName" gorm:"comment:角色名"`                                   // 角色名
	ParentId        *uint       `json:"parentId" gorm:"comment:父角色ID"`                                      // 父角色ID
	DataAuthorityId []Authority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id;"`
	Children        []Authority `json:"children" gorm:"-"`
	BaseMenus       []BaseMenu  `json:"menus" gorm:"many2many:sys_authority_menus;"`
	Users           []User      `json:"-" gorm:"many2many:sys_user_authority;"`
	DefaultRouter   string      `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"` // 默认菜单(默认dashboard)
}
