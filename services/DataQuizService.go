package services

import (
	"quiz-app/config/database"
	"quiz-app/models"
)

type DataQuizService interface {
	GetAllData() []models.DataQuiz
	GetById(id int) (models.DataQuiz, error)
	AddData(data models.DataQuiz) (models.DataQuiz, error)
	UpdateData(id int, newData models.DataQuiz) error
	DeleteData(id int) error
}

type dataQuizService struct{}

func NewDataQuizService() DataQuizService {
	return &dataQuizService{}
}

func (service *dataQuizService) GetAllData() []models.DataQuiz {
	var data []models.DataQuiz
	_ = database.DB.Find(&data)
	return data
}

func (service *dataQuizService) GetById(id int) (models.DataQuiz, error) {
	var data models.DataQuiz
	err := database.DB.First(&data, id).Error
	return data, err
}

func (service *dataQuizService) AddData(data models.DataQuiz) (models.DataQuiz, error) {
	err := database.DB.Create(&data).Error
	return data, err
}

func (service *dataQuizService) UpdateData(id int, newData models.DataQuiz) error {
	err := database.DB.Model(&newData).Where("id = ?", id).Updates(&newData).Error
	return err
}

func (serivce *dataQuizService) DeleteData(id int) error {
	err := database.DB.Delete(&models.DataQuiz{}, id).Error
	return err
}
