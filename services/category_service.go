package services

import (
	"github.com/ihksanghazi/api-online-course/models"
	"gorm.io/gorm"
)

type CategoryService interface {
	FindAll() ([]models.CategoryWebResponse, error)
	FindById(id string) (models.CategoryWebResponseDetail, error)
	Create(request *models.CategoryRequest) (models.CategoryRequest, error)
	Update(request *models.CategoryRequest, id string) (models.CategoryRequest, error)
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

func (c *CategeryServiceImpl) FindAll() ([]models.CategoryWebResponse, error) {
	var categories []models.Category
	var response []models.CategoryWebResponse

	err := c.DB.Model(&categories).Find(&response).Error
	return response, err
}

func (c *CategeryServiceImpl) FindById(id string) (models.CategoryWebResponseDetail, error) {
	var category models.Category
	var response models.CategoryWebResponseDetail

	err := c.DB.Model(&category).Preload("Classes.CreatedBy").Find(&response, "id = ?", id).Error
	return response, err
}

func (c *CategeryServiceImpl) Create(request *models.CategoryRequest) (models.CategoryRequest, error) {
	var category models.Category
	category.Name = request.Name

	// transaction
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&category).Create(&category).Error; err != nil {
			// return any error will rollback
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})

	return *request, err
}

func (c *CategeryServiceImpl) Update(request *models.CategoryRequest, id string) (models.CategoryRequest, error) {
	var category models.Category
	category.Name = request.Name
	// transaction
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&category).Where("id = ?", id).Update("name", category.Name).Error; err != nil {
			// return any error will rollback
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})

	return *request, err
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
