package models

import (
	"time"

	"github.com/google/uuid"
)

type Class struct {
	Model
	Name        string    `gorm:"size:50;unique;not null" json:"name"`
	CreatedByID uuid.UUID `gorm:"type:uuid;foreignKey" json:"created_by"`
	CategoryID  uuid.UUID `gorm:"type:uuid;foreignKey" json:"category_id"`
	Description string    `json:"description"`
	Thumbnail   string    `gorm:"size:100" json:"thumbnail"`
	Trailer     string    `gorm:"size:100" json:"trailer"`
	//other
	Students []User `gorm:"many2many:user_classes;foreignKey:ID;joinForeignKey:ClassID;References:ID;joinReferences:UserID" json:"students"`
	// relations
	Quizzes      []Quiz        `gorm:"foreignKey:ClassID"`
	UserClasses  []UserClass   `gorm:"foreignKey:ClassID"`
	ClassModules []ClassModule `gorm:"foreignKey:ClassID"`
}

type ClassWebRequest struct {
	Name        string    `json:"name" validate:"required"`
	CreatedByID uuid.UUID `json:"created_by" validate:"required"`
	CategoryID  uuid.UUID `json:"category_id" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Thumbnail   string    `json:"thumbnail"`
	Trailer     string    `json:"trailer"`
}

type ClassWebResponse struct {
	ID          uuid.UUID           `json:"id"`
	Name        string              `json:"name"`
	CreatedByID uuid.UUID           `json:"-"`
	CreatedBy   UserWebResponse     `gorm:"foreignKey:CreatedByID" json:"created_by"`
	CategoryID  uuid.UUID           `json:"-"`
	Category    CategoryWebResponse `gorm:"foreignKey:CategoryID" json:"category"`
	Description string              `json:"description"`
	Thumbnail   string              `json:"thumbnail"`
	CreatedAt   time.Time           `json:"created_at"`
}

func (c *ClassWebResponse) TableName() string {
	return "classes"
}

type ClassWebResponseNoCategory struct {
	ID          uuid.UUID       `json:"id"`
	Name        string          `json:"name"`
	CreatedByID uuid.UUID       `json:"-"`
	CreatedBy   UserWebResponse `gorm:"foreignKey:CreatedByID" json:"created_by"`
	CategoryID  uuid.UUID       `json:"-"`
	Description string          `json:"description"`
	Thumbnail   string          `json:"thumbnail"`
	CreatedAt   time.Time       `json:"created_at"`
}

func (c *ClassWebResponseNoCategory) TableName() string {
	return "classes"
}

type ClassWebResponseNoCreatedBy struct {
	ID          uuid.UUID           `json:"id"`
	Name        string              `json:"name"`
	CreatedByID uuid.UUID           `json:"-"`
	CategoryID  uuid.UUID           `json:"-"`
	Category    CategoryWebResponse `gorm:"foreignKey:CategoryID" json:"category"`
	Description string              `json:"description"`
	Thumbnail   string              `json:"thumbnail"`
	CreatedAt   time.Time           `json:"created_at"`
}

func (c *ClassWebResponseNoCreatedBy) TableName() string {
	return "classes"
}

type ClassWebResponseDetail struct {
	ID          uuid.UUID           `json:"id"`
	Name        string              `json:"name"`
	CreatedByID uuid.UUID           `json:"-"`
	CreatedBy   UserWebResponse     `gorm:"foreignKey:CreatedByID" json:"created_by"`
	CategoryID  uuid.UUID           `json:"-"`
	Category    CategoryWebResponse `gorm:"foreignKey:CategoryID" json:"category"`
	Description string              `json:"description"`
	Thumbnail   string              `json:"thumbnail"`
	Trailer     string              `json:"trailer"`
	Members     []UserWebResponse   `gorm:"many2many:user_classes;foreignKey:ID;joinForeignKey:ClassID;References:ID;joinReferences:UserID" json:"members"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}

func (c *ClassWebResponseDetail) TableName() string {
	return "classes"
}
