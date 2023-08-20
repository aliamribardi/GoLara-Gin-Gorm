package main

import (
	"GoLara-Gin-Gorm/Routes"
	"GoLara-Gin-Gorm/Database/Migrations"
	"GoLara-Gin-Gorm/Database"
)

func main() {
	Database.ConnectionDatabase()
	Migrations.Migration()
	Routes.Routes()
}