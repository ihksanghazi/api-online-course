package models

import "github.com/google/uuid"

type UserAnswer struct {
	Model
	ResponseID     uuid.UUID `gorm:"type:uuid;foreignKey" json:"response_id"`
	QuestionID     uuid.UUID `gorm:"type:uuid;foreignKey" json:"question_id"`
	ChosenAnswerID uuid.UUID `gorm:"type:uuid;foreignKey" json:"chosen_answer_id"`
	EssayResponse  string    `json:"essay_response"`
}
