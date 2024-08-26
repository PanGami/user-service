package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pangami/user-service/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	result := r.DB.Create(user)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

// create function to get user detail not just error
func (r *UserRepository) Detail(ctx context.Context, user *entity.User) (*entity.User, error) {
	result := r.DB.Where("id = ?", user.ID).First(user)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return user, nil
}

func (ur *UserRepository) List(ctx context.Context, page, pageSize int) ([]*entity.User, int, error) {
	// Implement the logic to fetch users from the database with pagination
	var users []*entity.User
	ur.DB.Find(&users)

	totalCount := len(users) // Total count of users
	start := (page - 1) * pageSize
	end := start + pageSize
	if end > totalCount {
		end = totalCount
	}

	return users[start:end], totalCount, nil
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) error {
	result := r.DB.Save(user)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, user *entity.User) (*entity.User, error) {
	result := r.DB.Where("id = ?", user.ID).Delete(user)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return user, nil
}

// SaveUserToCache caches user details in Redis with an expiration time
func SaveUserToCache(ctx context.Context, redisClient *redis.Client, userDetail *entity.User) error {
	cacheKey := "user_detail_" + string(userDetail.ID)
	userData, err := json.Marshal(userDetail)
	if err != nil {
		return err
	}

	err = redisClient.Set(ctx, cacheKey, userData, 5*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetUserFromCache retrieves user details from Redis
func GetUserFromCache(ctx context.Context, redisClient *redis.Client, userID int) (*entity.User, error) {
	cacheKey := "user_detail_" + fmt.Sprint(userID)
	cachedUser, err := redisClient.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		return nil, nil // Cache miss
	} else if err != nil {
		return nil, err
	}

	var userDetail entity.User
	err = json.Unmarshal([]byte(cachedUser), &userDetail)
	if err != nil {
		return nil, err
	}

	return &userDetail, nil
}
