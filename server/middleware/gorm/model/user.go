package model

import (
	"database/sql"
	"fmt"
	"github.com/zhime/gin-vue/server/middleware/gorm/db"
	"time"
)

type Model struct {
	UUID uint `gorm:"primaryKey"`
	Time time.Time
}
type User struct {
	Model        Model          `gorm:"embedded;embeddedPrefix:model_"`
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields

}

func TestUserCreate() {
	_ = db.Db.AutoMigrate(&User{})

	db.Db.Create(&[]User{
		{Model: Model{
			UUID: uint(1),
			Time: time.Now(),
		}, Name: "test01"},
		{Model: Model{
			UUID: uint(2),
			Time: time.Now(),
		}, Name: "test02"},
		{Model: Model{
			UUID: uint(3),
			Time: time.Now(),
		}, Name: "test03"},
	})
}

func TestUserSelect() {
	var u User
	var u1 User
	var u2 User
	var u3 User
	db.Db.Model(&User{}).First(&u1)
	db.Db.Model(&User{}).Take(&u2)
	db.Db.Model(&User{}).Last(&u3)
	db.Db.Find(&u)
	fmt.Println(u1)
	fmt.Println(u2)
	fmt.Println(u3)
	fmt.Println(u)
}
