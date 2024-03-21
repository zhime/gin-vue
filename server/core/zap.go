package core

import (
	"fmt"
	"github.com/zhime/gin-vue/server/global"
	"go.uber.org/zap"
)

func InitZap() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Errorf("can't initialize zap logger: %v \n", err))
	}
	global.Zap = logger
}
