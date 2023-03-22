package models

type Student struct {
	ID   uint   `gorm:"column:id;size:20"`
	Name string `gorm:"column:name;size:16;comment:姓名"`
	Age  int    `gorm:"column:age;size:3;comment:年龄"`
}
