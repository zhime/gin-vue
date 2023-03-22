package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

var DB *gorm.DB

func init() {
	host := "127.0.0.1"
	username := "root"
	password := "123456"
	port := 3306
	dbName := "gorm"
	config := "charset=utf8mb4&parseTime=True&loc=Local"

	//db := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := username + ":" + password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + dbName + "?" + config

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix:   "gorm_", // 表名前缀
		//	SingularTable: true,    // 单数表名
		//},
	})
	if err != nil {
		panic("数据库连接失败，" + err.Error())
	}
	DB = db
}
