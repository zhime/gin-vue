package config

type Mysql struct {
	Host     string `json:"host" yaml:"host" mapstructure:"host"`
	Port     int    `json:"port" yaml:"port" mapstructure:"port"`
	Username string `json:"username" yaml:"username" mapstructure:"username"`
	Password string `json:"password" yaml:"password" mapstructure:"password"`
	Database string `json:"database" yaml:"database" mapstructure:"database"`
	Config   string `json:"config" yaml:"config" mapstructure:"config"`
	LogLevel string `json:"log_level" yaml:"log_level" mapstructure:"log_level"`
}
