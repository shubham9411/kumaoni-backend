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
	)
	DB.Create(Category{Name: "Day and Time", HindiName: "दिन एवं समय", Icon: "61494", CreatedAt: time.Now().Local(), UpdatedAt: time.Now().Local(), ID: uuid.New().String()})
	DB.Create(Category{Name: "Frequent", HindiName: "नियमित", Icon: "61560", CreatedAt: time.Now().Local(), UpdatedAt: time.Now().Local(), ID: uuid.New().String()})
	DB.Create(Category{Name: "Preposition", HindiName: "पूर्वसर्ग", Icon: "59698", CreatedAt: time.Now().Local(), UpdatedAt: time.Now().Local(), ID: uuid.New().String()})
	DB.Create(Category{Name: "Relation", HindiName: "रिश्ते", Icon: "61546", CreatedAt: time.Now().Local(), UpdatedAt: time.Now().Local(), ID: uuid.New().String()})
	DB.Create(Category{Name: "Months", HindiName: "महीने", Icon: "58915", CreatedAt: time.Now().Local(), UpdatedAt: time.Now().Local(), ID: uuid.New().String()})
	DB.Create(Category{Name: "Kumaoni Food", HindiName: "कुमाऊँनी खाना", Icon: "59237", CreatedAt: time.Now().Local(), UpdatedAt: time.Now().Local(), ID: uuid.New().String()})
	DB.Create(Category{Name: "Misc", HindiName: "विविध", Icon: "59698", CreatedAt: time.Now().Local(), UpdatedAt: time.Now().Local(), ID: "e2c450eb-b26c-49d4-8945-6d30e54dd2a6"})
	DB.Create(Category{Name: "Fruits", HindiName: "फल", Icon: "59237", CreatedAt: time.Now().Local(), UpdatedAt: time.Now().Local(), ID: uuid.New().String()})
	DB.Create(Category{Name: "Animals and Birds", HindiName: "जानवर एवं पक्षी", Icon: "60989", CreatedAt: time.Now().Local(), UpdatedAt: time.Now().Local(), ID: uuid.New().String()})
}
