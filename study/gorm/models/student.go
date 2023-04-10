package models

type Student struct {
	ID     uint    `gorm:"column:id;size:20" json:"id"`
	Name   string  `gorm:"column:name;size:16;comment:姓名" json:"name"`
	Age    int     `gorm:"column:age;size:3;comment:年龄" json:"age"`
	Gender bool    `gorm:"column:gender;comment:性别" json:"gender"`
	Email  *string `gorm:"column:email;size:20;comment:邮箱" json:"email"`
}
