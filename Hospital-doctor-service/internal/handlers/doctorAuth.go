package handlers

import (
	"context"

	pb "github.com/anandu86130/Hospital-doctor-service/internal/pbD"
)

func (d *DoctorHandler) DoctorSignup(ctx context.Context, doctorpb *pb.Signup) (*pb.Response, error) {
	response, err := d.SVC.SignupService(doctorpb)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (d *DoctorHandler) VerifyOTP(ctx context.Context, doctorpb *pb.OTP) (*pb.Response, error) {
	response, err := d.SVC.VerifyOTP(doctorpb)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (d *DoctorHandler) DoctorLogin(ctx context.Context, doctorpb *pb.Login) (*pb.Response, error) {
	response, err := d.SVC.LoginService(doctorpb)
	if err != nil {
		return response, err
	}
	return response, nil
}
