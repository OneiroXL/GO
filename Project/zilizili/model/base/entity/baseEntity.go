package entity

import (
	"time"
)

// 基本模型的定义
type BaseEntity struct {
	ID 	uint `gorm:"primary_key;column:ID;not null"`
	CreatedAt time.Time `gorm:"column:CreatedAt;not null"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt"`
	DeletedAt time.Time	`gorm:"column:DeletedAt;DEFAULT:null"`
  }
  