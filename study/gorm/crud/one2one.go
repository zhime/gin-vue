package crud

import (
	"fmt"
	"github.com/zhime/gin-vue/study/gorm/db"
	models2 "github.com/zhime/gin-vue/study/gorm/models"
)

func One2OneInsert() {
	_ = db.DB.Debug().AutoMigrate(&models2.User{}, &models2.UserInfo{})

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
	var user models2.User
	db.DB.Debug().Preload("UserInfo").Take(&user)
	fmt.Println(user)
}
