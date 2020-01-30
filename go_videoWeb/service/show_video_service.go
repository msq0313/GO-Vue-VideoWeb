package service

import (
	"singo/model"
	"singo/serializer"
)

// ShowVideoService 视频详情的服务
type ShowVideoService struct {

}

// Show 视频详情
func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	//指针覆盖video
	if err := model.DB.First(&video, id).Error; err != nil {
		return serializer.ParamErr("视频不存在", nil)
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
