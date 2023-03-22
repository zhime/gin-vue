package crud

import (
	"github.com/zhime/gin-vue/gorm_study/db"
	"github.com/zhime/gin-vue/gorm_study/models"
	"gorm.io/gorm"
)

func SelectPro() {
	var studentList []models.Student
	// 查询日志
	loggerDB := db.DB.Session(&gorm.Session{
		Logger: db.MysqlLogger,
	})
	//loggerDB.Where("name", "武则天").Find(&studentList)
	//fmt.Println(studentList)
	//
	//loggerDB.Not("name", "武则天").Find(&studentList)
	//fmt.Println(studentList)

	//loggerDB.Where("name in ?", []string{"武则天", "李元芳"}).Find(&studentList)
	//fmt.Println(studentList)

	//loggerDB.Where("name like ?", "%天").Find(&studentList)
	//fmt.Println(studentList)

	//loggerDB.Select([]string{"name", "age"}).Find(&studentList)
	//fmt.Println(studentList)

	// 排序。SELECT * FROM `students` ORDER BY age desc
	//loggerDB.Order("age desc").Find(&studentList)
	//fmt.Println(studentList)

	// 分页查询。SELECT * FROM `students` LIMIT 2
	loggerDB.Limit(2).Offset(0).Find(&studentList)
	//fmt.Println(studentList)
	//limit := 2
	//page := 2
	//loggerDB.Limit(limit).Offset(limit * (page - 1)).Find(&studentList)
	//fmt.Println(studentList)

	// 分组查询。SELECT count(id) FROM `students` GROUP BY `gender`
	//var countList []int
	//loggerDB.Model(&models.Student{}).Select("count(id) as count").Group("gender").Scan(&countList)
	//fmt.Println(countList)

	//type Group struct {
	//	Count  int
	//	Gender string
	//}
	//
	//var groupList []Group
	//loggerDB.Model(&models.Student{}).Select("count(id) as count", "gender").Group("gender").Scan(&groupList)
	//fmt.Println(groupList)

	//var count []map[string]interface{}
	//loggerDB.Model(&models.Student{}).Select("count(id) as count", "gender").Group("gender").Scan(&count)
	//fmt.Println(count)
	//for _, value := range count {
	//	fmt.Printf("count:%d,gender:%d\n", value["count"], value["gender"])
	//}
}
