package services

import "gorm.io/gorm"

type ClassService interface {
}

func NewClassService(DB *gorm.DB) ClassService {
	return &ClassServiceImpl{
		DB: DB,
	}
}

type ClassServiceImpl struct {
	DB *gorm.DB
}
