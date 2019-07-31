package dao

import (
	"zilizili/model"
	"zilizili/model/base/entity"
	"zilizili/model/video/response"
	"zilizili/model/video/param"
)

type VideoDao struct{

}

/*
添加视频
*/
func (this *VideoDao) AddVideo(video entity.Video) (bool,error) {
	res := model.DB.Create(&video)
	if(res.Error != nil){
		return false,res.Error
	}
	return true,nil;
}

/*
更新视频
*/
func (this *VideoDao) UpdateVideo(video entity.Video) (bool,error){
	err := model.DB.Model(&video).Update(video).Error
	if(err != nil){
		return false,err
	} else {
		return true,nil
	}
}
/*
获取视频信息
*/
func (this *VideoDao) GetVideo(videoId int) (response.VideoResponseModel,error) {
	video := response.VideoResponseModel{}
	err := model.DB.First(&video,videoId).Error
	return video,err;
}

/*
获取视频信息列表
*/
func (this *VideoDao) GetVideoList(searchVideoModel param.SearchVideoModel) ([]response.VideoResponseModel,int,error) {

	videoList := []response.VideoResponseModel{} 

	var db = model.DB;
	if(searchVideoModel.Search != ""){
		db = model.DB.Where("(Title LIKE '%?%')",searchVideoModel.Search);
	}
	//总条数
	var count int
	db.Model(&videoList).Count(&count)
	//查出数据
	db = db.Order("ID DESC")
	err := db.Offset(searchVideoModel.PageSize * (searchVideoModel.PageIndex - 1)).Limit(searchVideoModel.PageSize).Find(&videoList).Error

	return videoList,count,err;
}
/*
删除视频信息
*/
func (this *VideoDao) DeleteVideo(videoId int)(bool,error){
	err := model.DB.Where("ID = ?", videoId).Delete(&entity.Video{}).Error
	if(err != nil){
		return false,err
	} else {
		return true,nil
	}
}
