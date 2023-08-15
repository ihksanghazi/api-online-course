package models

import (
	"time"

	"github.com/google/uuid"
)

type UserQuizResponse struct {
	Model
	UserID       uuid.UUID `gorm:"type:uuid;foreignKey" json:"user_id"`
	QuizID       uuid.UUID `gorm:"type:uuid;foreignKey" json:"quiz_id"`
	ResponseTime time.Time `json:"response_time"`
	Score        int       `json:"score"`
	// relations
	UserAnswers []UserAnswer `gorm:"foreignKey:ResponseID"`
}
