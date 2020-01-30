package service

import (
	"singo/model"
	"singo/serializer"
)

// ListVideoService 视频列表的服务
type ListVideoService struct {

}

// List 视频列表
func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video
	//指针覆盖video
	if err := model.DB.Find(&videos).Error; err != nil {
		return serializer.ParamErr("视频不存在", nil)
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
