package di

import (
	"log"

	"github.com/anandu86130/Hospital-doctor-service/config"
	"github.com/anandu86130/Hospital-doctor-service/internal/api/service"
	"github.com/anandu86130/Hospital-doctor-service/internal/booking"
	"github.com/anandu86130/Hospital-doctor-service/internal/db"
	"github.com/anandu86130/Hospital-doctor-service/internal/handlers"
	"github.com/anandu86130/Hospital-doctor-service/internal/repository"
	"github.com/anandu86130/Hospital-doctor-service/internal/server"
	"github.com/anandu86130/Hospital-doctor-service/internal/user"
)

func Init() {
	cfg := config.LoadConfig()

	DB := db.ConnectDB(cfg)

	redis, err := config.SetupRedis(cfg)
	if err != nil {
		log.Fatalf("failed to connect to redis")
	}

	bookingClient, err := booking.ClientDial(*cfg)
	if err != nil {
		log.Fatalf("something went wrong when dialing booking client %v", err)
	}

	userClient, err := user.ClientDial(*cfg)
	if err != nil {
		log.Fatalf("something went wrong when dialing user client %v", err)
	}

	doctorRepo := repository.NewDoctorRepository(DB)
	doctorService := service.NewDoctorService(doctorRepo, redis, bookingClient, userClient)
	doctorHandler := handlers.NewDoctorHandler(doctorService)

	err = server.NewGrpcDoctorServer(cfg.GrpcDoctorPort, doctorHandler)
	if err != nil {
		log.Fatalf("failed to start gRPC server %v", err.Error())
	}
}
