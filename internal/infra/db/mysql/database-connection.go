package mysql

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDataBaseConnection() (*gorm.DB, error) {

	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db_user,
		db_pass,
		db_host,
		db_port,
		db_name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Error Ping")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Database connection established")
	return db, nil
}

func GetDB() (*gorm.DB, error) {

	_, err := DB.DB()
	if err != nil {
		panic("Error Getting DB")
	}
	return DB, nil
}

func PingDB(db *gorm.DB) error {

	sqlDB, err := db.DB()
	if err != nil {
		log.Panic("Erro getting DB on Waiting for server")
	}

	maxAttempts := 10
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		if err = sqlDB.Ping(); err == nil {
			break
		}
		time.Sleep(time.Duration(attempts) * time.Second)
	}
	if err != nil {
		panic(fmt.Sprintf("Failed to connect database after %d attempts: %v", maxAttempts, err))
	}

	return nil
}
