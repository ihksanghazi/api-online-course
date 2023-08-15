package models

import "github.com/google/uuid"

type ClassModule struct {
	Model
	ClassID  uuid.UUID `gorm:"type:uuid;foreignKey" json:"class_id"`
	ModuleID uuid.UUID `gorm:"type:uuid;foreignKey" json:"module_id"`
}
