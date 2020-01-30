package service

import (
	"singo/model"
	"singo/serializer"
)

// CreateVideoService 更新视频的服务
type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info string `form:"info" json:"info" binding:"max=300"`
}

// Update 更新视频
func (service *UpdateVideoService) Update(id string) serializer.Response {
	var video model.Video

	if err := model.DB.First(&video, id).Error; err != nil {
		return serializer.ParamErr("视频不存在", nil)
	}

	video.Title = service.Title
	video.Info = service.Info

	if err := model.DB.Save(&video).Error; err != nil {
		return serializer.ParamErr("视频更新失败", nil)
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
