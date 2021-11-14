package controllers

import (
	"github.com/gin-gonic/gin"
	"quiz-app/models"
	"quiz-app/services"
)

type QuestionQuizController interface {
	GetAllData() []models.QuestionQuiz
	GetById(id int) (models.QuestionQuiz, error)
	AddData(ctx *gin.Context) (models.QuestionQuiz, error)
	UpdateData(id int, ctx *gin.Context) error
	DeleteData(id int) error
}

type questionQuizController struct {
	service services.QuestionQuizService
}

func NewQuestionQuizController(service services.QuestionQuizService) QuestionQuizController {
	return &questionQuizController{
		service: service,
	}
}

func (c *questionQuizController) AddData(ctx *gin.Context) (models.QuestionQuiz, error) {
	var data models.QuestionQuiz
	_ = ctx.ShouldBindJSON(&data)
	return c.service.AddData(data)
}

func (c *questionQuizController) UpdateData(id int, ctx *gin.Context) error {
	var data models.QuestionQuiz
	_ = ctx.Bind(&data)
	return c.service.UpdateData(id, data)
}

func (c *questionQuizController) DeleteData(id int) error {
	return c.service.DeleteData(id)
}

func (c *questionQuizController) GetAllData() []models.QuestionQuiz {
	return c.service.GetAllData()
}

func (c *questionQuizController) GetById(id int) (models.QuestionQuiz, error) {
	return c.service.GetById(id)
}
