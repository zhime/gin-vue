package config

type Config struct {
	System System `json:"system" yaml:"system" mapstructure:"system"`
	Mysql  Mysql  `json:"mysql" yaml:"mysql" mapstructure:"mysql"`
	Redis  Redis  `json:"redis" yaml:"redis" mapstructure:"redis"`
	Zap    Zap    `json:"zap" yaml:"zap" mapstructure:"zap"`
}
