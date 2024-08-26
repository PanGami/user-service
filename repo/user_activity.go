package repo

import (
	"context"
	"log"

	"github.com/pangami/user-service/entity"
	"gorm.io/gorm"
)

type UserActivityRepository struct {
	DB *gorm.DB
}

func NewUserActivityRepository(db *gorm.DB) *UserActivityRepository {
	return &UserActivityRepository{DB: db}
}

func (r *UserActivityRepository) Create(ctx context.Context, activity *entity.UserActivity) error {
	result := r.DB.Create(activity)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func (r *UserActivityRepository) GetActivitiesByUserID(ctx context.Context, userID int32) ([]*entity.UserActivity, error) {
	var activities []*entity.UserActivity
	result := r.DB.Where("user_id = ?", userID).Find(&activities)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return activities, nil
}

func (r *UserRepository) CreateUserWithActivities(ctx context.Context, user *entity.User, activities []*entity.UserActivity) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		for _, activity := range activities {
			activity.UserID = user.ID
			if err := tx.Create(activity).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
