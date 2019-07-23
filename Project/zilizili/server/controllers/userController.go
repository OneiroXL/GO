package controllers


import (
	"zilizili/server/controllers/base"
	"zilizili/model/base/result"
	"zilizili/service/user"
	"zilizili/model/user/param"
	"zilizili/common/enum/resultEnum"
)

type UserController struct {
	baseControllers.BaseController
}



/*
添加用户
*/
func (this *UserController) AddUser() {

	result := result.NewFailedResult()

	userParamModel :=  param.UserParamModel{}
	//检查惨参数并绑定
	if (!this.BindParseForm(&userParamModel)) {
		result.Message = "参数有误"
		this.ResponseJSON(result)
		return
	}
	//验证参数
	if isPass,msg := this.ValidParam(userParamModel);!isPass {
		result.Message = msg
		this.ResponseJSON(result)
		return
	}

	userService := service.UserService{}
	
	if isOk,err := userService.AddUser(userParamModel);!isOk{
		result.Message = err
		result.Status = resultEnum.Error
		this.ResponseJSON(result)
		return
	}
	result.Message = "成功"
	result.Status = 10000
	this.ResponseJSON(result)
}

/*
更新用户信息
*/
func (this *UserController) UpdateUser() {
	result := result.NewFailedResult()

	userParamModel :=  param.UserParamModel{}
	//检查惨参数并绑定
	if (!this.BindParseForm(&userParamModel)) {
		result.Message = "参数有误"
		this.ResponseJSON(result)
		return
	}
	//验证参数
	if isPass,msg := this.ValidParam(userParamModel);!isPass {
		result.Message = msg
		this.ResponseJSON(result)
		return
	}

	userService := service.UserService{}
	if isOk,err := userService.UpdateUser(userParamModel);!isOk{
		result.Message = err
		this.ResponseJSON(result)
		return
	}
	result.Message = "成功"
	result.Status = 10000
	this.ResponseJSON(result)
}

/*
获取用户
*/
func (this *UserController) GetUser(){
	result := result.NewFailedResultData()

	id,err := this.GetInt("ID")
	if(id == 0 || err != nil){
		result.Message = "传入的ID有误"
		this.ResponseJSON(result)
		return
	}
	userService := service.UserService{}

	user,_ := userService.GetUser(id)
	
	result.Data = user
	result.Message = "成功"
	result.Status = 10000
	this.ResponseJSON(result)
}

/*
获取用户列表
*/
func (this *UserController) GetUserList() {
	result := result.NewFailedResultDataPage()

	userService := service.UserService{}

	userList,count,_ := userService.GetUserList()

	result.Data = userList
	result.Count = count
	result.Message = "成功"
	result.Status = 10000
	this.ResponseJSON(result)
}
/*
删除用户
*/
func (this *UserController) DeleteUser(){
	result := result.NewFailedResult()

	id,err := this.GetInt("ID")
	if(id == 0 || err != nil){
		result.Message = "传入的ID有误"
		this.ResponseJSON(result)
		return
	}
	userService := service.UserService{}

	if isOk,_ := userService.DeleteUser(id);!isOk{
		result.Message = "删除失败"
		this.ResponseJSON(result)
		return
	}
	result.Message = "成功"
	result.Status = 10000
	this.ResponseJSON(result)
}
