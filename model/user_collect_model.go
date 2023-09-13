package model

import "time"

type UserCollectModel struct {
	UserID       uint         `gorm:"primarykey"`
	UserModel    UserModel    `gorm:"foreignKey:UserID;"`
	ArticleID    uint         `gorm:"primarykey"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID;"`
	CreateAt     time.Time
}
