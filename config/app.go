package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "gouser:gouser123@/gotest?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal("Error in connecting db :::", err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
