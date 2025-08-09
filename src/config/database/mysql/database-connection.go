package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	DB_PASS  string
	dbConfig = map[string]string{
		"host":   "127.0.0.1",
		"port":   "3306",
		"user":   "rodrigodip",
		"dbname": "todo_api",
	}
)

func NewDataBaseConnection() error {

	DB_PASS = os.Getenv("SQL_DB_PASS")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig["user"],
		DB_PASS,
		dbConfig["host"],
		dbConfig["port"],
		dbConfig["dbname"],
	)

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
