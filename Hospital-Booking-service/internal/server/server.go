package server

import (
	pb "booking-service/internal/booking/pbB"
	"booking-service/internal/handlers"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func NewGrpcBookingServer(port string, handlr *handlers.BookingHandler) error {
	log.Println("Connecting to gRPC server")
	addr := fmt.Sprintf(":%s", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("error creating listener on %v", addr)
		return err
	}
	grpc := grpc.NewServer()
	// pbD.RegisterDoctorServer(grpc, handlr)
	pb.RegisterBookingServiceServer(grpc, handlr)
	log.Printf("listening on gRPC server %v", port)
	err = grpc.Serve(lis)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}
	return nil
}
