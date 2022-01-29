package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shubham9411/kumaoni-backend/utils"
)

func Connect() *gorm.DB {
	user := make(chan string, 1)
	pass := make(chan string, 1)
	dbname := make(chan string, 1)

	if os.Getenv("GO_ENV") != "production" {
		user <- utils.GodotEnv("DB_USER")
		pass <- utils.GodotEnv("DB_PASSWORD")
		dbname <- utils.GodotEnv("DB_NAME")
	} else {
		user <- os.Getenv("DB_USER")
		pass <- os.Getenv("DB_PASSWORD")
		dbname <- os.Getenv("DB_NAME")
	}

	// db, err := gorm.Open(postgres.Open(<-databaseURI), &gorm.Config{})
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local", <-user, <-pass, <-dbname))

	if err != nil {
		log.Fatal("Error in connecting db :::", err)
	}

	if os.Getenv("GO_ENV") != "production" {
		db.LogMode(true)
	}
	return db
}
