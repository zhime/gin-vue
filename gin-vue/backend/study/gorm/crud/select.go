package crud

import (
	"encoding/json"
	"fmt"
	"github.com/zhime/gin-vue/gin-vue/study/gorm/db"
	"github.com/zhime/gin-vue/gin-vue/study/gorm/models"
	"gorm.io/gorm"
)

func Select() {
	var student models.Student

	// 查询日志
	loggerDB := db.DB.Session(&gorm.Session{
		Logger: db.MysqlLogger,
	})
	// 查询一条记录
	//loggerDB.Take(&student)
	//fmt.Println(student)
	//
	//student = models.Student{}
	//loggerDB.First(&student)
	//fmt.Println(student)
	//
	//student = models.Student{}
	//loggerDB.Last(&student)
	//fmt.Println(student)

	// 根据主键查询
	//loggerDB.Take(&student, 2)
	//fmt.Println(student)

	// 根据字段查询
	loggerDB.Take(&student, "name = ?", "golang8")
	fmt.Println(student)

	// 查询多条记录
	var students []models.Student
	count := loggerDB.Find(&students).RowsAffected
	fmt.Println(count)
	//for _, value := range students {
	//	fmt.Println(value)
	//}
	data, _ := json.Marshal(&students)
	fmt.Println(string(data))

	// 根据主键列表查询
	loggerDB.Find(&students, []int{2, 3, 5})
	fmt.Println(students)

	// 根据其他条件查询
	loggerDB.Find(&students, "name in ?", []string{"golang1", "golang2"})
	fmt.Println(students)
}
