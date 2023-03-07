package core

import (
	"fmt"
	"github.com/zhime/gin-vue/global"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

func InitConfig() {
	const ConfigFile = "settings.yaml"
	config := &global.Config
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, config)
	if err != nil {
		log.Fatalf("config init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load init success.")
}
