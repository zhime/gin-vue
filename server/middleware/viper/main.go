package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	// 设置默认值
	viper.SetDefault("fileDir", "./")
	// 读取配置文件
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	//viper.SetConfigFile("./settings.yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Fatal error config file: %v\n", err)
	}
}
