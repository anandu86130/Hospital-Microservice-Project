package admin

import (
	"log"

	"github.com/anandu86130/Hospital-api-gateway/config"
	pb "github.com/anandu86130/Hospital-api-gateway/internal/admin/pb"
	middleware "github.com/anandu86130/Hospital-api-gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

// Admin represents the admin route handler, containing configuration and gRPC client.
type Admin struct {
	Cfg    *config.Config
	Client pb.AdminServiceClient
}

func NewAdminRoute(c *gin.Engine, cfg config.Config) {
	client, err := ClientDial(cfg)
	if err != nil {
		log.Fatalf("error not connected with grpc client : %v", err.Error())
	}

	adminHandler := &Admin{
		Cfg:    &cfg,
		Client: client,
	}

	apiVersion := c.Group("/api/v1")

	admin := apiVersion.Group("/admin")
	{
		admin.POST("/login", adminHandler.Login)

	}

	auth := admin.Group("/auth")
	auth.Use(middleware.AdminAuthorization(cfg.SECRETKEY, "admin"))
	{
		auth.PATCH("/user/:id", adminHandler.BlockUser)
		auth.PATCH("/user/unblock/:id", adminHandler.UnBlockUser)
		auth.PATCH("/doctor/:id", adminHandler.BlockDoctor)
		auth.PATCH("/doctor/unblock/:id", adminHandler.UnBlockDoctor)
		auth.PATCH("/doctor/verify/:id", adminHandler.IsVerified)
		auth.GET("/user", adminHandler.UserList)
		auth.GET("/doctor", adminHandler.DoctorList)
		auth.GET("/appointment", adminHandler.ViewAllAppointment)
	}
}
