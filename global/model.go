package global

import (
	"gorm.io/gorm"
	"time"
)

type MODEL struct {
	ID        uint           `gorm_study:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm_study:"index" json:"-"` // 删除时间
}
