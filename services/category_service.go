package services

import (
	"github.com/ihksanghazi/api-online-course/models"
	"gorm.io/gorm"
)

type CategoryService interface {
	FindAll() ([]models.Category, error)
	FindById(CategoryModel *models.Category) (models.Category, error)
	Create(CategoryModel *models.Category) (models.Category, error)
	Update(CategoryModel *models.Category) (models.Category, error)
	Delete(CategoryModel *models.Category) (models.Category, error)
}

type CategeryServiceImpl struct {
	DB *gorm.DB
}

func NewCategoryService(DB *gorm.DB) CategoryService {
	return &CategeryServiceImpl{
		DB: DB,
	}
}

func (c *CategeryServiceImpl) FindAll() ([]models.Category, error) {
	var categories []models.Category
	err := c.DB.Find(&categories).Error
	return categories, err
}

func (c *CategeryServiceImpl) FindById(CategoryModel *models.Category) (models.Category, error) {
	var category models.Category
	err := c.DB.Find(&category, CategoryModel.ID).Error
	return category, err
}

func (c *CategeryServiceImpl) Create(CategoryModel *models.Category) (models.Category, error) {
	// transaction
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&CategoryModel).Error; err != nil {
			// return any error will rollback
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})

	return *CategoryModel, err
}

func (c *CategeryServiceImpl) Update(CategoryModel *models.Category) (models.Category, error) {
	// transaction
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&CategoryModel).Where("id = ?", CategoryModel.ID).Update("name", CategoryModel.Name).Error; err != nil {
			// return any error will rollback
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})

	return *CategoryModel, err
}

func (c *CategeryServiceImpl) Delete(CategoryModel *models.Category) (models.Category, error) {
	// transaction
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&CategoryModel, CategoryModel.ID).Error; err != nil {
			// return any error will rollback
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})

	return *CategoryModel, err
}
