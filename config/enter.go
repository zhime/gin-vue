package config

type Config struct {
	System System `yaml:"system"`
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
}
