package flag

import (
	"github.com/zhime/gin-vue/global"
	"github.com/zhime/gin-vue/model"
)

func MakeMigrations() {
	_ = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&model.User{},
		&model.Authority{},
		//&model.BaseMenu{},
		//&model.BaseMenuBtn{},
		//&model.BaseMenuParameter{},
		//&model.Meta{},
	)
}
