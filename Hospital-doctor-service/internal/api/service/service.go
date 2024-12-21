package service

import (
	"github.com/anandu86130/Hospital-doctor-service/config"
	"github.com/anandu86130/Hospital-doctor-service/internal/api/service/interfaces"
	bookingpb "github.com/anandu86130/Hospital-doctor-service/internal/booking/pbB"
	inter "github.com/anandu86130/Hospital-doctor-service/internal/repository/interfaces"
	userpb "github.com/anandu86130/Hospital-doctor-service/internal/user/pbU"
)

type DoctorService struct {
	Repo          inter.DoctorRepository
	Redis         *config.RedisService
	BookingClient bookingpb.BookingServiceClient
	UserClient    userpb.UserServiceClient
}

func NewDoctorService(repo inter.DoctorRepository, redis *config.RedisService, bookingClient bookingpb.BookingServiceClient, userClient userpb.UserServiceClient) interfaces.DoctorServiceInter {
	return &DoctorService{
		Repo:          repo,
		Redis:         redis,
		BookingClient: bookingClient,
		UserClient:    userClient,
	}
}
