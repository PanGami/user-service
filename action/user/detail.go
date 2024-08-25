package action

import (
	"context"

	"github.com/pangami/user-service/entity"
	"github.com/pangami/user-service/repo"
)

type DetailUser struct {
	repoUser repo.IUser
}

func NewDetailUser(userRepo repo.IUser) *DetailUser {
	return &DetailUser{
		repoUser: userRepo,
	}
}

func (a *DetailUser) Handler(ctx context.Context, user *entity.User) (*entity.User, error) {
	user, err := a.repoUser.Detail(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
