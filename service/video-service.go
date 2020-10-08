package service

import "vaibhav/try-try-gg/entity"

// this provide two services
type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

// inorder to implement the interface
type videoService struct {
	videos []entity.Video
}

// default constructor - returning an implementation of the VideoService
func New() VideoService {
	// returning a pointer to the struct videoService
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
