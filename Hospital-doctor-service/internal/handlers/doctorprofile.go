package handlers

import (
	"context"

	pb "github.com/anandu86130/Hospital-doctor-service/internal/pbD"
)

// ViewProfile fetch the user profile from database
func (d *DoctorHandler) ViewProfile(ctx context.Context, doctorpb *pb.ID) (*pb.DoctorProfile, error) {
	response, err := d.SVC.ViewProfileService(doctorpb)
	if err != nil {
		return response, err
	}
	return response, nil
}

// EditProfile update the user profile in database
func (d *DoctorHandler) EditProfile(ctx context.Context, doctorpb *pb.DoctorProfile) (*pb.DoctorProfile, error) {
	response, err := d.SVC.EditProfileService(doctorpb)
	if err != nil {
		return response, err
	}
	return response, nil
}

// ChangePassword update the user profile in database
func (d *DoctorHandler) ChangePassword(ctx context.Context, doctorpb *pb.Password) (*pb.Response, error) {
	response, err := d.SVC.ChangePassword(doctorpb)
	if err != nil {
		return response, err
	}
	return response, nil
}

// BlockUser update the user as blocked in database
func (d *DoctorHandler) BlockDoctor(ctx context.Context, doctorpb *pb.ID) (*pb.Response, error) {
	response, err := d.SVC.BlockDoctorService(doctorpb)
	if err != nil {
		return response, err
	}
	return response, nil
}

// BlockUser update the user as blocked in database
func (d *DoctorHandler) IsVerified(ctx context.Context, doctorpb *pb.ID) (*pb.Response, error) {
	response, err := d.SVC.IsVerifiedSVC(doctorpb)
	if err != nil {
		return response, err
	}
	return response, nil
}

// UnblockUser updates the user as unblocked in the database
func (d *DoctorHandler) UnblockDoctor(ctx context.Context, doctorpb *pb.ID) (*pb.Response, error) {
	// Call the service function that handles the unblocking logic
	response, err := d.SVC.UnBlockDoctorService(doctorpb)
	if err != nil {
		return response, err
	}
	return response, nil
}

// UserList fetech the userdata from database
func (d *DoctorHandler) DoctorList(ctx context.Context, doctorpb *pb.NoParam) (*pb.DoctorListResponse, error) {
	response, err := d.SVC.ViewDoctorList(doctorpb)
	if err != nil {
		return response, err
	}
	return response, nil
}
