package entity

import (
	"time"
)

type UserActivity struct {
	ID        int32     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int32     `gorm:"not null" json:"user_id"`
	Action    string    `gorm:"size:255;not null" json:"action"`
	Timestamp time.Time `gorm:"not null" json:"timestamp"`
}

// UserActivityResponse struct to include in User details response
type UserActivityResponse struct {
	ID        int32  `json:"id"`
	Action    string `json:"action"`
	Timestamp string `json:"timestamp"`
}
