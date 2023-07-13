package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var (
	db *gorm.DB
)

func ConnectDB(){

	d, err := gorm.Open("mysql", "")

	if err != nil {
		panic("failed to connect database")
	}

	db = d
}

// GetDB returns a handle to the DB object 
func GetDB() *gorm.DB {
	return db
}