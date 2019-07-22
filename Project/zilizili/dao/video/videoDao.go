package dao

import (
	"../../model"
	"../../model/base/entity"
	"../../model/video/response"
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
func (this *VideoDao) GetVideoList() ([]response.VideoResponseModel,error) {
	videoList := []response.VideoResponseModel{}
	err := model.DB.Find(&videoList).Error
	return videoList,err;
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