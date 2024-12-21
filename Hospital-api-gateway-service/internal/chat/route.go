package chat

import (
	"log"

	"github.com/anandu86130/Hospital-api-gateway/config"
	pb "github.com/anandu86130/Hospital-api-gateway/internal/chat/pbC"
	"github.com/anandu86130/Hospital-api-gateway/internal/user"
	userpb "github.com/anandu86130/Hospital-api-gateway/internal/user/pbU"
	"github.com/gin-gonic/gin"
)

type Chat struct {
	cfg        *config.Config
	userClient userpb.UserServiceClient
	client     pb.ChatServiceClient
}

func NewChatRoutes(c *gin.Engine, cfg config.Config) {
	client, err := ClientDial(cfg)
	if err != nil {
		log.Fatalf("error not connected with grpc client : %v", err.Error())
	}

	userClient, err := user.ClientDial(cfg)
	if err != nil {
		log.Fatalf("error not connected with grpc user client : %v", err.Error())
	}
	chatHandler := &Chat{
		cfg:        &cfg,
		client:     client,
		userClient: userClient,
	}

	apiVersion := c.Group("/api/v1")

	user := apiVersion.Group("/user")
	{
		user.GET("/chat", chatHandler.Chat)
		user.POST("/video-call", chatHandler.VideoCall)

	}
	c.GET("/chat", chatHandler.ChatPage)
}
