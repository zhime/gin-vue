package model

type MenuBannerModel struct {
	MenuID      uint        `json:"menu_id"`
	MenuModel   MenuModel   `json:"menu_model" gorm:"foreignKey:MenuID"`
	BannerID    uint        `json:"banner_id"`
	BannerModel BannerModel `json:"banner_model" gorm:"foreignKey:BannerID"`
	Sort        int         `json:"sort"`
}
