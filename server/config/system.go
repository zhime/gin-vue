package config

type System struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Env  string `json:"env"`
}
