package config

type Config struct {
	System System `json:"system"`
	Mysql  Mysql  `json:"mysql"`
	Redis  Redis  `json:"redis"`
	Zap    Zap    `json:"zap"`
}
