package dao

import (
	"zilizili/model"
	"zilizili/model/base/entity"
)


type FileDao struct {

}

/*
添加文件
*/
func (this *FileDao) AddFile(file entity.File) (bool,error) {
	res := model.DB.Create(&file)
	if (res.Error != nil){
		return false,res.Error
	}
	return true,nil;
}

