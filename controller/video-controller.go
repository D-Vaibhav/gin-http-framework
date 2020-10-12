package controller

import (
	"vaibhav/try-try-gg/entity"
	"vaibhav/try-try-gg/service"
	"vaibhav/try-try-gg/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// controller will handle various types of requests like - GET, POST
type VideoController interface {
	Save(ctx *gin.Context) error
	FindAll() []entity.Video
}

type controller struct {
	service service.VideoService
}

// for validation
var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New() // for validation
	validate.RegisterValidation("is-cool", validators.ValidateIsCool)

	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	// for validation error handling
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}
