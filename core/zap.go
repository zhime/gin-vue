package core

import (
	"fmt"
	"github.com/zhime/gin-vue/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"

	"github.com/zhime/gin-vue/global"
)

func getEncoder() zapcore.Encoder {
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	//return zapcore.NewConsoleEncoder(encoderConfig)
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create(global.G_CONFIG.Zap.Director + "./server.log")
	return zapcore.AddSync(file)
}

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.G_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create directory %v \n", global.G_CONFIG.Zap.Director)
		_ = os.Mkdir(global.G_CONFIG.Zap.Director, os.ModePerm)
	}

	core := zapcore.NewCore(getEncoder(), getLogWriter(), zapcore.DebugLevel)
	logger = zap.New(core, zap.AddCaller())
	return logger
}
