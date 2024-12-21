package booking

import (
	"fmt"
	"log"

	"github.com/anandu86130/Hospital-api-gateway/config"
	pb "github.com/anandu86130/Hospital-api-gateway/internal/booking/pbB"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.BookingServiceClient, error) {
	grpcAddr := fmt.Sprintf("booking-service:%s", cfg.BOOKINGPORT)
	grpc, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing to grpc to client : %s", err.Error())
		return nil, err
	}
	log.Printf("Successfully connected to booking client at port : %s", cfg.BOOKINGPORT)
	return pb.NewBookingServiceClient(grpc), nil
}
