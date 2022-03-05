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
	host := make(chan string, 1)

	if os.Getenv("GO_ENV") != "production" {
		user <- utils.GodotEnv("DB_USER")
		pass <- utils.GodotEnv("DB_PASSWORD")
		dbname <- utils.GodotEnv("DB_NAME")
		host <- utils.GodotEnv("DB_HOST")
	} else {
		user <- os.Getenv("DB_USER")
		pass <- os.Getenv("DB_PASSWORD")
		dbname <- os.Getenv("DB_NAME")
		host <- os.Getenv("DB_HOST")
	}

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", <-user, <-pass, <-host, <-dbname))
	if err != nil {
		log.Fatal("Error in connecting db :::", err)
	}

	if os.Getenv("GO_ENV") != "production" {
		db.LogMode(true)
	}
	return db
}
