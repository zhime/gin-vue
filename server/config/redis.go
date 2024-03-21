package config

type Redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database int    `json:"database"`
	Password string `json:"password"`
}
