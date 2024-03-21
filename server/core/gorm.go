package core

import (
	"fmt"
	"github.com/zhime/gin-vue/server/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGorm() {
	db, err := gorm.Open(mysql.Open(global.Config.Mysql.Dsn()), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("can't init mysql: %v \n", err))
	}
	global.Db = db
}
