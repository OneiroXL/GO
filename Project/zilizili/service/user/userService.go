package service

import (
	"zilizili/dao/user"
	"zilizili/model/base/entity"
	"zilizili/model/user/param"
	"zilizili/model/user/response"
	"zilizili/common/constant/userConstant"
	"zilizili/common/tools"
	
)

var (
	userDao dao.UserDao
)

type UserService struct{

}


/*
添加用户
*/
func (this *UserService) AddUser(userParamModel param.UserParamModel) (bool,error) {
	user := entity.User{}
	
	user.UserName = userParamModel.UserName;
	cryptoPasswor,_ := this.GetCryptoPassword(userParamModel.Password);
	user.Password = cryptoPasswor;
	user.Nickname = userParamModel.Nickname;
	user.Status = userParamModel.Status;
	user.Avatar = userParamModel.Avatar;

	return userDao.AddUser(user);
}

/*
更新用户信息
*/
func (this *UserService) UpdateUser(userParamModel param.UserParamModel) (bool,error) {
	user := entity.User{}
	
	user.ID = userParamModel.ID;

	user.UserName = userParamModel.UserName;
	user.Nickname = userParamModel.Nickname;
	user.Status = userParamModel.Status;
	user.Avatar = userParamModel.Avatar;

	return userDao.AddUser(user);
}

/*
获取用户
*/
func (this *UserService) GetUser(userID int) (response.UserResponseModel,error) {
	return userDao.GetUser(userID);
}

/*
获取用户列表
*/
func (this *UserService) GetUserList() ([]response.UserResponseModel,int,error) {
	return userDao.GetUserList();
}
/*
删除用户
*/
func (this *UserService) DeleteUser(userID int)(bool,error){
	return userDao.DeleteUser(userID);
}


/*
加密密码
*/
func (user *UserService) GetCryptoPassword(password string) (string,error) {
	return cryptoTool.CryptoPassword(password,userConstant.PASSWORD_COST);
}

/*
验证密码
*/
func (user *UserService) CheckPassword(password string,cryptoPassword string) bool {
	return cryptoTool.CheckPassword(password,cryptoPassword)
}
