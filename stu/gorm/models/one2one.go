package models

//type Usr struct {
//	ID       int
//	Name     string
//	Age      int
//	Gender   string
//	UserInfo UserInfo
//}

type UserInfo struct {
	ID     uint
	UserID uint `gorm:"size:4"`
	Addr   string
	Like   string
}
