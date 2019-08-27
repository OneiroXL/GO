package accountInfoBLL

import (
	"chat/model/base/entity"
	"chat/server/dao/accountInfo"
	"chat/model/user"
)

var (
	userDAO accountInfo.UserDAO = accountInfo.UserDAO{}
)

type UserBLL struct {
	
}

/*
添加用户
*/
func (this *UserBLL) AddUser(userRegisterModel user.UserRegisterModel) error {
	user := entity.User{}

	user.UserCode = userRegisterModel.UserCode
	user.Passcode = userRegisterModel.Passcode
	user.UserName = userRegisterModel.UserName

	return userDAO.AddUser(user)
}


/*
获取用户信息
*/
func (this *UserBLL) GetUser(userCode string) (entity.User,error) {
	user,err := userDAO.GetUser(userCode)
	return user,err
}

/*
是否存在用户
*/
func (this *UserBLL) IsExistUser(userCode string) bool {
	user,_ := this.GetUser(userCode)
	if(user.ID != 0){
		return true
	} else {
		return false
	}
}


/*
登录验证
*/
func (this *UserBLL) LoginVerification(userLoginModel user.UserLoginModel) (bool,string,error) {
	//获取用户信息
	user,err := this.GetUser(userLoginModel.UserCode)
	if(user.ID == 0){
		return false,"该用户不存在",err
	}
	if(err != nil){
		return false,"错误",err
	}
	if(user.Passcode != userLoginModel.Passcode){
		return false,"密码错误",err
	}
	return true,"登录成功",nil
}
