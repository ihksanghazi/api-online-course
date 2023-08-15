package models

import "github.com/google/uuid"

type Quiz struct {
	Model
	ClassID         uuid.UUID `gorm:"type:uuid;foreignKey" json:"class_id"`
	QuizTitle       string    `gorm:"size:255;not null" json:"quiz_title"`
	QuizDescription string    `json:"quiz_description"`
	// relations
	Questions         []Question         `gorm:"foreignKey:QuizID"`
	UserQuizResponses []UserQuizResponse `gorm:"foreignKey:QuizID"`
}
