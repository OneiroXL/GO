package entity

import (

)

type Video struct{
	BaseEntity
	Title string `gorm:"column:Title;type:varchar(100);not null"`
	Info  string `gorm:"column:Info;type:varchar(300);not null"`
	URL   string `gorm:"column:URL;type:varchar(300)"`
	Cover string `gorm:"column:Cover;type:varchar(100)"`
}

func (video Video) TableName() string {
	return "Video"
}