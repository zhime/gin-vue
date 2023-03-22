package main

import (
	"fmt"
	"github.com/zhime/gin-vue/gorm_study/db"
	"github.com/zhime/gin-vue/gorm_study/models"
)

func main() {
	//err := db.DB.Debug().AutoMigrate(&models.Student{})
	err := db.DB.AutoMigrate(&models.Student{})
	if err != nil {
		fmt.Println(err)
	}
}
