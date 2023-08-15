package models

import "github.com/google/uuid"

type Discussion struct {
	Model
	ModuleID uuid.UUID `gorm:"type:uuid;foreignKey" json:"module_id"`
	UserID   uuid.UUID `gorm:"type:uuid;foreignKey" json:"user_id"`
	Content  string    `json:"content_id"`
	ParentID uuid.UUID `gorm:"type:uuid" json:"parent_id"`
}
