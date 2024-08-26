package util

import (
	"errors"
	"log"
	"time"

	"github.com/pangami/user-service/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	DateFormatRFC3339 = time.RFC3339
	StatusInActive    = 0
	StatusActive      = 1
	StatusForgot      = 2
	StatusDraft       = "10"
	PaidStatusPro     = "pro"
	PaidStatusBasic   = "basic"
	SubscribePending  = "pending"
	SubscribePaid     = "paid"
	SubscribeExpired  = "expired"
	SubscribeFailed   = "failed"
	DefaultPage       = 1
	DefaultCount      = 15
	limitPage         = 9999
)

// hashPassword hashes the given plain-text password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash password")
	}
	return string(bytes), nil
}

// InsertMockData inserts default users data (mock) into the database
func InsertMockData(db *gorm.DB) {
	hashedPassword, _ := HashPassword("p4Nc4~Pass!")

	mockUsers := []entity.User{
		{Username: "panca", FullName: "Panca Nugraha Wicaksana", Password: hashedPassword},
		{Username: "nugraha", FullName: "Nugraha Wicaksana", Password: hashedPassword},
	}

	for _, user := range mockUsers {
		var existingUser entity.User

		result := db.Where("username = ?", user.Username).First(&existingUser)

		if result.Error == nil {
			log.Printf("Default user %s already exists, skipping insert.", user.Username)
			continue
		}

		if result.Error != gorm.ErrRecordNotFound {
			log.Printf("Error checking for existing user %s: %v", user.Username, result.Error)
			continue
		}

		err := db.Create(&user).Error

		if err == nil {
			log.Printf("Inserted default user %s", user.Username)
		}
	}
}
