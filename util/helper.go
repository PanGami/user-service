package util

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/pangami/user-service/entity"
	"github.com/pangami/user-service/repo"
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

// InsertMockData inserts default users and their activities (mock) into the database
func InsertMockData(db *gorm.DB, userRepo *repo.UserRepository) {
	if db == nil {
		log.Fatalf("Database connection is nil")
	}

	if userRepo == nil {
		log.Fatalf("UserRepository is nil")
	}

	if db == nil {
		log.Fatalf("GORM DB instance is nil. Check if the database connection was initialized correctly.")
	}

	hashedPassword, _ := HashPassword("p4Nc4~Pass!")

	mockUsers := []struct {
		User       entity.User
		Activities []entity.UserActivity
	}{
		{
			User: entity.User{
				Username: "panca",
				FullName: "Panca Nugraha Wicaksana",
				Password: hashedPassword,
			},
			Activities: []entity.UserActivity{
				{Action: "Login", Timestamp: time.Now()},
				{Action: "View Dashboard", Timestamp: time.Now()},
			},
		},
		{
			User: entity.User{
				Username: "nugraha",
				FullName: "Nugraha Wicaksana",
				Password: hashedPassword,
			},
			Activities: []entity.UserActivity{
				{Action: "Login", Timestamp: time.Now()},
				{Action: "Edit Profile", Timestamp: time.Now()},
			},
		},
	}

	for _, mock := range mockUsers {
		var existingUser entity.User

		result := db.Where("username = ?", mock.User.Username).First(&existingUser)

		if result.Error == nil {
			log.Printf("Default user %s already exists, skipping insert.", mock.User.Username)
			continue
		}

		if result.Error != gorm.ErrRecordNotFound {
			log.Printf("Error checking for existing user %s: %v", mock.User.Username, result.Error)
			continue
		}

		// Use CreateUserWithActivities method to insert user and activities
		activities := make([]*entity.UserActivity, len(mock.Activities))
		for i, activity := range mock.Activities {
			activities[i] = &activity
		}
		err := userRepo.CreateUserWithActivities(context.Background(), &mock.User, activities)
		if err == nil {
			log.Printf("Inserted default user %s with activities", mock.User.Username)
		} else {
			log.Printf("Failed to insert user %s with activities: %v", mock.User.Username, err)
		}
	}
}
