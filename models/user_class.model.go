package models

import "github.com/google/uuid"

type UserClass struct {
	Model
	UserID  uuid.UUID     `gorm:"type:uuid;foreignKey" json:"user_id"`
	ClassID uuid.UUID     `gorm:"type:uuid;foreignKey" json:"class_id"`
	Role    UserClassRole `gorm:"size:50" json:"role"`
}

type UserClassRole string

const (
	TeacherUserClassRole UserClassRole = "teacher"
	StudentUserClassRole UserClassRole = "student"
)
