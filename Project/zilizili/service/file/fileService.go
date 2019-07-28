package service

import (
	"zilizili/dao/file"
	"zilizili/model/base/entity"
	"zilizili/common/tools"
	"mime/multipart"
	"os"
	"io"
)


var (
	fileDao dao.FileDao
)

type FileService struct{

}

/*
文件处理
*/
func (this *FileService) FileHandle(file multipart.File,fileHeader *multipart.FileHeader,fileType int) (*entity.File,error) {
	//文件信息
	var savePath  string = "static/upload/videoCover/"
	var fileName string =  fileHeader.Filename
	var fileID string = tools.GetGuid()
	var filePrefix string  = fileID + "_"
	//没有文件夹要先创建
	f, err := os.OpenFile(savePath + filePrefix + fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return nil,err
	}
	defer f.Close()
	io.Copy(f, file)

	fileEntity := &entity.File{}
	fileEntity.FileID = fileID
	fileEntity.FileName = fileName
	fileEntity.FileType = fileType
	fileEntity.AbsolutePath = ""
	fileEntity.VirtualAddress = savePath + filePrefix + fileName

	//添加
	if isOk,err := this.AddFile(*fileEntity); !isOk{
		return fileEntity,err
	} 
	return fileEntity,nil
}

/*
添加文件
*/
func (this *FileService) AddFile(file entity.File) (bool,error) {
	return fileDao.AddFile(file)
}