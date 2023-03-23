package crud

import (
	"github.com/zhime/gin-vue/gorm_study/db"
	"github.com/zhime/gin-vue/gorm_study/models"
)

// One2ManyInsert 一对多
func One2ManyInsert() {
	_ = db.DB.Debug().AutoMigrate(&models.User{}, &models.Article{})
	db.DB.Debug().Create(&models.User{
		Name: "张三",
		Article: []models.Article{
			{
				Title: "golang",
			},
		},
	})

	db.DB.Debug().Create(&models.Article{
		Title: "python",
		Uid:   1,
	})

	db.DB.Debug().Create(&models.Article{
		Title: "java",
		User: models.User{
			Name: "李四",
		},
	})

	//db.DB.Debug().Create(&models.Article{
	//	Title: "php",
	//	User:  models.User{},
	//})
	//
	//var user models.User
	//db.DB.Take(&user, 2)
	//var article models.Article
	//db.DB.Take(&article, 4)
	//_ = db.DB.Model(&user).Association("Article").Append(&article)
	//_ = db.DB.Model(&article).Association("User").Append(&user)

}
