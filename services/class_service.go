package services

import (
	"github.com/ihksanghazi/api-online-course/models"
	"gorm.io/gorm"
)

type ClassService interface {
	Create(request models.ClassWebRequest) (models.ClassWebRequest, error)
	GetAll() ([]models.ClassWebResponse, error)
	GetById(classId string) (models.ClassWebResponseDetail, error)
	AddClass(request models.UserClassWebRequest) (models.UserClassWebRequest, error)
}

func NewClassService(DB *gorm.DB) ClassService {
	return &ClassServiceImpl{
		DB: DB,
	}
}

type ClassServiceImpl struct {
	DB *gorm.DB
}

func (c *ClassServiceImpl) Create(request models.ClassWebRequest) (models.ClassWebRequest, error) {
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

	return request, errTransaction
}

func (c *ClassServiceImpl) GetAll() ([]models.ClassWebResponse, error) {
	var class []models.Class
	var response []models.ClassWebResponse

	errModel := c.DB.Model(&class).Preload("CreatedBy").Preload("Category").Find(&response).Error

	return response, errModel
}

func (c *ClassServiceImpl) GetById(classId string) (models.ClassWebResponseDetail, error) {
	var class models.Class
	var response models.ClassWebResponseDetail

	err := c.DB.Model(&class).Preload("CreatedBy").Preload("Category").Preload("Members").First(&response, "id = ?", classId).Error

	return response, err
}

func (c *ClassServiceImpl) AddClass(request models.UserClassWebRequest) (models.UserClassWebRequest, error) {
	var userClass models.UserClass
	userClass.ClassID = request.ClassID
	userClass.UserID = request.UserID
	userClass.Role = "student"

	errTransaction := c.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&userClass).Create(&userClass).Error; err != nil {
			return err
		}
		return nil
	})

	return request, errTransaction
}
