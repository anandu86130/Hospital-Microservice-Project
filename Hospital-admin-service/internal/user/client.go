package user

import (
	"fmt"
	"log"

	"github.com/anandu86130/Hospital-admin-service/config"
	pb "github.com/anandu86130/Hospital-admin-service/internal/user/pbU"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.UserServiceClient, error) {
	grpcAddr := fmt.Sprintf("user-service:%s", cfg.GrpcUserPort)
	grpc, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc user client: %s, ", cfg.GrpcUserPort)
		return nil, err
	}
	log.Printf("Succesfully connected to user client at port: %v", cfg.GrpcUserPort)
	return pb.NewUserServiceClient(grpc), nil
}
