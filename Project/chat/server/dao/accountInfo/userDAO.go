package accountInfo

import (
	"chat/model/base/entity"
	DBBase "chat/server/dao"
)


type UserDAO struct {

}

/*
添加用户
*/
func (this *UserDAO) AddUser(user entity.User) error {
	res := DBBase.DB.Create(&user)
	return res.Error
}

/*
获取用户信息
*/
func (this *UserDAO) GetUser(userCode string) (entity.User,error) {
	user := entity.User{}
	res := DBBase.DB.Where("UserCode = ?", userCode).First(&user)
	return user,res.Error;
}
