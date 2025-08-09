package database

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDataBaseConnection() error {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "rodrigodip:Mysql@121982mysql@tcp(127.0.0.1:3306)/todo_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("Error Ping")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Database connection established")
	DB = db
	return nil
}
func GetDB() (*gorm.DB, error) {

	sqlDB, err := DB.DB()
	if err != nil {
		panic("Error Getting DB")
	}

	if err = sqlDB.Ping(); err != nil {
		panic("error pinging DB")
	}

	return DB, nil
}
