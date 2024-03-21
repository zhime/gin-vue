package global

import (
	"github.com/zhime/gin-vue/server/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config config.Config
	Zap    *zap.Logger
	Db     *gorm.DB
)
