package user

import (
	"fmt"
	"log"

	"github.com/anandu86130/Hospital-api-gateway/config"
	pb "github.com/anandu86130/Hospital-api-gateway/internal/user/pbU"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.UserServiceClient, error) {
	grpcAddr := fmt.Sprintf("user-service:%s", cfg.USERPORT)
	grpc, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing to grpc to client : %s", err.Error())
		return nil, err
	}
	log.Printf("Successfully connected to user client at port : %s", cfg.USERPORT)
	return pb.NewUserServiceClient(grpc), nil
}
