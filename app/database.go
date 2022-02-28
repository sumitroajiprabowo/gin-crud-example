// models/setup.go
package app

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SetupDB : initializing mysql database
func SetupDB() *gorm.DB {
	USER := "root"
	PASS := "kmzway87aa"
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "learn_gin_gonic_crud"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open(mysql.Open(URL))
	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("Database connection established")
	return db
}
