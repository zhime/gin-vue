package models

type Tag struct {
	ID      uint
	Name    string
	Article []Article `gorm:"many2many:article_tags"`
}
