package core

import (
	"github.com/spf13/viper"
	"github.com/zhime/gin-vue/study/gin/global"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile("study/gin/settings.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic("读取配置文件出错")
	}

	//var conf config.Config
	err = v.Unmarshal(&global.Config)
	if err != nil {
		panic("unmarshal error.")
	}
	//fmt.Println(global.Config)
}
