package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID        string     `gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index"`

	Name      string `gorm:"unique;not null" json:"name"`
	HindiName string `json:"hindiName"`
	Icon      string `json:"icon"`
	ImageUrl  string `json:"imageUrl"`
}

func (entity *Category) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Category) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}

func GetAllCategories() ([]Category, error) {
	var Categories []Category
	dbe := DB.Find(&Categories)
	if dbe.Error != nil {
		return nil, dbe.Error
	}
	return Categories, nil
}
