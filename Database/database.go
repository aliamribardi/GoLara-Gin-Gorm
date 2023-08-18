package Database

import (
	  "os"
	  "log"
  	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
    "github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectionDatabase() {
  godotenv.Load()

  DB_Name := os.Getenv("DB_Name")
	
  // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
  dsn := "root:@tcp(127.0.0.1:3306)/" + DB_Name + "?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
	log.Fatal("DB Connection Error")
  }

  DB = db
}