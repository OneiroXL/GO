package dao


import (
	"xfund/core/model/base"
	"xfund/core/model/base/entity"
)

type UserDAO struct{
	
}

/*
添加用户
*/
func (this *UserDAO) AddUser(user entity.User) (bool,error) {
	res := base.DB.Create(&user)
	if (res.Error != nil){
		return false,res.Error
	}

	return true,nil;
}