package handlers

import (
	inter "github.com/anandu86130/Hospital-doctor-service/internal/api/service/interfaces"
	pb "github.com/anandu86130/Hospital-doctor-service/internal/pbD"
)

type DoctorHandler struct{
	SVC inter.DoctorServiceInter
	pb.DoctorServiceServer
}

func NewDoctorHandler(svc inter.DoctorServiceInter) *DoctorHandler {
	return &DoctorHandler{
		SVC: svc,
	}
}
