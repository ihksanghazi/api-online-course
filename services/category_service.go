package services

import (
	"github.com/ihksanghazi/api-online-course/models"
	"gorm.io/gorm"
)

type CategoryService interface {
	FindAll(categories *[]models.Category) ([]models.Category, error)
	FindById(category *models.Category, id string) (models.Category, error)
	Create(CategoryModel *models.Category) (models.Category, error)
	Update(category *models.Category, id string) (models.Category, error)
	Delete(category *models.Category, id string) (models.Category, error)
}

type CategeryServiceImpl struct {
	DB *gorm.DB
}

func NewCategoryService(DB *gorm.DB) CategoryService {
	return &CategeryServiceImpl{
		DB: DB,
	}
}

func (c *CategeryServiceImpl) FindAll(categories *[]models.Category) ([]models.Category, error) {
	err := c.DB.Find(&categories).Error
	return *categories, err
}

func (c *CategeryServiceImpl) FindById(category *models.Category, id string) (models.Category, error) {
	err := c.DB.Find(&category, "id = ?", id).First(&category).Error
	return *category, err
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

func (c *CategeryServiceImpl) Update(category *models.Category, id string) (models.Category, error) {
	// transaction
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&category).Where("id = ?", id).Update("name", category.Name).Error; err != nil {
			// return any error will rollback
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})

	return *category, err
}

func (c *CategeryServiceImpl) Delete(category *models.Category, id string) (models.Category, error) {
	// transaction
	err := c.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&category, "id = ?", id).Error; err != nil {
			// return any error will rollback
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})

	return *category, err
}
