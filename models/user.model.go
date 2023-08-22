package models

import "github.com/google/uuid"

type User struct {
	Model
	Username     string `gorm:"size:50;unique;not null" json:"username"`
	Email        string `gorm:"size:100;unique;not null" json:"email"`
	Password     string `gorm:"size:100" json:"password"`
	RefreshToken string `gorm:"size:255" json:"refresh_token"`
	ProfileUrl   string `gorm:"size:255" json:"profile_url"`
	Role         string `gorm:"not null;default:member" json:"role"`
	// Relation
	MessagesSent      []Message          `gorm:"foreignKey:SenderID"`
	MessagesReceived  []Message          `gorm:"foreignKey:RecipientID"`
	Classes           []Class            `gorm:"foreignKey:CreatedByID"`
	UserClasses       []UserClass        `gorm:"foreignKey:UserID"`
	Discussions       []Discussion       `gorm:"foreignKey:UserID"`
	UserQuizResponses []UserQuizResponse `gorm:"foreignKey:UserID"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"oneof='member' 'teacher' 'admin' '' "`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Role         string    `json:"role"`
	RefreshToken string    `json:"refresh_token"`
	ProfileUrl   string    `json:"profile_url"`
}
