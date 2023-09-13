package model

import (
	"github.com/google/uuid"
	"github.com/zhime/gin-vue/model/ctype"
)

type UserModel struct {
	Model
	UUID         uuid.UUID        `json:"uuid"`
	NickName     string           `json:"nick_name" gorm:"comment:用户昵称"`
	UserName     string           `json:"user_name" gorm:"index;comment:用户登录名"`
	Password     string           `json:"-"  gorm:"comment:用户登录密码"`
	Avatar       string           `json:"avatar" gorm:"comment:头像id"`
	Email        string           `json:"email" gorm:"comment:用户邮箱"`
	Tel          string           `json:"tel" gorm:"comment:手机号"`
	Addr         string           `json:"addr" gorm:"comment:地址"`
	Token        string           `json:"token" gorm:"comment：token"`
	IP           string           `json:"ip" gorm:"comment:ip"`
	Role         ctype.Role       `json:"role" gorm:"comment:角色"`
	SignStatus   ctype.SignStatus `json:"sign_status" gorm:"comment:注册来源"`
	ArticleModel []ArticleModel   `json:"article_models" gorm:"foreignKey:UserID;comment:发布的文章列表"`
	CollectModel []ArticleModel   `json:"collect_models" gorm:"many2many:user_collect_model;JoinForeignKey:UserID;JoinReferences:ArticleID;comment:收藏的文章列表"`
}
