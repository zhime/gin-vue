package models

type User struct {
	ID       uint `gorm:"size:4"`
	Name     string
	Age      uint
	Gender   string
	UserInfo UserInfo
	Article  []Article `gorm:"foreignKey:Uid"`
}

type Article struct {
	ID    uint
	Title string
	Uid   uint  `gorm:"size:4"`
	User  User  `gorm:"foreignKey:Uid"`
	Tag   []Tag `gorm:"many2many:article_tags"`
}
