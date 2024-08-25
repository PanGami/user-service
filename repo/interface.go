package repo

import (
	"context"

	"github.com/pangami/user-service/entity"
)

type i map[string]interface{}

type IUser interface {
	Create(ctx context.Context, user *entity.User) error
	Detail(ctx context.Context, user *entity.User) (*entity.User, error)
	List(ctx context.Context, page, pageSize int) ([]*entity.User, int, error)
	Update(ctx context.Context, user *entity.User) error
}
