package services

import (
	"github.com/devprajwalgotnochill/gin-frmwork/entity"
)

type VideoService interface {
	Save(entity.Videos) entity.Videos
	FindAll() []entity.Videos
}

type videoService struct {
	Videos []entity.Videos
}

func New() *videoService {
	return &videoService{}
}

func (services *videoService) Save(vdo entity.Videos) entity.Videos {
	services.Videos = append(services.Videos, vdo)
	return vdo
}

func (services *videoService) FindAll() []entity.Videos {
	return services.Videos
}
