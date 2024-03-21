package config

type Zap struct {
	Level         string `json:"level"`
	Format        string `json:"format"`
	Prefix        string `json:"prefix"`
	Director      string `json:"director"`
	ShowLine      bool   `json:"show_line"`
	EncodeLevel   string `json:"encode_level"`
	StacktraceKey string `json:"stacktrace_key"`
	LogInConsole  bool   `json:"log_in_console"`
}
