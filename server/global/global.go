package global

import (
	"github.com/zhime/gin-vue/server/config"
	"go.uber.org/zap"
)

var (
	Config config.Config
	Zap    *zap.Logger
)
