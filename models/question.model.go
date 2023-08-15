package models

import "github.com/google/uuid"

type Question struct {
	Model
	QuizID       uuid.UUID `gorm:"type:uuid;foreignKey" json:"quiz_id"`
	QuestionText string    `json:"question_text"`
	QuestionType string    `gorm:"size:50" json:"question_type"`
	// relations
	ChosenAnswers []ChosenAnswer `gorm:"foreignKey:QuestionID"`
	UserAnswers   []UserAnswer   `gorm:"foreignKey:QuestionID"`
}
