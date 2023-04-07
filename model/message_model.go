package model

type MessageModel struct {
	Model
	SendUserID       uint      `json:"send_user_id" gorm:"primaryKey"`
	SendUserModel    UserModel `json:"send_user_model" gorm:"foreignKey:SendUserID"`
	SendUserNickName string    `json:"send_user_nick_name"`
	SendUserAvatar   string    `json:"send_user_avatar"`
	RevUserID        uint      `json:"rev_user_id" gorm:"primaryKey"`
	RevUserModel     UserModel `json:"rev_user_model" gorm:"foreignKey:RevUserID"`
	RevUserNickName  string    `json:"rev_user_nick_name"`
	RevUserAvatar    string    `json:"rev_user_avatar"`
	IsRead           bool      `json:"is_read"`
	Content          string    `json:"content"`
}
