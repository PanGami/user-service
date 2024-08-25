package action

import (
	"context"

	"github.com/pangami/user-service/entity"
	"github.com/pangami/user-service/repo"
)

type ListUsers struct {
	repoUser repo.IUser
}

func NewListUsers(userRepo repo.IUser) *ListUsers {
	return &ListUsers{
		repoUser: userRepo,
	}
}

func (a *ListUsers) Handler(ctx context.Context, page, pageSize int) (*entity.ListUsersResponse, error) {
	users, totalCount, err := a.repoUser.List(ctx, page, pageSize)
	if err != nil {
		return nil, err
	}

	// Convert *entity.User to *entity.Data
	var userList []*entity.Data
	for _, u := range users {
		userList = append(userList, &entity.Data{
			ID:       u.ID,
			Username: u.Username,
		})
	}

	return &entity.ListUsersResponse{
		Users:      userList,
		TotalCount: int32(totalCount),
	}, nil
}
