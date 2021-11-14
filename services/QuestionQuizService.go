package services

import (
	"quiz-app/config/database"
	"quiz-app/models"
)

type QuestionQuizService interface {
	GetAllData() []models.QuestionQuiz
	GetById(id int) (models.QuestionQuiz, error)
	AddData(question models.QuestionQuiz) (models.QuestionQuiz, error)
	UpdateData(id int, newData models.QuestionQuiz) error
	DeleteData(id int) error
}

type questionQuizService struct{}

func NewQuestionQuizService() QuestionQuizService {
	return &questionQuizService{}
}

func (service *questionQuizService) GetAllData() []models.QuestionQuiz {
	var question []models.QuestionQuiz
	_ = database.DB.Find(&question)
	return question
}

func (service *questionQuizService) GetById(id int) (models.QuestionQuiz, error) {
	var question models.QuestionQuiz
	err := database.DB.First(&question, id).Error
	return question, err
}

func (service *questionQuizService) AddData(question models.QuestionQuiz) (models.QuestionQuiz, error) {
	err := database.DB.Create(&question).Error
	return question, err
}

func (service *questionQuizService) UpdateData(id int, newData models.QuestionQuiz) error {
	err := database.DB.Model(&newData).Where("id = ?", id).Updates(&newData).Error
	return err
}

func (serivce *questionQuizService) DeleteData(id int) error {
	err := database.DB.Delete(&models.QuestionQuiz{}, id).Error
	return err
}
