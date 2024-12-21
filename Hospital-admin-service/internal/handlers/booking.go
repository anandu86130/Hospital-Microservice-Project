package handlers

import (
	"context"

	pb "github.com/anandu86130/Hospital-admin-service/internal/pb"
)

func (d *AdminHandler) ViewAllAppointment(ctx context.Context, p *pb.NoParam) (*pb.ViewAppointmentList, error) {
	response, err := d.SVC.ViewAllAppointmentService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
