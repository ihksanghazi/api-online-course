package services

import (
	"github.com/ihksanghazi/api-online-course/models"
	"gorm.io/gorm"
)

type CategoryService interface {
	FindAll() ([]models.CategoryResponse, error)
	FindById(id string) (models.CategoryWithClassResponse, error)
	Create(request *models.CategoryRequest) (models.CategoryResponse, error)
	Update(request *models.CategoryRequest, id string) (models.CategoryResponse, error)
	Delete(id string) error
}

type CategeryServiceImpl struct {
	DB *gorm.DB
}

func NewCategoryService(DB *gorm.DB) CategoryService {
	return &CategeryServiceImpl{
		DB: DB,
	}
}

func (c *CategeryServiceImpl) FindAll() ([]models.CategoryResponse, error) {
	var categories []models.Category
	var response []models.CategoryResponse

	err := c.DB.Model(&categories).Find(&response).Error
	return response, err
}

func (c *CategeryServiceImpl) FindById(id string) (models.CategoryWithClassResponse, error) {
	var category models.Category
	var response models.CategoryWithClassResponse

	err := c.DB.Model(&category).Find(&category, "id = ?", id).First(&response).Error
	return response, err
}

func (c *CategeryServiceImpl) Create(request *models.CategoryRequest) (models.CategoryResponse, error) {
	var category models.Category
	category.Name = request.Name
	var response models.CategoryResponse

	// transaction
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&category).Create(&category).First(&response, "id = ?", category.ID).Error; err != nil {
			// return any error will rollback
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})

	return response, err
}

func (c *CategeryServiceImpl) Update(request *models.CategoryRequest, id string) (models.CategoryResponse, error) {
	var category models.Category
	category.Name = request.Name
	var response models.CategoryResponse
	// transaction
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&category).Where("id = ?", id).Update("name", category.Name).First(&response, "id = ?", id).Error; err != nil {
			// return any error will rollback
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})

	return response, err
}

func (c *CategeryServiceImpl) Delete(id string) error {
	var category models.Category
	// transaction
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&category, "id = ?", id).Error; err != nil {
			// return any error will rollback
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})

	return err
}
