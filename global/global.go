package global

import (
	"github.com/sirupsen/logrus"
	"github.com/zhime/gin-vue/gin-vue/config"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
