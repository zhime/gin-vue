package main

import (
	"fmt"
	"github.com/zhime/gin-vue/gorm_study/crud"
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
	crud.SelectPro()
}
