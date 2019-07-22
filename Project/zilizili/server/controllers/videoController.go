package controllers

import (
	"./base"
	"../../model/base/result"
	"../../service/video"
	"../../model/video/param"
)

type VideoController struct {
	baseControllers.BaseController
}

//添加视频投稿
func (this *VideoController) AddVideo() {
	//返回消息
	result := result.Result{}
	//创建接收数据模型
	videoParamModel := param.VideoParamModel{}
	//检查惨参数并绑定
	if (!this.BindParseForm(&videoParamModel)) {
		result.Message = "参数有误"
		result.Status = 20000
		this.ResponseJSON(result)
		return
	}
	//验证参数
	if isSucceess,msg := this.ValidParam(videoParamModel);!isSucceess {
		result.Message = msg
		result.Status = 20000
		this.ResponseJSON(result)
		return
	}

	videoService := service.VideoService{}
	
	if isOk,err := videoService.AddVideo(videoParamModel);!isOk{
		result.Message = err
		result.Status = 20000
		this.ResponseJSON(result)
		return
	}
	result.Message = "成功"
	result.Status = 10000
	this.ResponseJSON(result)
}


//更新视频的接口
func (this *VideoController) UpdateVideo() {
	//返回消息
	result := result.Result{}
	//创建接收数据模型
	videoParamModel := param.VideoParamModel{}
	//检查惨参数并绑定
	if (!this.BindParseForm(&videoParamModel)) {
		result.Message = "参数有误"
		result.Status = 20000
		this.ResponseJSON(result)
		return
	}
	//验证参数
	if isSucceess,msg := this.ValidParam(videoParamModel);!isSucceess {
		result.Message = msg
		result.Status = 20000
		this.ResponseJSON(result)
		return
	}
	videoService := service.VideoService{}
	if isOk,err := videoService.UpdateVideo(videoParamModel);!isOk{
		result.Message = err
		result.Status = 20000
		this.ResponseJSON(result)
		return
	}
	result.Message = "成功"
	result.Status = 10000
	this.ResponseJSON(result)
}

//视频详情接口
func (this *VideoController) GetVideo() {
	//返回消息
	result := result.Result{}

	id,err := this.GetInt("ID")
	if(id == 0 || err != nil){
		result.Message = "传入的ID有误"
		result.Status = 20000
		this.ResponseJSON(result)
		return
	}
	videoService := service.VideoService{}
	video,_ := videoService.GetVideo(id)
	
	result.Data = video
	result.Message = "成功"
	result.Status = 10000
	this.ResponseJSON(result)
}

//视频列表接口
func (this *VideoController) GetVideoList() {
	//返回消息
	result := result.Result{}

	videoService := service.VideoService{}
	videoList,_ := videoService.GetVideoList()

	result.Data = videoList
	result.Message = "成功"
	result.Status = 10000
	this.ResponseJSON(result)
}


//删除视频的接口
func (this *VideoController) DeleteVideo() {
	//返回消息
	result := result.Result{}

	id,err := this.GetInt("ID")
	if(id == 0 || err != nil){
		result.Message = "传入的ID有误"
		result.Status = 20000
		this.ResponseJSON(result)
		return
	}
	videoService := service.VideoService{}
	if isOk,_ := videoService.DeleteVideo(id); !isOk{
		result.Message = "删除失败"
		result.Status = 10000
		this.ResponseJSON(result)
		return
	}

	result.Message = "成功"
	result.Status = 10000
	this.ResponseJSON(result)
}