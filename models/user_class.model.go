package models

import "github.com/google/uuid"

type UserClass struct {
	Model
	UserID  uuid.UUID `gorm:"type:uuid;foreignKey" json:"user_id"`
	ClassID uuid.UUID `gorm:"type:uuid;foreignKey" json:"class_id"`
	Role    string    `gorm:"size:50" json:"role"`
}

type UserClassWebRequest struct {
	UserID  uuid.UUID `json:"user_id"`
	ClassID uuid.UUID `json:"class_id"`
}
