package model

type CommentModel struct {
	Model
	SubComment         []*CommentModel `json:"sub_comment" gorm:"foreignKey:ParentCommentID"`
	ParentCommentModel *CommentModel   `json:"parent_comment_model" gorm:"foreignKey:ParentCommentID"`
	ParentCommentID    *uint           `json:"parent_comment_id"`
	Content            string          `json:"content"`
	DiggCount          int             `json:"digg_count"`
	CommentCount       int             `json:"comment_count"`
	ArticleModel       ArticleModel    `json:"article_model" gorm:"foreignKey:ArticleID"`
	ArticleID          uint            `json:"article_id"`
	UserModel          UserModel       `json:"user" gorm:"foreignKey:UserID"`
	UserID             uint            `json:"user_id"`
}
