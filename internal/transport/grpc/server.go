package grpc

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	userpb "github.com/Dmitriihub/project-protos/proto/user"
	"github.com/Dmitriihub/users-service/internal/user"
)

func RunGRPC(svc *user.Service) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("не удалось запустить listener: %w", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, NewHandler(svc))

	log.Println("gRPC сервер запущен на порту :50051")
	return grpcServer.Serve(lis)
}
