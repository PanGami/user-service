package grpc

import (
	"context"
	"log"

	action "github.com/pangami/user-service/action/user"
	builder "github.com/pangami/user-service/builder"
	user_service "github.com/pangami/user-service/repo"
	"github.com/pangami/user-service/repo/mysql"
	"github.com/pangami/user-service/repo/redis"
	user "github.com/pangami/user-service/transport/grpc/proto"
)

type GrpcServer struct {
	builder *builder.Grpc
	user.UnimplementedUserServer
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{
		builder: builder.NewGrpc(),
	}
}

func (s *GrpcServer) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.NoResponse, error) {
	service := user_service.NewUserRepository(mysql.DOTestDB)
	request := s.builder.CreateUserRequest(req)

	err := action.NewCreateUser(service).Handler(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &user.NoResponse{}, nil
}

func (s *GrpcServer) DetailUser(ctx context.Context, req *user.DetailUserRequest) (*user.DetailUserResponse, error) {
	service := user_service.NewUserRepository(mysql.DOTestDB)

	log.Println("REQ ID: ", req)

	// Create a request entity for the action handler
	request := s.builder.DetailUserRequest(req)

	log.Println("REQUEST ID: ", request.ID)

	// Check cache first
	userDetail, err := user_service.GetUserFromCache(ctx, redis.Client, int(request.ID))
	if err != nil {
		log.Println("Error fetching user from cache:", err)
	}

	if userDetail == nil {
		// Cache miss; fetch from DB
		log.Println("Cache miss; fetching from database")
		userDetail, err = action.NewDetailUser(service).Handler(ctx, &request)
		if err != nil {
			return nil, err
		}

		// Save the fetched user details to cache
		err = user_service.SaveUserToCache(ctx, redis.Client, userDetail)
		if err != nil {
			log.Println("Error saving user to cache:", err)
		}
	} else {
		log.Println("Cache hit; returning user from cache")
	}

	// Prepare and return the gRPC response
	response := &user.DetailUserResponse{
		Id:       int32(userDetail.ID),
		Username: userDetail.Username,
		FullName: userDetail.FullName,
		// Add more fields as necessary
	}

	return response, nil
}

func (s *GrpcServer) ListUsers(ctx context.Context, req *user.ListUsersRequest) (*user.ListUsersResponse, error) {
	service := user_service.NewUserRepository(mysql.DOTestDB)

	log.Println("REQ Page: ", req.Page)
	log.Println("REQ PageSize: ", req.PageSize)

	// Call the action to get the list of users
	entityResponse, err := action.NewListUsers(service).Handler(ctx, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}

	// Convert entity.ListUsersResponse to proto.ListUsersResponse
	var users []*user.Data
	for _, u := range entityResponse.Users {
		users = append(users, &user.Data{
			Id:       u.ID,
			Username: u.Username,
		})
	}

	return &user.ListUsersResponse{
		Users:      users,
		TotalCount: entityResponse.TotalCount,
	}, nil
}

func (s *GrpcServer) UpdateUser(ctx context.Context, req *user.CreateUserRequest) (*user.NoResponse, error) {
	service := user_service.NewUserRepository(mysql.DOTestDB)
	request := s.builder.UpdateUserRequest(req)

	err := action.NewUpdateUser(service).Handler(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &user.NoResponse{}, nil
}

func (s *GrpcServer) DeleteUser(ctx context.Context, req *user.DetailUserRequest) (*user.NoResponse, error) {
	service := user_service.NewUserRepository(mysql.DOTestDB)
	request := s.builder.DeleteUserRequest(req)

	_, err := action.NewDeleteUser(service).Handler(ctx, &request)
	if err != nil {
		return nil, err
	}

	return &user.NoResponse{}, nil
}
