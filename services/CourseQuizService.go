package services

import (
	"gorm.io/gorm/clause"
	"quiz-app/config/database"
	"quiz-app/models"
)

type CourseQuizService interface {
	GetAllCourse() []models.CourseQuiz
	GetById(id int) (models.CourseQuiz, error)
	AddCourse(course models.CourseQuiz) (models.CourseQuiz, error)
	UpdateCourse(id int, newCourse models.CourseQuiz) error
	DeleteCourse(id int) error
}

type courseQuizService struct{}

func NewCourseQuizService() CourseQuizService {
	return &courseQuizService{}
}

func (service *courseQuizService) GetAllCourse() []models.CourseQuiz {
	var courses []models.CourseQuiz
	_ = database.DB.Preload(clause.Associations).Find(&courses)
	return courses
}

func (service *courseQuizService) GetById(id int) (models.CourseQuiz, error) {
	var course models.CourseQuiz
	err := database.DB.First(&course, id).Error
	return course, err
}

func (service *courseQuizService) AddCourse(course models.CourseQuiz) (models.CourseQuiz, error) {
	err := database.DB.Create(&course).Error
	return course, err
}

func (service *courseQuizService) UpdateCourse(id int, newCourse models.CourseQuiz) error {
	err := database.DB.Model(&newCourse).Where("id = ?", id).Updates(&newCourse).Error
	return err
}

func (service *courseQuizService) DeleteCourse(id int) error {
	err := database.DB.Delete(&models.CourseQuiz{}, id).Error
	return err
}
