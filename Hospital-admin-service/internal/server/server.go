package server

import (
	"fmt"
	"log"
	"net"

	"github.com/anandu86130/Hospital-admin-service/internal/handlers"
	pb "github.com/anandu86130/Hospital-admin-service/internal/pb"
	"google.golang.org/grpc"
)

func NewGrpcAdminServer(port string, handlr *handlers.AdminHandler) error {
	log.Println("connecting to gRPC server")
	addr := fmt.Sprintf(":%s", port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("error creating listener on %v", addr)
		return err
	}

	grpc := grpc.NewServer()
	pb.RegisterAdminServiceServer(grpc, handlr)

	log.Printf("listening on gRPC server %v", port)
	err = grpc.Serve(lis)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}
	return nil
}
