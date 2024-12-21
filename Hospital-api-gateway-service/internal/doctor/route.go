package doctor

import (
	"log"

	"github.com/anandu86130/Hospital-api-gateway/config"
	pb "github.com/anandu86130/Hospital-api-gateway/internal/doctor/pbD"
	"github.com/anandu86130/Hospital-api-gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Doctor struct {
	cfg    *config.Config
	client pb.DoctorServiceClient
}

func NewDoctorRoutes(c *gin.Engine, cfg config.Config) {
	client, err := ClientDial(cfg)
	if err != nil {
		log.Fatalf("error not connected with grpc client : %v", err.Error())
	}

	doctorhandler := &Doctor{
		cfg:    &cfg,
		client: client,
	}

	apiVersion := c.Group("/api/v1")

	user := apiVersion.Group("/doctor")

	{
		user.POST("/signup", doctorhandler.DoctorSignup)
		user.POST("/verify", doctorhandler.VerifyOTP)
		user.POST("/login", doctorhandler.DoctorLogin)
	}
	auth := user.Group("/auth")
	auth.Use(middleware.Authorization(cfg.SECRETKEYDOCTOR))
	{
		auth.GET("/profile", doctorhandler.ViewProfile)
		auth.PATCH("/profile", doctorhandler.EditProfile)
		auth.PATCH("/password", doctorhandler.ChangePassword)
		auth.GET("/user", doctorhandler.UserList)
		auth.GET("/availability", doctorhandler.ViewAvailability)
		auth.POST("/availability", doctorhandler.AddAvailability)
		auth.POST("/appointment", doctorhandler.ViewAppointment)
		auth.POST("/prescription", doctorhandler.AddPrescription)
	}
}
