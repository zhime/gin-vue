package config

type Zap struct {
	Level         string `json:"level" yaml:"level" mapstructure:"level"`                            // 级别
	Prefix        string `json:"prefix" yaml:"prefix" mapstructure:"prefix"`                         // 日志前缀
	Format        string `json:"format" yaml:"format" mapstructure:"format"`                         // 输出
	Director      string `json:"director" yaml:"director" mapstructure:"director"`                   // 日志文件夹
	EncodeLevel   string `json:"encode-level" yaml:"encode-level" mapstructure:"encode-level"`       // 编码级
	StacktraceKey string `json:"stacktrace-key" yaml:"stacktrace-key" mapstructure:"stacktrace-key"` // 栈名
	MaxAge        int    `json:"max-age" yaml:"max-age" mapstructure:"max-age"`                      // 日志留存时间
	ShowLine      bool   `json:"show-line" yaml:"show-line" mapstructure:"show-line"`                // 显示行
	LogInConsole  bool   `json:"log-in-console" yaml:"log-in-console" mapstructure:"log-in-console"` // 输出控制台
}
