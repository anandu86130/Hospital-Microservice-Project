package di

import (
	"booking-service/config"
	"booking-service/internal/db"
	"booking-service/internal/doctor"
	"booking-service/internal/handlers"
	"booking-service/internal/repository"
	"booking-service/internal/server"
	"booking-service/internal/service"
	"booking-service/internal/user"
	"booking-service/internal/utils"
	"log"
)

func Init() {
	cfg := config.LoadConfig()

	DB := db.ConnectDB(cfg)

	doctorClient, err := doctor.ClientDial(*cfg)
	if err != nil {
		log.Fatalf("something went wrong when dialing user client %v", err)
	}

	userClient, err := user.ClientDial(*cfg)
	if err != nil {
		log.Fatalf("something went wrong when dialing user client %v", err)
	}

	redis, err := config.SetupRedis(cfg)
	if err != nil {
		log.Fatalf("failed to connect to redis")
	}

	stripeClient := utils.NewStripeClient(*cfg, redis)

	bookingRepo := repository.NewBookingRepository(DB)
	bookingService := service.NewBookingService(bookingRepo, doctorClient, userClient, redis, stripeClient)
	bookingHandler := handlers.NewBookingHandler(bookingService)

	err = server.NewGrpcBookingServer(cfg.BookingPort, bookingHandler)
	if err != nil {
		log.Fatalf("failed to start gRPC server %v", err.Error())
	}
}
