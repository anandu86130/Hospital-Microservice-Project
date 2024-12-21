package services

import (
	"context"

	doctorpb "github.com/anandu86130/Hospital-admin-service/internal/doctor/pbD"
	pb "github.com/anandu86130/Hospital-admin-service/internal/pb"
)

// BlockUserService handle the admin to block the users using the provided information
func (a *AdminService) BlockDoctorService(p *pb.AdID) (*pb.AdminResponse, error) {
	ctx := context.Background()
	doctor := &doctorpb.ID{
		ID: p.ID,
	}
	_, err := a.DoctorClient.BlockDoctor(ctx, doctor)
	if err != nil {
		return nil, err
	}
	return &pb.AdminResponse{
		Status:  pb.AdminResponse_OK,
		Message: "Doctor blocked successfully",
	}, nil
}

// UnblockUserService implements interfaces.AdminServiceInter.
func (a *AdminService) UnblockDoctorService(p *pb.AdID) (*pb.AdminResponse, error) {
	ctx := context.Background()
	doctor := &doctorpb.ID{
		ID: p.ID,
	}
	_, err := a.DoctorClient.UnblockDoctor(ctx, doctor)
	if err != nil {
		return nil, err
	}
	return &pb.AdminResponse{
		Status:  pb.AdminResponse_OK,
		Message: "Doctor unblocked successfully",
	}, nil
}

func (a *AdminService) IsVerifiedService(p *pb.AdID) (*pb.AdminResponse, error) {
	ctx := context.Background()
	doctor := &doctorpb.ID{
		ID: p.ID,
	}
	_, err := a.DoctorClient.IsVerified(ctx, doctor)
	if err != nil {
		return nil, err
	}
	return &pb.AdminResponse{
		Status:  pb.AdminResponse_OK,
		Message: "Doctor Verified successfully",
	}, nil
}

func (a *AdminService) GetDoctorListService(p *pb.NoParam) (*pb.DoctorListResponse, error) {
	ctx := context.Background()
	doctor := &doctorpb.NoParam{}

	// Call the UserList method from UserClient
	response, err := a.DoctorClient.DoctorList(ctx, doctor)
	if err != nil {
		return nil, err
	}

	// Prepare the UserListResponse
	var profiles []*pb.DoctorProfile
	for _, doctorProfile := range response.Profiles {
		profiles = append(profiles, &pb.DoctorProfile{
			Id:                doctorProfile.Id,
			Name:              doctorProfile.Name,
			Email:             doctorProfile.Email,
			Specialization:    doctorProfile.Specialization,
			YearsOfExperience: doctorProfile.YearsOfExperience,
			LiscenceNumber:    doctorProfile.LiscenceNumber,
			Fees:              doctorProfile.Fees,
		})
	}

	return &pb.DoctorListResponse{
		Profiles: profiles,
	}, nil
}
