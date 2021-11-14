package controllers

import (
	"github.com/gin-gonic/gin"
	"quiz-app/models"
	"quiz-app/services"
)

type CourseQuizController interface {
	GetAllCourse() []models.CourseQuiz
	GetById(id int) (models.CourseQuiz, error)
	AddCourse(ctx *gin.Context) (models.CourseQuiz, error)
	UpdateCourse(id int, ctx *gin.Context) error
	DeleteCourse(id int) error
}

type courseQuizController struct {
	service services.CourseQuizService
}

func NewCourseQuizController(service services.CourseQuizService) CourseQuizController {
	return &courseQuizController{
		service: service,
	}
}

func (c *courseQuizController) GetAllCourse() []models.CourseQuiz {
	return c.service.GetAllCourse()
}

func (c *courseQuizController) GetById(id int) (models.CourseQuiz, error) {
	return c.service.GetById(id)
}

func (c *courseQuizController) AddCourse(ctx *gin.Context) (models.CourseQuiz, error) {
	var course models.CourseQuiz
	_ = ctx.ShouldBindJSON(&course)
	return c.service.AddCourse(course)
}

func (c *courseQuizController) UpdateCourse(id int, ctx *gin.Context) error {
	var newCourse models.CourseQuiz
	_ = ctx.Bind(&newCourse)
	return c.service.UpdateCourse(id, newCourse)
}

func (c *courseQuizController) DeleteCourse(id int) error {
	return c.service.DeleteCourse(id)
}
