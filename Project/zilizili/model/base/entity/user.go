package entity

import (

)

type User struct {
	BaseEntity
	UserName       string `gorm:"column:UserName;type:varchar(50);not null"`
	Password  	   string `gorm:"column:Password;type:varchar(50);not null"`
	Nickname       string `gorm:"column:Nickname;type:varchar(50);not null"`
	Status         string `gorm:"column:Status;type:int;not null"`
	Avatar         string `gorm:"size:1000;column:Avatar"`
}

//为Video绑定表名
func (video *User) TableName() string {
	return "User"
}