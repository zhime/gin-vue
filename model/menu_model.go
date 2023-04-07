package model

import (
	"github.com/zhime/gin-vue/model/ctype"
)

type MenuModel struct {
	Model
	MenuTitle    string        `json:"menu_title"`
	MenuTitleEn  string        `json:"menu_title_en"`
	Slogan       string        `json:"slogan"`
	Abstract     ctype.Array   `json:"abstract"`
	AbstractTime int           `json:"abstract_time"`
	Banners      []BannerModel `json:"banners" gorm:"many2many:menu_banner_models;JoinForeignKey:MenuID;JoinReferences:BannerID"`
	BannerTime   int           `json:"banner_time"`
	Sort         int           `json:"sort"`
}
