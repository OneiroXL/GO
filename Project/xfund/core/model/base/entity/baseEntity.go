package entity

import (
	"time"
)

// 基本模型的定义
type BaseEntity struct {
	ID 	uint `gorm:"primary_key;column:ID;not null"`
	CreatedTime time.Time `gorm:"column:CreatedTime;not null"`
	UpdatedTime time.Time `gorm:"column:UpdatedTime"`
	DeletedTime time.Time	`gorm:"column:DeletedTime;DEFAULT:null"`
}
  