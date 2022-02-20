package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func SetDatabase(db *gorm.DB) {
	DB = db

	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&Word{},
		&Category{},
		&Phrase{},
		&User{},
	)
	var defaultId string = "e2c450eb-b26c-49d4-8945-6d30e54dd2a6"
	var defaultCategory Category
	dbt := DB.Where("ID = ?", defaultId).Find(&defaultCategory)
	if dbt.RowsAffected == 0 {
		defaultCategories(defaultId)
	}
}

func defaultCategories(defaultId string) {
	DB.Create(Category{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
		Name:      "Day and Time",
		HindiName: "दिन एवं समय",
		Icon:      "61494",
	})
	DB.Create(Category{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
		Name:      "Frequent",
		HindiName: "नियमित",
		Icon:      "61560",
	})
	DB.Create(Category{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
		Name:      "Preposition",
		HindiName: "पूर्वसर्ग",
		Icon:      "59698",
	})
	DB.Create(Category{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
		Name:      "Relation",
		HindiName: "रिश्ते",
		Icon:      "61546",
	})
	DB.Create(Category{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
		Name:      "Months",
		HindiName: "महीने",
		Icon:      "58915",
	})
	DB.Create(Category{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
		Name:      "Kumaoni Food",
		HindiName: "कुमाऊँनी खाना",
		Icon:      "59237",
	})
	DB.Create(Category{
		ID:        defaultId,
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
		Name:      "Misc",
		HindiName: "विविध",
		Icon:      "59698",
	})
	DB.Create(Category{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
		Name:      "Fruits",
		HindiName: "फल",
		Icon:      "59237",
	})
	DB.Create(Category{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
		Name:      "Animals and Birds",
		HindiName: "जानवर एवं पक्षी",
		Icon:      "60989",
	})
}
