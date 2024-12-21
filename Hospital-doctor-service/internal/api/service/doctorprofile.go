package service

import (
	"errors"
	"fmt"

	"github.com/anandu86130/Hospital-doctor-service/internal/password"
	pb "github.com/anandu86130/Hospital-doctor-service/internal/pbD"
)

func (d *DoctorService) ViewProfileService(doctorpb *pb.ID) (*pb.DoctorProfile, error) {
	doctor, err := d.Repo.FindDoctorByID(uint(doctorpb.ID))
	if err != nil {
		return nil, err
	}
	doctorprofile := &pb.DoctorProfile{
		Id:                doctorpb.ID,
		Name:              doctor.Name,
		Email:             doctor.Email,
		Specialization:    doctor.Specialization,
		YearsOfExperience: uint32(doctor.YearsOfExperience),
		LiscenceNumber:    doctor.LicenceNumber,
		Fees:              uint32(doctor.Fees),
	}
	return doctorprofile, nil
}

func (d *DoctorService) EditProfileService(doctorpb *pb.DoctorProfile) (*pb.DoctorProfile, error) {
	// Find the doctor by ID
	doctor, err := d.Repo.FindDoctorByID(uint(doctorpb.Id))
	if err != nil {
		return nil, err
	}

	// Ensure fees are valid (at least 100)
	if doctorpb.Fees < 100 {
		return &pb.DoctorProfile{}, errors.New("fees should be at least 100RS")
	}

	// Ensure name is at least 4 characters
	if len(doctorpb.Name) < 4 {
		return &pb.DoctorProfile{}, errors.New("name should be at least 4 characters long")
	}

	// Update only the allowed fields
	doctor.Name = doctorpb.Name
	doctor.YearsOfExperience = uint32(doctorpb.YearsOfExperience)
	doctor.Fees = uint32(doctorpb.Fees)

	// Call the repository to update the doctor in the database
	err = d.Repo.UpdateDoctor(doctor)
	if err != nil {
		return nil, err
	}

	// Return the updated doctor profile
	return doctorpb, nil
}

func (d *DoctorService) BlockDoctorService(doctorpb *pb.ID) (*pb.Response, error) {
	doctor, err := d.Repo.FindDoctorByID(uint(doctorpb.ID))
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in fetching doctor from database",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}
	doctor.IsBlocked = true
	err = d.Repo.UpdateDoctor(doctor)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in updating doctor",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}
	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "doctor blocked successfully",
	}, nil
}

func (d *DoctorService) UnBlockDoctorService(doctorpb *pb.ID) (*pb.Response, error) {
	doctor, err := d.Repo.FindDoctorByID(uint(doctorpb.ID))
	fmt.Println("user data", doctor)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in getting doctor from database",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}
	doctor.IsBlocked = false
	fmt.Println("user", doctor.IsBlocked)
	err = d.Repo.UpdateDoctor(doctor)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in updating doctor",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "doctor unblocked successfully",
	}, nil
}

func (d *DoctorService) ChangePassword(doctorpb *pb.Password) (*pb.Response, error) {
	doctor, err := d.Repo.FindDoctorByID(uint(doctorpb.User_ID))
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error when fetching user from database",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}
	fmt.Println(doctorpb.Old_Password, doctor.Password)
	if !password.CheckPassword(doctorpb.Old_Password, doctor.Password) {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "old password is incorrect",
		}, errors.New("old password mismatch")
	}

	if doctorpb.New_Password != doctorpb.Confirm_Password {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "password is incorrect, passwords not matching",
		}, errors.New("new password mismatch")
	}

	newPassword, err := password.HashPassword(doctorpb.New_Password)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Error while hashing new password",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	doctor.Password = newPassword

	err = d.Repo.UpdateDoctor(doctor)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Error while updating password",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Password changed successfully",
	}, nil
}

func (d *DoctorService) ViewDoctorList(p *pb.NoParam) (*pb.DoctorListResponse, error) {
	// Fetch the user list from the repository to interact with db
	doctors, err := d.Repo.GetDoctorList()
	if err != nil {
		return nil, err
	}

	//Create a response message to hold the user profiles
	doctorListResponse := &pb.DoctorListResponse{}

	//Map the retrieved user data to the gRPC profile format
	for _, doctor := range doctors {
		// Construct each Profile from the user data
		profile := &pb.DoctorProfile{
			Id:                uint32(doctor.ID),
			Name:              doctor.Name,
			Email:             doctor.Email,
			Specialization:    doctor.Specialization,
			YearsOfExperience: uint32(doctor.YearsOfExperience),
			LiscenceNumber:    doctor.LicenceNumber,
			Fees:              uint32(doctor.Fees),
		}

		// Append each Profile to the UserListResponse
		doctorListResponse.Profiles = append(doctorListResponse.Profiles, profile)
	}

	//  Return the populated UserListResponse and nil error
	return doctorListResponse, nil
}

func (d *DoctorService) IsVerifiedSVC(doctorpb *pb.ID) (*pb.Response, error) {
	doctor, err := d.Repo.FindDoctorByID(uint(doctorpb.ID))
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in fetching doctor from database",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}
	doctor.IsVerified = true
	err = d.Repo.UpdateDoctor(doctor)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in updating doctor",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}
	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "doctor verified successfully",
	}, nil
}
