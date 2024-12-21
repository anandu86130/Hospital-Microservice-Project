package handlers

import (
	"context"

	pb "github.com/anandu86130/Hospital-doctor-service/internal/pbD"
)

func (d *DoctorHandler) AddAvailability(ctx context.Context, doctorpb *pb.Availability) (*pb.Response, error) {
	// Call the service function that handles the unblocking logic
	response, err := d.SVC.AddAvailabilityService(doctorpb)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (d *DoctorHandler) ViewAvailability(ctx context.Context, p *pb.NoParam) (*pb.AvailabilityListResponse, error) {
	response, err := d.SVC.ViewAvailabilityService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (d *DoctorHandler) ViewAppointment(ctx context.Context, p *pb.ID) (*pb.AppointmentList, error) {
	response, err := d.SVC.ViewAppointmentService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (d *DoctorHandler) AddPrescription(ctx context.Context, p *pb.Prescription) (*pb.Response, error) {
	response, err := d.SVC.AddPrescriptionService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (d *DoctorHandler) DoctorDetails(ctx context.Context, p *pb.Doctor) (*pb.Doctorresponse, error) {
	response, err := d.SVC.DoctorDetailsService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
