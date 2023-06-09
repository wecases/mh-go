package models

import (
	"time"
)

// 需要迁移的表
var Tables []interface{} = []interface{}{
	&User{},
}

// 需要重建的表
var AgainTables []interface{} = []interface{}{
	// &User{},
}

type BaseModel struct {
	ID uint `gorm:"primary_key; comment:'ID'" json:"id"`
}

type Model struct {
	BaseModel
	CreatedAt time.Time `gorm:"comment:'创建时间'" json:"create_at"`
	UpdatedAt time.Time `gorm:"comment:'更新时间'" json:"update_at"`
}

type SoftDelete struct {
	Model
	DeletedAt *time.Time `gorm:"comment:'删除时间'" json:"delete_at"`
}
