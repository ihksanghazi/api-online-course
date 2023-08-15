package models

import "github.com/google/uuid"

type ChosenAnswer struct {
	Model
	QuestionID uuid.UUID `gorm:"type:uuid;foreignKey" json:"question_id"`
	AnswerText string    `json:"answer_text"`
	IsCorrect  bool      `json:"is_correct"`
	// relations
	UserAnswers []UserAnswer `gorm:"foreignKey:ChosenAnswerID"`
}
