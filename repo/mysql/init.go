package mysql

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pangami/user-service/entity"
	"github.com/pangami/user-service/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DOTestDB *gorm.DB

func InitCon() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	fmt.Printf("MYSQL Server %s:%s ... \n", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	log.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("CONNECT DB FAILED: %s\n", err)
	}

	// Automatically migrate the `User` struct to the database schema
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("DB FAILED : %s\n", err)
	}

	// Set the maximum idle connections
	sqlDB.SetMaxIdleConns(10)

	// Set the maximum open connections
	sqlDB.SetMaxOpenConns(100)

	// Enable connection reuse
	sqlDB.SetConnMaxLifetime(time.Hour)

	DOTestDB = db

	// Insert default users data (mock) into the database
	util.InsertMockData(db)
}
