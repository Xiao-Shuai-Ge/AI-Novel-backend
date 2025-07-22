package model

import (
	"gorm.io/gorm"
	"time"
)

type TimeModel struct {
	CreatedTime int64 `gorm:"column:created_time;type:bigint"`
	UpdatedTime int64 `gorm:"column:updated_time;type:bigint"`
}

func (b *TimeModel) BeforeCreate(db *gorm.DB) error {
	// 生成雪花ID
	b.CreatedTime = time.Now().UnixMilli()
	b.UpdatedTime = time.Now().UnixMilli()

	return nil
}

func (b *TimeModel) BeforeUpdate(db *gorm.DB) error {
	b.UpdatedTime = time.Now().UnixMilli()

	return nil
}
