package crud

import (
	"fmt"
	"github.com/zhime/gin-vue/stu/gorm/db"
	"github.com/zhime/gin-vue/stu/gorm/models"
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

func One2ManySelect() {
	var user models.User
	db.DB.Debug().Preload("Article").Take(&user)
	fmt.Println(user)

	var article models.Article
	db.DB.Debug().Preload("User").Take(&article)
	fmt.Println(article)
}

func One2ManyDelete() {
	// 外键置为null
	//var user models.User
	//db.DB.Debug().Preload("Article").Take(&user, 2)
	//_ = db.DB.Debug().Model(&user).Association("Article").Delete(&user.Article)
	//db.DB.Debug().Delete(&user)

	// 级联删除
	var user models.User
	db.DB.Debug().Take(&user, 1)
	db.DB.Debug().Select("Article").Delete(&user)
}
