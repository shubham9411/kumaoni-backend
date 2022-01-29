package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name      string `gorm:"" json:"name"`
	HindiName string `json:"hindiName"`
	Icon      string `json:"icon"`
	ImageUrl  string `json:"imageUrl"`
}
