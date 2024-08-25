package grpc

import (
	"fmt"
	"log"
	"net"
	"os"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/reflection"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	pb "github.com/pangami/user-service/transport/grpc/proto"
)

// Run handler running grpc
func Run() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	} else {
		fmt.Printf(" Connect grpc to port : %s \n", os.Getenv("GRPC_PORT"))
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		apmgrpc.NewUnaryServerInterceptor(apmgrpc.WithRecovery()),
	)))

	hsrv := health.NewServer()
	hsrv.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(s, hsrv)

	pb.RegisterUserServer(s, NewGrpcServer())

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
