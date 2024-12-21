package handlers

import (
	inter "github.com/anandu86130/Hospital-admin-service/internal/services/interfaces"
	pb "github.com/anandu86130/Hospital-admin-service/internal/pb"
)

type AdminHandler struct {
	SVC inter.AdminServiceInter
	pb.AdminServiceServer
}

func NewAdminHandler(svc inter.AdminServiceInter) *AdminHandler {
	return &AdminHandler{
		SVC: svc,
	}
}
