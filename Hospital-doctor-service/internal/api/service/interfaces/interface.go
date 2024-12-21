package interfaces

import (
	pb "github.com/anandu86130/Hospital-doctor-service/internal/pbD"
)

type DoctorServiceInter interface {
	SignupService(doctorpb *pb.Signup) (*pb.Response, error)
	VerifyOTP(doctorpb *pb.OTP) (*pb.Response, error)
	LoginService(doctorpb *pb.Login) (*pb.Response, error)
	ViewProfileService(doctorpb *pb.ID) (*pb.DoctorProfile, error)
	EditProfileService(doctorpb *pb.DoctorProfile) (*pb.DoctorProfile, error)
	ChangePassword(doctorpb *pb.Password) (*pb.Response, error)
	BlockDoctorService(doctorpb *pb.ID) (*pb.Response, error)
	UnBlockDoctorService(doctorpb *pb.ID) (*pb.Response, error)
	ViewDoctorList(doctorpb *pb.NoParam) (*pb.DoctorListResponse, error)
	IsVerifiedSVC(doctorpb *pb.ID) (*pb.Response, error)
	AddAvailabilityService(booking *pb.Availability) (*pb.Response, error)
	ViewAvailabilityService(p *pb.NoParam) (*pb.AvailabilityListResponse, error)
	GetUserListService(p *pb.NoParam) (*pb.UserListResponse, error)
	ViewAppointmentService(p *pb.ID) (*pb.AppointmentList, error)
	AddPrescriptionService(p *pb.Prescription) (*pb.Response, error)
	DoctorDetailsService(p *pb.Doctor) (*pb.Doctorresponse, error)
}
