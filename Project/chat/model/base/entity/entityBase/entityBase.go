package entityBase


import (
	"time"
)

type EntityBase struct  {
	ID 	uint `gorm:"primary_key;column:ID;not null"`
	CreatedAt time.Time `gorm:"column:CreatedAt;not null"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt"`
	DeletedAt time.Time	`gorm:"column:DeletedAt;DEFAULT:null"`
}

