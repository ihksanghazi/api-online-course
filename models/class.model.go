package models

import "github.com/google/uuid"

type Class struct {
	Model
	Name        string    `gorm:"size:50;unique;not null" json:"name"`
	CreatedByID uuid.UUID `gorm:"type:uuid;foreignKey" json:"created_by"`
	CategoryID  uuid.UUID `gorm:"type:uuid;foreignKey" json:"category_id"`
	Description string    `json:"description"`
	Thumbnail   string    `gorm:"size:100" json:"thumbnail"`
	Trailer     string    `gorm:"size:100" json:"trailer"`
	// relations
	Quizzes      []Quiz        `gorm:"foreignKey:ClassID"`
	UserClasses  []UserClass   `gorm:"foreignKey:ClassID"`
	ClassModules []ClassModule `gorm:"foreignKey:ClassID"`
}
