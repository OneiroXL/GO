package controllers

import (
	"zilizili/server/controllers/base"
	"zilizili/model/base/result"
	"zilizili/model/base/upload"
	"zilizili/service/file"
)

type UploadController struct {
	baseControllers.BaseController
}

/*
上传文件
*/
func (this *UploadController) UploadFile()  {
	result := result.NewFailedResultData()

	f, h, err := this.GetFile("file")
	//log.Println(this.Ctx.Request)
    if err != nil {
		result.Message = "未传入文件"
		this.ResponseJSON(result)
		return
    }
	defer f.Close()
	
	fileService := service.FileService{} 
	//文件处理
	entityFile,err := fileService.FileHandle(f,h,100)
	if(err != nil){
		result.Message = err
		this.ResponseJSON(result)
		return
	}

	uploadResultModel := upload.UploadResultModel{}
	uploadResultModel.FileID = entityFile.FileID
	uploadResultModel.FileName = entityFile.FileName
	uploadResultModel.VirtualAddress = entityFile.VirtualAddress

	result.Data = uploadResultModel
	result.Message = "成功"
	result.Status = 10000
	this.ResponseJSON(result)
}