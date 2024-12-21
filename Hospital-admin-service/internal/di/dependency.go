package di

import (
	"log"

	"github.com/anandu86130/Hospital-admin-service/config"
	"github.com/anandu86130/Hospital-admin-service/internal/booking"
	"github.com/anandu86130/Hospital-admin-service/internal/db"
	"github.com/anandu86130/Hospital-admin-service/internal/doctor"
	"github.com/anandu86130/Hospital-admin-service/internal/handlers"
	"github.com/anandu86130/Hospital-admin-service/internal/repositories"
	server "github.com/anandu86130/Hospital-admin-service/internal/server"
	"github.com/anandu86130/Hospital-admin-service/internal/services"
	"github.com/anandu86130/Hospital-admin-service/internal/user"
)

func Init() {
	cfg := config.LoadConfig()
	dbconn := db.ConnectDB(cfg)

	userClient, err := user.ClientDial(*cfg)
	if err != nil {
		log.Fatalf("something went wrong when dialing user client %v", err)
	}

	doctorClient, err := doctor.ClientDial(*cfg)
	if err != nil {
		log.Fatalf("something went wrong when dialing user client %v", err)
	}

	bookingClient, err := booking.ClientDial(*cfg)
	if err != nil {
		log.Fatalf("something went wrong when dialing user client %v", err)
	}

	adminRepo := repositories.NewAdminRepository(dbconn)

	adminService := services.NewAdminRepository(adminRepo, userClient, doctorClient, bookingClient)
	adminHandler := handlers.NewAdminHandler(adminService)
	err = server.NewGrpcAdminServer(cfg.AdminPort, adminHandler)
	if err != nil {
		log.Fatalf("failed to start gRPC server %v", err.Error())
	}
}
