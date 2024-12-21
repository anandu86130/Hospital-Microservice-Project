package handlers

import (
	"context"
	pb "github.com/anandu86130/Hospital-doctor-service/internal/pbD"
)

func (a *DoctorHandler) UserList(ctx context.Context, p *pb.NoParam) (*pb.UserListResponse, error) {
	response, err := a.SVC.GetUserListService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
