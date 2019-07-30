package service

import (
	"zilizili/dao/video"
	"zilizili/model/base/entity"
	"zilizili/model/video/param"
	"zilizili/model/video/response"
)


type VideoService struct{
	
}

//私有变量
var (
	videoDao dao.VideoDao
)
//初始化
func init(){
	videoDao = dao.VideoDao{}
}


/*
添加视频
*/
func (this *VideoService) AddVideo(videoParamModel param.VideoParamModel) (bool,error) {
	video := entity.Video{
		Title: videoParamModel.Title,
		Info:  videoParamModel.Info,
		URL:   videoParamModel.URL,
	}
	return videoDao.AddVideo(video);
}

/*
更新视频
*/
func (this *VideoService) UpdateVideo(videoParamModel param.VideoParamModel) (bool,error){
	video := entity.Video{
		Title: videoParamModel.Title,
		Info:  videoParamModel.Info,
		URL:   videoParamModel.URL,
	}
	video.BaseEntity.ID = videoParamModel.ID
	return videoDao.UpdateVideo(video);
}

/*
获取视频信息
*/
func (this *VideoService) GetVideo(videoId int) (response.VideoResponseModel,error) {
	return videoDao.GetVideo(videoId)
}

/*
获取视频信息列表
*/
func (this *VideoService) GetVideoList() ([]response.VideoResponseModel,error) {
	videoList,err := videoDao.GetVideoList()
	for index,_ := range videoList {
		videoList[index].CoverUrl = "http://127.0.0.1:3000/static/upload/videoCover/19dabca7465a5ee93b1369d8d68de39a_1.jpeg"
	}
	return videoList,err
}

/*
删除视频信息
*/
func (this *VideoService) DeleteVideo(videoId int)(bool,error){
	return videoDao.DeleteVideo(videoId);
}

