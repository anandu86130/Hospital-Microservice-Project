package interfaces

import (
	pb "github.com/anandu86130/Hospital-admin-service/internal/pb"
)

type AdminServiceInter interface {
	LoginService(p *pb.AdminLogin) (*pb.AdminResponse, error)
	BlockUserService(p *pb.AdID) (*pb.AdminResponse, error)
	UnblockUserService(p *pb.AdID) (*pb.AdminResponse, error)
	BlockDoctorService(p *pb.AdID) (*pb.AdminResponse, error)
	UnblockDoctorService(p *pb.AdID) (*pb.AdminResponse, error)
	IsVerifiedService(p *pb.AdID) (*pb.AdminResponse, error)
	GetUserListService(p *pb.NoParam) (*pb.UserListResponse, error)
	GetDoctorListService(p *pb.NoParam) (*pb.DoctorListResponse, error)
	ViewAllAppointmentService(p *pb.NoParam) (*pb.ViewAppointmentList, error)
}