package entity

import (

)

type File struct {
	BaseEntity
	FileID		string `gorm:"column:FileID;type:varchar(50);not null"`
	FileName       string `gorm:"column:FileName;type:varchar(50);not null"`
	FileType       int `gorm:"column:FileType;type:int;not null"`
	AbsolutePath  	   string `gorm:"column:AbsolutePath;type:varchar(150)"`
	VirtualAddress       string `gorm:"column:VirtualAddress;type:varchar(150);not null"`
}

func (this *File) TableName() string {
	return "File"
}