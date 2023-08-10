package model

type TagModel struct {
	Model
	Title        string
	ArticleModel []ArticleModel `json:"article_model" gorm:"many2many:article_tag_models"`
}
