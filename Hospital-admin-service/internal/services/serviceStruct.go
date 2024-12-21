package services

import (
	bookingpb "github.com/anandu86130/Hospital-admin-service/internal/booking/pbB"
	doctorpb "github.com/anandu86130/Hospital-admin-service/internal/doctor/pbD"
	inter "github.com/anandu86130/Hospital-admin-service/internal/repositories/interfaces"
	"github.com/anandu86130/Hospital-admin-service/internal/services/interfaces"
	userpb "github.com/anandu86130/Hospital-admin-service/internal/user/pbU"
)

type AdminService struct {
	Repo          inter.AdminRepoInter
	UserClient    userpb.UserServiceClient
	DoctorClient  doctorpb.DoctorServiceClient
	BookingClient bookingpb.BookingServiceClient
}

func NewAdminRepository(repo inter.AdminRepoInter, userClient userpb.UserServiceClient, doctorClient doctorpb.DoctorServiceClient, bookingClient bookingpb.BookingServiceClient) interfaces.AdminServiceInter {
	return &AdminService{
		Repo:          repo,
		UserClient:    userClient,
		DoctorClient:  doctorClient,
		BookingClient: bookingClient,
	}
}
