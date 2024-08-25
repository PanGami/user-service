package repo

import (
	"context"
	"log"

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
