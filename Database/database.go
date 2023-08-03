package Database

import (
	"log"
  	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	
  // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
  dsn := "root:@tcp(127.0.0.1:3306)/go-crud-gin?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
	log.Fatal("DB Connection Error")
  }

  DB = db
}