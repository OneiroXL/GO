package controllers

import (
	"xfund/server/controllers/base"
)

type UserController struct {
	baseControllers.BaseController
}

/*
获取用户列表
*/
func (this *UserController) GetUserList() {
	this.ResponseJSON("")
}