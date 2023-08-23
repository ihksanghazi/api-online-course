package services

import (
	"github.com/ihksanghazi/api-online-course/models"
	"gorm.io/gorm"
)

type ClassService interface {
	Create(request models.ClassWebRequest) (models.Class, error)
}

func NewClassService(DB *gorm.DB) ClassService {
	return &ClassServiceImpl{
		DB: DB,
	}
}

type ClassServiceImpl struct {
	DB *gorm.DB
}

func (c *ClassServiceImpl) Create(request models.ClassWebRequest) (models.Class, error) {
	var class models.Class
	var userClass models.UserClass

	// transaction
	errTransaction := c.DB.Transaction(func(tx *gorm.DB) error {
		class.Name = request.Name
		class.CreatedByID = request.CreatedByID
		class.CategoryID = request.CategoryID
		class.Description = request.Description
		class.Thumbnail = request.Thumbnail
		class.Trailer = request.Trailer

		if err := tx.Model(&class).Create(&class).Error; err != nil {
			return err
		}
		userClass.UserID = request.CreatedByID
		userClass.ClassID = class.ID
		userClass.Role = "teacher"

		if err := tx.Model(&userClass).Create(&userClass).Error; err != nil {
			return err
		}
		return nil
	})

	return class, errTransaction
}
