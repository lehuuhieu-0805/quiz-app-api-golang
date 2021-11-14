package models

type CourseQuiz struct {
	Id       int        `json:"id" gorm:"primaryKey"`
	Name     string     `json:"name"`
	QuestionQuiz []QuestionQuiz `gorm:"foreignKey:CourseId;reference:Id" json:"-"`
}

func (CourseQuiz) TableName() string {
	return "course_quiz"
}
