package crud

import (
	"github.com/zhime/gin-vue/stu/gorm/db"
	"github.com/zhime/gin-vue/stu/gorm/models"
	"gorm.io/gorm"
)

func Update() {
	var student models.Student
	loggerDB := db.DB.Session(&gorm.Session{
		Logger: db.MysqlLogger,
	})

	loggerDB.Take(&student, 1)
	student.Name = "golang2"

	// 更新单个字段。UPDATE `students` SET `name`='go10' WHERE `id` = 11
	//loggerDB.Select("name").Save(&student)

	// 更新整行。UPDATE `students` SET `name`='golang10',`age`=30,`gender`=false,`email`='admin@qq.com' WHERE `id` = 11
	//loggerDB.Save(&student)

	// 单个字段批量更新。UPDATE `students` SET `gender`=false WHERE `students`.`id` IN (10,11) AND `id` IN (10,11)
	var students []models.Student
	//loggerDB.Find(&students, []int{10, 11}).Update("gender", false)

	// 多字段批量更新。UPDATE `students` SET `age`=20,`gender`=true WHERE `students`.`id` IN (10,11) AND `id` IN (10,11)
	//loggerDB.Find(&students, []int{10, 11}).Updates(models.Student{Age: 20, Gender: true})

	// UPDATE `students` SET `age`=22,`gender`=false WHERE `students`.`id` IN (10,11) AND `id` IN (10,11)
	//loggerDB.Find(&students, []int{10, 11}).Updates(map[string]interface{}{
	//	"age":    22,
	//	"gender": false,
	//})

	// 删除。DELETE FROM `students` WHERE `students`.`id` = 11
	//loggerDB.Delete(&student, 11)

	// 批量删除。DELETE FROM `students` WHERE `students`.`id` IN (8,9)
	loggerDB.Delete(&students, []int{8, 9})
}
