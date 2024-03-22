package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type config struct {
	Host  string `json:"host" mapstructure:"host"`
	Port  string `json:"port" mapstructure:"port"`
	Mysql Mysql  `json:"mysql" mapstructure:"mysql"`
}

type Mysql struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	Database string `json:"database" mapstructure:"database"`
}

func main() {
	// 设置默认值
	//viper.SetDefault("fileDir", "./middleware/viper/settings.yaml")
	// 读取配置文件
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./middleware/viper") // 配置文件查找路径
	viper.AddConfigPath(".")
	//viper.SetConfigFile("./middleware/viper/settings.yaml") // 指定配置文件
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Fatal error config file: %v\n", err)
	}

	var c config
	if err := viper.Unmarshal(&c); err != nil {
		fmt.Printf("viper.Unmarshal error: %v\n", err)
	}
	//if err := viper.UnmarshalKey("mysql", &c.Mysql); err != nil {
	//	fmt.Printf("viper.Unmarshal error: %v\n", err)
	//}

	fmt.Println(c)
}
