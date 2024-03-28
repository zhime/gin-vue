package table

import "github.com/zhime/gin-vue/server/middleware/gorm/db"

func InitTable() {
	type User struct {
		Name string
	}
	//_ = db.AutoMigrate(&User{})  // 自动迁移

	// 手动创建表,Migrator
	m := db.Db.Migrator()
	// 查询是否存在表
	if m.HasTable(&User{}) {
		// 存在删除表
		//_ = m.DropTable(&User{})

		// 重命名表
		_ = m.RenameTable(&User{}, "old_user")
	} else {
		// 不存在创建表
		_ = m.CreateTable(&User{})
	}
}
