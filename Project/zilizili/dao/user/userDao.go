package dao

import (
	"zilizili/model"
	"zilizili/model/base/entity"
	"zilizili/model/user/response"
)


type UserDao struct {

}

/*
添加用户
*/
func (this *UserDao) AddUser(user entity.User) (bool,error) {
	res := model.DB.Create(&user)
	if (res.Error != nil){
		return false,res.Error
	}

	return true,nil;
}

/*
更新用户信息
*/
func (this *UserDao) UpdateUser(user entity.User) (bool,error) {
	res := model.DB.Model(&user).Update(user)
	if(res.Error != nil){
		return false,res.Error
	}
	return true,nil
}

/*
获取用户
*/
func (this *UserDao) GetUser(userID int) (response.UserResponseModel,error) {
	userResponseModel := response.UserResponseModel{}
	res := model.DB.First(&userResponseModel,userID)
	return userResponseModel,res.Error;
}

/*
获取用户列表
*/
func (this *UserDao) GetUserList() ([]response.UserResponseModel,int,error) {
	userResponseList :=  []response.UserResponseModel{}
	res := model.DB.Find(&userResponseList)
	var count int
	model.DB.Model(&userResponseList).Count(&count)
	return  userResponseList,count,res.Error
}
/*
删除用户
*/
func (this *UserDao) DeleteUser(userId int)(bool,error){
	res := model.DB.Where("ID = ?", userId).Delete(&entity.User{})
	if(res.Error != nil){
		return false,res.Error
	} else {
		return true,nil
	}
}
