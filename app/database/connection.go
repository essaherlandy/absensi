package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:@/absensi"), &gorm.Config{})

	if err != nil {
		panic("Could not connect database")
	}

	DB = connection

	// connection.AutoMigrate(&models.User)
}
