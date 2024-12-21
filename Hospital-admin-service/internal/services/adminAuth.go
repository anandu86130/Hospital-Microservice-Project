package services

import (
	"errors"
	"log"

	"github.com/anandu86130/Hospital-admin-service/config"
	pb "github.com/anandu86130/Hospital-admin-service/internal/pb"
	"github.com/anandu86130/Hospital-admin-service/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func (a *AdminService) LoginService(p *pb.AdminLogin) (*pb.AdminResponse, error) {
	admin, err := a.Repo.FindAdminByEmail(p.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(p.Password))
	if err != nil {
		log.Printf("Password comparison failed: %v", err)
		return &pb.AdminResponse{
			Status:  pb.AdminResponse_ERROR,
			Message: "password comparison failed",
			Payload: &pb.AdminResponse_Error{Error: err.Error()},
		}, errors.New("password comparison failed")
	}

	// if admin.Password != p.Password {
	// 	return &pb.AdminResponse{
	// 		Status:  pb.AdminResponse_ERROR,
	// 		Message: "Type correct Password",
	// 		Payload: &pb.AdminResponse_Error{Error: "Incorrect Password"},
	// 	}, errors.New("incorrect password")
	// }

	jwtToken, err := utils.GenerateToken(config.LoadConfig().SECERETKEY, admin.Email)
	if err != nil {
		return &pb.AdminResponse{
			Status:  pb.AdminResponse_ERROR,
			Message: "error in generating token",
			Payload: &pb.AdminResponse_Error{Error: err.Error()},
		}, errors.New("error generating token")
	}

	return &pb.AdminResponse{
		Status:  pb.AdminResponse_OK,
		Message: "Login successful",
		Payload: &pb.AdminResponse_Data{Data: jwtToken},
	}, nil
}
