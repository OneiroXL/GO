package entity

import (
	"chat/model/base/entity/entityBase"
)

type User struct {
	entityBase.EntityBase
	UserCode string `gorm:"column:UserCode;not null"`
	Passcode string `gorm:"column:Passcode;not null"`
	UserName string `gorm:"column:UserName;not null"`
}