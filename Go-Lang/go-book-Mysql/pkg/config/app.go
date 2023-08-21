package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "nakul:00000000/simplerest?charset=utf8&parseTime=True&loc=local")
	if err != nil {
		panic(err)
	}
	db = d
}

func Getdb() *gorm.DB {
	return db
}

//func Getdb() *gorm.DB: This line defines a function named Getdb that returns a pointer to a gorm.DB type. The *gorm.DB type represents a GORM database connection.
