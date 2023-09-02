package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	Model
	Name string `gorm:"size:50;unique;not null"`
	// relations
	Classes []Class `gorm:"foreignKey:CategoryID"`
}

type CategoryRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}

type CategoryWebResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *CategoryWebResponse) TableName() string {
	return "categories"
}

type CategoryWebResponseDetail struct {
	ID        uuid.UUID                    `json:"id"`
	Name      string                       `json:"name"`
	Classes   []ClassWebResponseNoCategory `gorm:"foreignKey:CategoryID;references:ID" json:"classes"`
	CreatedAt time.Time                    `json:"created_at"`
	UpdatedAt time.Time                    `json:"updated_at"`
}
