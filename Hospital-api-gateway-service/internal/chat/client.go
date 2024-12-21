package chat

import (
	"fmt"
	"log"

	"github.com/anandu86130/Hospital-api-gateway/config"
	pb "github.com/anandu86130/Hospital-api-gateway/internal/chat/pbC"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.ChatServiceClient, error) {
	grpcAddr := fmt.Sprintf("chat-service:%s", cfg.ChatPort)
	grpc, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing to grpc to client : %s", err.Error())
		return nil, err
	}
	log.Printf("Successfully connected to chat client at port : %s", cfg.ChatPort)
	return pb.NewChatServiceClient(grpc), nil
}
