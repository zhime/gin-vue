package config

type Redis struct {
	Host     string `json:"host" yaml:"host" mapstructure:"host"`
	Port     int    `json:"port" yaml:"port" mapstructure:"port"`
	Database int    `json:"database" yaml:"database" mapstructure:"database"`
	Password string `json:"password" yaml:"password" mapstructure:"password"`
}
