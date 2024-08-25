package action

import (
	"context"

	"github.com/pangami/user-service/entity"
	"github.com/pangami/user-service/repo"
)

type UpdateUser struct {
	repoUser repo.IUser
}

func NewUpdateUser(userRepo repo.IUser) *UpdateUser {
	return &UpdateUser{
		repoUser: userRepo,
	}
}

func (a *UpdateUser) Handler(ctx context.Context, user *entity.User) error {
	err := a.repoUser.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
