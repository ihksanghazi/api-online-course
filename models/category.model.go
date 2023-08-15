package models

type Category struct {
	Model
	Name string `gorm:"size:50;unique;not null" json:"name"`
	// relations
	Classes []Class `gorm:"foreignKey:CategoryID"`
}
