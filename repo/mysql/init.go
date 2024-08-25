package mysql

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var UnigoDB *gorm.DB

func InitCon() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	fmt.Printf("⇨ Connect MYSQL to Server %s:%s ... \n", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	log.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("⇨ Failed to connect to database: %s\n", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("⇨ Failed to get DB: %s\n", err)
	}

	// Set the maximum idle connections
	sqlDB.SetMaxIdleConns(10)

	// Set the maximum open connections
	sqlDB.SetMaxOpenConns(100)

	// Enable connection reuse
	sqlDB.SetConnMaxLifetime(time.Hour)

	UnigoDB = db
}
