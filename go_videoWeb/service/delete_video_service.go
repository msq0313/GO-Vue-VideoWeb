package service

import (
	"singo/model"
	"singo/serializer"
)

// DeleteVideoService 删除视频的服务
type DeleteVideoService struct {
}

// Delete 删除视频
func (service *DeleteVideoService) Delete(id string) serializer.Response {
	var video model.Video
	if err := model.DB.First(&video, id).Error; err != nil {
		return serializer.ParamErr("视频不存在", nil)
	}

	if err := model.DB.Delete(&video).Error; err != nil {
		return serializer.ParamErr("删除失败", nil)
	}

	return serializer.Response{}
}
