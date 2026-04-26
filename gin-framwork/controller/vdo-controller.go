package controller

import (
	"github.com/devprajwalgotnochill/gin-frmwork/entity"
	"github.com/devprajwalgotnochill/gin-frmwork/services"
)

type VideoController interface {
	FindAll() []entity.Videos
	Save(video entity.Videos) entity.Videos
}

type controller struct {
	service services.VideoService
}

func New(service services.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Videos {
	return c.service.FindAll()
}

func (c *controller) Save(video entity.Videos) entity.Videos {
	return c.service.Save(video)
}
