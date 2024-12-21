package booking

import (
	"fmt"
	"log"

	"github.com/anandu86130/Hospital-doctor-service/config"
	pb "github.com/anandu86130/Hospital-doctor-service/internal/booking/pbB"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pb.BookingServiceClient, error) {
	grpcAddr := fmt.Sprintf("booking-service:%s", cfg.GrpcBookingPort)
	grpc, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc booking client: %s, ", cfg.GrpcBookingPort)
		return nil, err
	}
	log.Printf("Succesfully connected to booking client at port: %v", cfg.GrpcBookingPort)
	return pb.NewBookingServiceClient(grpc), nil
}
