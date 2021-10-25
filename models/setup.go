package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase(database *gorm.DB, err error) {
	database, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("unable to connect to database")
	}

	database.AutoMigrate(&User{}, &Transaction{})
}
