package model

type LoginModel struct {
	Model
	UserID    uint      `json:"user_id"`
	UserModel UserModel `json:"user_model" gorm:"foreignKey:UserID"`
	IP        string    `json:"ip"`
	NickName  string    `json:"nick_name"`
	Token     string    `json:"token"`
	Device    string    `json:"device"`
	Addr      string    `json:"addr"`
}
