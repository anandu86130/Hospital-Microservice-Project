package handlers

import (
	"context"

	pb "github.com/anandu86130/Hospital-admin-service/internal/pb"
)

func (a *AdminHandler) AdminLoginRequest(ctx context.Context, p *pb.AdminLogin) (*pb.AdminResponse, error) {
	response, err := a.SVC.LoginService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
