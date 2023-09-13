package crud

import (
	"github.com/zhime/gin-vue/study/gorm/db"
	"github.com/zhime/gin-vue/study/gorm/models"
	"strconv"
)

func Insert() {
	email := "admin@qq.com"

	// 单条插入
	db.DB.Create(&models.Student{
		Name:   "golang",
		Age:    20,
		Gender: true,
		Email:  &email,
	})

	// 批量插入
	var students []models.Student
	var gender bool
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			gender = true
		} else {
			gender = false
		}
		students = append(students, models.Student{
			Name:   "golang" + strconv.Itoa(i+1),
			Age:    20 + i + 1,
			Gender: gender,
			Email:  &email,
		})
	}
	db.DB.Create(&students)
}
