package action

import (
	"context"

	"github.com/pangami/user-service/entity"
	"github.com/pangami/user-service/repo"
)

type DeleteUser struct {
	repoUser repo.IUser
}

func NewDeleteUser(userRepo repo.IUser) *DeleteUser {
	return &DeleteUser{
		repoUser: userRepo,
	}
}

func (a *DeleteUser) Handler(ctx context.Context, user *entity.User) (*entity.User, error) {
	user, err := a.repoUser.Delete(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
