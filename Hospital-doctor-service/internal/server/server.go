package server

import (
	"fmt"
	"log"
	"net"

	"github.com/anandu86130/Hospital-doctor-service/internal/handlers"
	pbD "github.com/anandu86130/Hospital-doctor-service/internal/pbD"
	"google.golang.org/grpc"
)

func NewGrpcDoctorServer(port string, handlr *handlers.DoctorHandler) error {
	log.Println("Connecting to gRPC server")
	addr := fmt.Sprintf(":%s", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("error creating listener on %v", addr)
		return err
	}
	grpc := grpc.NewServer()
	// pbD.RegisterDoctorServer(grpc, handlr)
	pbD.RegisterDoctorServiceServer(grpc, handlr)
	log.Printf("listening on gRPC server %v", port)
	err = grpc.Serve(lis)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}
	return nil
}
