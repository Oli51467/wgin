package model

import (
	"gorm.io/gorm"
)

// ID 自增ID主键
type ID struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

// Timestamps 创建、更新时间
type Timestamps struct {
	CreateTime *LocalTime `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime *LocalTime `json:"update_time" gorm:"autoCreateTime"`
}

// SoftDeletes 软删除
type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
