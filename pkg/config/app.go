package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	moviedb, err := gorm.Open("mysql", "daksh:pwd@/movie?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = moviedb
}

func GetDB() *gorm.DB {
	return db
}
