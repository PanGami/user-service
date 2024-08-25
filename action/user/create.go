package action

import (
	"context"

	"github.com/pangami/user-service/entity"
	"github.com/pangami/user-service/repo"
)

type CreateUser struct {
	repoUser repo.IUser
}

func NewCreateUser(userRepo repo.IUser) *CreateUser {
	return &CreateUser{
		repoUser: userRepo,
	}
}

func (a *CreateUser) Handler(ctx context.Context, user *entity.User) error {
	err := a.repoUser.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
