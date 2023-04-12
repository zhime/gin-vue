package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type ALiYunSms struct {
	AccessKeyId     string `mapstructure:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret"`
	RegionId        string `mapstructure:"region_id"`
	SignName        string `mapstructure:"sign_name"`
	TemplateCode    string `mapstructure:"template_code"`
}

type Config struct {
	ALiYunSms `mapstructure:"aliyun_sms"`
}

func main() {
	v := viper.New()
	v.SetConfigFile("study/gin/aliyunsms.yaml")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	var a Config
	if err = v.Unmarshal(&a); err != nil {
		fmt.Println(err)
	}
	fmt.Println(a.AccessKeyId)
}
