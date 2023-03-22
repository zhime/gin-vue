package model

import (
	"github.com/satori/go.uuid"
	"github.com/zhime/gin-vue/global"
)

type User struct {
	global.MODEL
	UUID        uuid.UUID `json:"uuid" gorm_study:"index;comment:用户UUID"`                                                     // 用户UUID
	NickName    string    `json:"nickName" gorm_study:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	UserName    string    `json:"userName" gorm_study:"index;comment:用户登录名"`                                                  // 用户登录名
	Password    string    `json:"-"  gorm_study:"comment:用户登录密码"`                                                             // 用户登录密码
	SideMode    string    `json:"sideMode" gorm_study:"default:dark;comment:用户侧边主题"`                                          // 用户侧边主题
	HeaderImg   string    `json:"headerImg" gorm_study:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	BaseColor   string    `json:"baseColor" gorm_study:"default:#fff;comment:基础颜色"`                                           // 基础颜色
	ActiveColor string    `json:"activeColor" gorm_study:"default:#1890ff;comment:活跃颜色"`                                      // 活跃颜色
	AuthorityId uint      `json:"authorityId" gorm_study:"default:888;comment:用户角色ID"`                                        // 用户角色ID
	Authority   Authority `json:"authority" gorm_study:"foreignKey:AuthorityId;comment:用户角色"`
	//Authority   Authority `json:"authority" gorm_study:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	//Authorities []Authority `json:"authorities" gorm_study:"many2many:user_authority;"`
	//Phone       string      `json:"phone"  gorm_study:"comment:用户手机号"`                     // 用户手机号
	//Email       string      `json:"email"  gorm_study:"comment:用户邮箱"`                      // 用户邮箱
	//Enable      int         `json:"enable" gorm_study:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
}
