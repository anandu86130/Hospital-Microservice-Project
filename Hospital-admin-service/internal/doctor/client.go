package doctor

import (
	"fmt"
	"log"

	"github.com/anandu86130/Hospital-admin-service/config"
	pbD "github.com/anandu86130/Hospital-admin-service/internal/doctor/pbD"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ClientDial(cfg config.Config) (pbD.DoctorServiceClient, error) {
	grpcAddr := fmt.Sprintf("doctor-service:%s", cfg.DoctorPort)
	grpc, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error Dialing to grpc doctor client: %s, ", cfg.DoctorPort)
		return nil, err
	}
	log.Printf("Succesfully connected to doctor client at port: %v", cfg.DoctorPort)
	return pbD.NewDoctorServiceClient(grpc), nil
}
