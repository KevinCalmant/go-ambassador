package database

import (
	"ambassador/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	// _, err := gorm.Open(mysql.Open("root:root@tcp(db:3306)/ambassador"), &gorm.Config{})
	DB, err = gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/ambassador"), &gorm.Config{})
	if err != nil {
		panic("Could not connect with the database!")
	}
}

func Migrate() {
	err := DB.AutoMigrate(models.User{})
	if err != nil {
		panic("Migration has failed")
	}
}
