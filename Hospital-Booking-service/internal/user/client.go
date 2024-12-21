package user

import (
	"booking-service/config"
	"fmt"
	"log"

	pbU "booking-service/internal/user/pbU"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pbU.UserServiceClient, error) {
	grpcAddr := fmt.Sprintf("user-service:%s", cfg.UserPort)
	grpc, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc user client: %s, ", cfg.DoctorPort)
		return nil, err
	}
	log.Printf("Succesfully connected to user client at port: %v", cfg.DoctorPort)
	return pbU.NewUserServiceClient(grpc), nil
}
