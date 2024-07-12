package config

type System struct {
	Host string `json:"host" yaml:"host" mapstructure:"host"`
	Port int    `json:"port" yaml:"port" mapstructure:"port"`
	Env  string `json:"env" yaml:"env" mapstructure:"env"`
}
