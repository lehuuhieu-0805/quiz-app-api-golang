package models

type QuestionQuiz struct {
	Id            int     `json:"id" gorm:"primaryKey"`
	Question      string  `json:"question"`
	AnswerA       string  `json:"answerA"`
	AnswerB       string  `json:"answerB"`
	AnswerC       string  `json:"answerC"`
	AnswerD       string  `json:"answerD"`
	CorrectAnswer string  `json:"correctAnswer"`
	Point         float64 `json:"point"`
	CourseId      int     `json:"courseId"`
}

func (QuestionQuiz) TableName() string {
	return "question_quiz"
}
