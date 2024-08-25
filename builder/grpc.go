package builder

import (
	"encoding/json"

	"github.com/pangami/user-service/entity"
	pb "github.com/pangami/user-service/transport/grpc/proto"
)

type Grpc struct{}

func NewGrpc() *Grpc {
	return &Grpc{}
}

func (g *Grpc) CreateUserRequest(reqPb *pb.CreateUserRequest) (data entity.User) {
	bytes, _ := json.Marshal(&reqPb)
	json.Unmarshal(bytes, &data)

	return data
}

func (g *Grpc) DetailUserRequest(reqPb *pb.DetailUserRequest) (data entity.User) {
	bytes, _ := json.Marshal(&reqPb)
	json.Unmarshal(bytes, &data)

	return data
}

func (g *Grpc) UpdateUserRequest(reqPb *pb.CreateUserRequest) (data entity.User) {
	bytes, _ := json.Marshal(&reqPb)
	json.Unmarshal(bytes, &data)

	return data
}
