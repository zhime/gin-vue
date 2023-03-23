package main

import (
	"fmt"
	"github.com/zhime/gin-vue/gorm_study/db"
	"github.com/zhime/gin-vue/gorm_study/models"
)

func main() {
	// 创建表
	//err := db.DB.Debug().AutoMigrate(&models.Student{})
	err := db.DB.AutoMigrate(&models.Student{})
	if err != nil {
		fmt.Println(err)
	}

	// 添加记录
	//crud.Insert()

	// 查询
	//crud.Select()

	// 更新
	//crud.Update()

	//crud.InsertPro()

	// 高级查询
	//crud.SelectPro()

	// 一对多
	//_ = db.DB.Debug().AutoMigrate(&models.User{}, &models.Article{})
	//db.DB.Debug().Create(&models.User{
	//	Name: "zz",
	//	Article: []models.Article{
	//		{
	//			Title: "Hello",
	//		},
	//	},
	//})

	//db.DB.Debug().Create(&models.Article{
	//	Title: "golang",
	//	Uid:   1,
	//})

	//db.DB.Debug().Create(&models.Article{
	//	Title: "golang",
	//	User: models.User{
	//		Name: "张三",
	//	},
	//})
}
