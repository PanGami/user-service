package util

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
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
