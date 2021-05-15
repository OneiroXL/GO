package account

import (
	"HuoBiClient/model/account"
	HBSDKAccount "HuoBiClient/HBSDK/account"
)


var (
	//用户信息map[用户类型][用户信息]
	userInfpMap map[int64]*account.UserModel = make(map[int64]*account.UserModel,2)
)

//初始化用户信息
func UserInit(){
	//请求用户信息
	userInfoSilce := HBSDKAccount.GetUserInfo();
	if(userInfoSilce != nil){
		for _,userInfo := range userInfoSilce {

			userModel := new(account.UserModel)

			userModel.Id = userInfo.Id
			userModel.State = userInfo.State
			userModel.Subtype = userInfo.Subtype
			userModel.Type = userInfo.Type

			userInfpMap[userInfo.Id] = userModel
		}
	}
}


//获取所有用户
func GetAllUser() map[int64]*account.UserModel {
	return userInfpMap
}

//依据用户Id获取用户信息
func GetUserById(userId int64) *account.UserModel {
	return userInfpMap[userId]
}


//依据用户类型获取用户信息
func GetUserByType(userType string) *account.UserModel {
	for _,userInfo := range userInfpMap {
		if userInfo.Type == userType {
			return userInfo
		}
	}
	return nil
}