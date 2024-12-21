package user

import (
	"fmt"

	"github.com/anandu86130/Hospital-api-gateway/config"
	middleware "github.com/anandu86130/Hospital-api-gateway/internal/middleware"
	pb "github.com/anandu86130/Hospital-api-gateway/internal/user/pbU"
	"github.com/gin-gonic/gin"
)

type User struct {
	cfg    *config.Config
	client pb.UserServiceClient
}

func NewUserRoute(c *gin.Engine, cfg config.Config) error {
	client, err := ClientDial(cfg)
	if err != nil {
		return fmt.Errorf("error connecting to gRPC client: %v", err)
	}

	userHandler := &User{
		cfg:    &cfg,
		client: client,
	}

	apiVersion := c.Group("/api/v1")
	user := apiVersion.Group("/user")

	// Unauthenticated routes
	{
		user.POST("/signup", userHandler.UserSignup)
		user.POST("/verify", userHandler.VerifyOTP)
		user.POST("/login", userHandler.UserLogin)
	}

	// Authenticated routes
	auth := user.Group("/auth")
	auth.Use(middleware.Authorization(cfg.SECRETKEY))
	{
		auth.GET("/profile", userHandler.ViewProfile)
		auth.PATCH("/profile", userHandler.EditProfile)
		auth.PATCH("/password", userHandler.ChangePassword)
		auth.GET("/doctor", userHandler.ViewDoctor)
		auth.GET("/availability", userHandler.ViewAvailability)
		auth.POST("/appointment", userHandler.BookAppoinment)
		auth.POST("/appointment/view", userHandler.ViewAppointment)
		auth.GET("cancel-appointment", userHandler.CancelAppointment)
		auth.POST("/prescription", userHandler.ViewPrescription)
	}

	// Payment routes
	user.GET("/payment", userHandler.Payment)
	user.POST("/payment/success", userHandler.PaymentSuccess)
	user.GET("/payment-success", userHandler.PaymentSuccessPage)

	return nil
}
