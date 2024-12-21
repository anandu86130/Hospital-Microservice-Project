package doctor

import (
	"fmt"
	"log"

	"github.com/anandu86130/Hospital-api-gateway/config"
	pbD "github.com/anandu86130/Hospital-api-gateway/internal/doctor/pbD"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pbD.DoctorServiceClient, error) {
	grpcAddr := fmt.Sprintf("doctor-service:%s", cfg.DOCTORPORT)
	grpc, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing to grpc to client : %s", err.Error())
		return nil, err
	}
	log.Printf("Successfully connected to docotor client at port : %s", cfg.DOCTORPORT)
	return pbD.NewDoctorServiceClient(grpc), nil
}
