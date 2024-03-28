package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strconv"
	"time"
)

var (
	Db          *gorm.DB
	MysqlLogger logger.Interface
)

func init() {
	host := "127.0.0.1"
	username := "root"
	password := "123456"
	port := 3306
	dbName := "gorm"
	config := "charset=utf8mb4&parseTime=True&loc=Local"

	//mysqlLogger := logger.Default.LogMode(logger.Info) // 设置日志等级
	MysqlLogger = logger.Default.LogMode(logger.Info)

	//db := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := username + ":" + password + "@tcp(" + host + ":" + strconv.Itoa(port) + ")/" + dbName + "?" + config

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction: false, // 跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "gorm_", // 表名前缀
			SingularTable: true,    // 表名单数
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束
	})

	// 连接池
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		fmt.Printf("Connect mysql err: %v", err)
	}

	//db, err := gorm.Open(mysql.New(mysql.Config{
	//	DSN:                       dsn,   // DSN data source name
	//	DefaultStringSize:         256,   // string 类型字段的默认长度
	//	DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	//	DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	//	DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	//	SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	//}), &gorm.Config{})

	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	//NamingStrategy: schema.NamingStrategy{
	//	//	TablePrefix:   "gorm_", // 表名前缀
	//	//	SingularTable: true,    // 单数表名
	//	//},
	//	//Logger: mysqlLogger,  // 日志
	//})
	//if err != nil {
	//	panic("数据库连接失败，" + err.Error())
	//}

	//DB = db.Session(&gorm.Session{
	//	Logger: mysqlLogger,
	//})

	Db = db
}
