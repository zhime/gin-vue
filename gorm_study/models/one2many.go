package models

type User struct {
	ID      uint `gorm:"size:4"`
	Name    string
	Article []Article `gorm:"foreignKey:Uid"`
}

type Article struct {
	ID    uint
	Title string
	Uid   uint `gorm:"size:4"`
	User  User `gorm:"foreignKey:Uid"`
}
