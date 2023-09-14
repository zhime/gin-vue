package crud

import (
	"github.com/zhime/gin-vue/stu/gorm/db"
	"github.com/zhime/gin-vue/stu/gorm/models"
)

func InsertPro() {
	var studentList []models.Student
	studentList = []models.Student{
		{Name: "武则天", Age: 18, Gender: false, Email: PtrSting("wzt@qq.com")},
		{Name: "狄仁杰", Age: 20, Gender: true, Email: PtrSting("drj@qq.com")},
		{Name: "孙尚香", Age: 21, Gender: false, Email: PtrSting("ssx@qq.com")},
		{Name: "王昭君", Age: 19, Gender: false, Email: PtrSting("wzj@qq.com")},
		{Name: "后裔", Age: 22, Gender: true, Email: PtrSting("hy@qq.com")},
		{Name: "李元芳", Age: 23, Gender: true, Email: PtrSting("lyf@qq.com")},
		{Name: "张良", Age: 25, Gender: true, Email: PtrSting("zl@qq.com")},
		{Name: "李白", Age: 21, Gender: true, Email: PtrSting("lb@qq.com")},
	}

	db.DB.Create(&studentList)
}

func PtrSting(email string) *string {
	return &email
}
