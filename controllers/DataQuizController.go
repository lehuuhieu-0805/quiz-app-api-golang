package controllers

import (
	"github.com/gin-gonic/gin"
	"quiz-app/models"
	"quiz-app/services"
)

type DataQuizController interface {
	GetAllData() []models.DataQuiz
	GetById(id int) (models.DataQuiz, error)
	AddData(ctx *gin.Context) (models.DataQuiz, error)
	UpdateData(id int, ctx *gin.Context) error
	DeleteData(id int) error
}

type dataQuizController struct {
	service services.DataQuizService
}

func NewDataQuizController(service services.DataQuizService) DataQuizController {
	return &dataQuizController{
		service: service,
	}
}

func (c *dataQuizController) AddData(ctx *gin.Context) (models.DataQuiz, error) {
	var data models.DataQuiz
	_ = ctx.ShouldBindJSON(&data)
	return c.service.AddData(data)
}

func (c *dataQuizController) UpdateData(id int, ctx *gin.Context) error {
	var data models.DataQuiz
	_ = ctx.Bind(&data)
	return c.service.UpdateData(id, data)
}

func (c *dataQuizController) DeleteData(id int) error {
	return c.service.DeleteData(id)
}

func (c *dataQuizController) GetAllData() []models.DataQuiz {
	return c.service.GetAllData()
}

func (c *dataQuizController) GetById(id int) (models.DataQuiz, error) {
	return c.service.GetById(id)
}
