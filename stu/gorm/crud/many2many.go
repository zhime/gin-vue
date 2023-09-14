package crud

import (
	"fmt"
	"github.com/zhime/gin-vue/stu/gorm/db"
	models2 "github.com/zhime/gin-vue/stu/gorm/models"
)

func Many2ManyInsert() {
	_ = db.DB.Debug().AutoMigrate(&models2.Tag{}, &models2.Article{})

	db.DB.Debug().Create(&models2.Article{
		Title: "django进阶",
		User: models2.User{
			Name:   "张三",
			Age:    18,
			Gender: "男",
		},
		Tag: []models2.Tag{
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
	var article models2.Article
	db.DB.Debug().Preload("Tag").Take(&article)
	fmt.Println(article)
}

func Many2ManyDelete() {
	var article models2.Article
	db.DB.Debug().Preload("Tag").Take(&article)
	// 级联删除
	_ = db.DB.Debug().Model(&article).Association("Tag").Delete(article.Tag)
}
