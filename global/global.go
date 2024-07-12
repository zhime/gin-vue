package global

import (
	"github.com/spf13/viper"
	"github.com/zhime/gin-vue/config"
)

var (
	G_CONFIG config.Config
	G_VIPER  *viper.Viper
)
