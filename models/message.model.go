package models

import (
	"github.com/google/uuid"
)

type Message struct {
	Model
	SenderID    uuid.UUID `gorm:"type:uuid;foreignKey" json:"sender_id"`
	RecipientID uuid.UUID `gorm:"type:uuid;foreignKey" json:"recipient_id"`
	Content     string    `json:"content"`
}
