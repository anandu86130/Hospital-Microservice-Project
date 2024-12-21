package admin

import (
	"fmt"
	"log"

	"github.com/anandu86130/Hospital-api-gateway/config"
	pb "github.com/anandu86130/Hospital-api-gateway/internal/admin/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.AdminServiceClient, error) {
	grpcAddr := fmt.Sprintf("admin-service:%s", cfg.ADMINPORT)
	grpc, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing to grpc to client : %s", err.Error())
		return nil, err
	}
	log.Printf("Successfully connected to admin client at port : %s", cfg.ADMINPORT)
	return pb.NewAdminServiceClient(grpc), nil
}
