package core

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zhime/gin-vue/global"
)

func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigName("settings")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	v.WatchConfig()

	if err := v.Unmarshal(&global.G_CONFIG); err != nil {
		panic(err)
	}

	return v
}
