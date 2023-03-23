package crud

import (
	"fmt"
	"github.com/zhime/gin-vue/gorm_study/db"
	"github.com/zhime/gin-vue/gorm_study/models"
)

func One2OneInsert() {
	_ = db.DB.Debug().AutoMigrate(&models.User{}, &models.UserInfo{})

	// 新增
	//db.DB.Debug().Create(&models.User{
	//	Name:   "张三",
	//	Age:    18,
	//	Gender: "男",
	//	UserInfo: models.UserInfo{
	//		Addr: "杭州",
	//		Like: "游泳",
	//	},
	//})

	// 查询
	var user models.User
	db.DB.Debug().Preload("UserInfo").Take(&user)
	fmt.Println(user)
}
