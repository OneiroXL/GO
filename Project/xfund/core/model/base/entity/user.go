package entity

import (

)

type User struct {
	BaseEntity
	UserCode		string `gorm:"column:UserCode;type:varchar(50);not null"`
	UserName       string `gorm:"column:UserName;type:varchar(50);not null"`
	Password  	   string `gorm:"column:Password;type:varchar(50);not null"`
	Nickname       string `gorm:"column:Nickname;type:varchar(50);not null"`
	Status         string `gorm:"column:Status;type:int;not null"`
	Avatar         string `gorm:"size:1000;column:Avatar"`
	MobileNumber         string `gorm:"column:MobileNumber;type:varchar(20);not null"`
}

func (this *User) TableName() string {
	return "User"
}