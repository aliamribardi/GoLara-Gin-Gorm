package Migrations

import (
	"fmt"
	"GoLara-Gin-Gorm/Models"
	"GoLara-Gin-Gorm/Database"
)

func Migration () {
	err := Database.DB.AutoMigrate(
		&Models.User{},
		&Models.Kelas{},
	)

	if err != nil {
		fmt.Println("Can't running Migrations")
	}

	fmt.Println("Migration Success")
}