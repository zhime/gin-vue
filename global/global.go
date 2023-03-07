package global

import (
	"github.com/zhime/gin-vue/config"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
