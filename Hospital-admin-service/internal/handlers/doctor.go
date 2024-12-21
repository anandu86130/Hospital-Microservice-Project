package handlers

import (
	"context"

	pb "github.com/anandu86130/Hospital-admin-service/internal/pb"
)

func (a *AdminHandler) AdminBlockDoctor(ctx context.Context, p *pb.AdID) (*pb.AdminResponse, error) {
	response, err := a.SVC.BlockDoctorService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// AdminBlockUser helps to find the user and block by Admin service.
func (a *AdminHandler) AdminUnblockDoctor(ctx context.Context, p *pb.AdID) (*pb.AdminResponse, error) {
	response, err := a.SVC.UnblockDoctorService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (a *AdminHandler) AdminIsVerified(ctx context.Context, p *pb.AdID) (*pb.AdminResponse, error) {
	response, err := a.SVC.IsVerifiedService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (a *AdminHandler) DoctorList(ctx context.Context, p *pb.NoParam) (*pb.DoctorListResponse, error) {
	response, err := a.SVC.GetDoctorListService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

