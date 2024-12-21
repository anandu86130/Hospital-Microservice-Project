package handlers

import (
	"context"

	pb "github.com/anandu86130/Hospital-admin-service/internal/pb"
)

func (a *AdminHandler) AdminBlockUser(ctx context.Context, p *pb.AdID) (*pb.AdminResponse, error) {
	response, err := a.SVC.BlockUserService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

// AdminBlockUser helps to find the user and block by Admin service.
func (a *AdminHandler) AdminUnblockUser(ctx context.Context, p *pb.AdID) (*pb.AdminResponse, error) {
	response, err := a.SVC.UnblockUserService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (a *AdminHandler) UserList(ctx context.Context, p *pb.NoParam) (*pb.UserListResponse, error) {
	response, err := a.SVC.GetUserListService(p)
	if err != nil {
		return response, err
	}
	return response, nil	
}
