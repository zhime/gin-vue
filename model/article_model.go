package model

import "github.com/zhime/gin-vue/model/ctype"

type ArticleModel struct {
	Model
	Title        string         `json:"title" gorm:"comment:文章标题"`
	Abstract     string         `json:"abstract" gorm:"comment:文章简介"`
	Content      string         `json:"content" gorm:"comment:文章内容"`
	LookCount    int            `json:"look_count" gorm:"comment:浏览量"`
	CommentCount int            `json:"comment_count" gorm:"comment:评论量"`
	DiggCount    int            `json:"digg_count" gorm:"comment:点赞量"`
	CollectCount int            `json:"collect_count" gorm:"comment:收藏量"`
	TagModel     []TagModel     `json:"tag_models" gorm:"many2many:article_tag_models;comment:文章标签"`
	CommentModel []CommentModel `json:"comment_models" gorm:"foreignKey:ArticleID;comment:评论列表"`
	UserModel    UserModel      `json:"user_models" gorm:"foreignKey:UserID;comment:文章作者"`
	UserID       uint           `json:"user_id" gorm:"comment:用户id"`
	Category     string         `json:"category" gorm:"comment:文章分类"`
	Source       string         `json:"source" gorm:"comment:文章来源"`
	Link         string         `json:"link" gorm:"comment:文章链接"`
	Banner       BannerModel    `json:"banner" gorm:"foreignKey:BannerID;comment:文章封面"`
	BannerID     uint           `json:"banner_id"`
	NickName     string         `json:"nick_name" gorm:"comment:用户昵称"`
	BannerPath   string         `json:"banner_path" gorm:"comment:文章封面"`
	Tag          ctype.Array    `json:"tag" gorm:"comment:文章标签"`
}
