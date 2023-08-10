package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/zhime/gin-vue/gin-vue/global"
	"os"
	"path"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// 自定义日期格式
	timeStamp := entry.Time.Format("2006-01-02 15:04:05")
	log := global.Config.Logger
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		// 自定义输出格式
		_, _ = fmt.Fprintf(b, "[%s] %s \x1b[%dm[%s]\x1b[0m %s %s %s\n", timeStamp, log.Prefix, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		_, _ = fmt.Fprintf(b, "[%s] %s \x1b[%dm[%s]\x1b[0m %s\n", timeStamp, log.Prefix, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() *logrus.Logger {
	mLog := logrus.New()                                        // 新建一个实例
	mLog.SetOutput(os.Stdout)                                   // 设置输出类型
	mLog.SetReportCaller(global.Config.Logger.ShowLine)         // 开启返回函数名和行号
	mLog.SetFormatter(&LogFormatter{})                          // 设置日志格式
	level, err := logrus.ParseLevel(global.Config.Logger.Level) // 设置日志等级
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level)
	//InitDefaultLogger()
	return mLog
}

//func InitDefaultLogger() {
//	logrus.SetOutput(os.Stdout)
//	logrus.SetReportCaller(global.Config.Logger.ShowLine)
//	logrus.SetFormatter(&LogFormatter{})
//	level, err := logrus.ParseLevel(global.Config.Logger.Level)
//	if err != nil {
//		level = logrus.InfoLevel
//	}
//	logrus.SetLevel(level)
//}
