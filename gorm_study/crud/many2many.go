package crud

import (
	"fmt"
	"github.com/zhime/gin-vue/gorm_study/db"
	"github.com/zhime/gin-vue/gorm_study/models"
)

func Many2ManyInsert() {
	_ = db.DB.Debug().AutoMigrate(&models.Tag{}, &models.Article{})

	db.DB.Debug().Create(&models.Article{
		Title: "django进阶",
		User: models.User{
			Name:   "张三",
			Age:    18,
			Gender: "男",
		},
		Tag: []models.Tag{
			{
				Name: "python",
			},
			{
				Name: "django",
			},
		},
	})
}

func Many2ManySelect() {
	var article models.Article
	db.DB.Debug().Preload("Tag").Take(&article)
	fmt.Println(article)
}

func Many2ManyDelete() {
	var article models.Article
	db.DB.Debug().Preload("Tag").Take(&article)
	// 级联删除
	_ = db.DB.Debug().Model(&article).Association("Tag").Delete(article.Tag)
}
