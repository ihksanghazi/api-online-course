package models

type Module struct {
	Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	VideoURL string `gorm:"size:100" json:"video_url"`
	// relations
	ClassModules []ClassModule `gorm:"foreignKey:ModuleID"`
	Discussions  []Discussion  `gorm:"foreignKey:ModuleID"`
}
