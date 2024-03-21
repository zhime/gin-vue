package core

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zhime/gin-vue/server/global"
)

func InitViper() {
	v := viper.New()
	v.SetConfigFile("settings.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := v.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
