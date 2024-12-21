package handlers

import (
	pb "booking-service/internal/booking/pbB"
	inter "booking-service/internal/service/interfaces"
	"context"
)

type BookingHandler struct {
	SVC inter.BookingServiceInter
	pb.BookingServiceServer
}

func (d *BookingHandler) AddAvailability(ctx context.Context, pb *pb.Availability) (*pb.Response, error) {
	response, err := d.SVC.AddAvailabilityService(pb)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (b *BookingHandler) ViewAvailability(context.Context, *pb.NoParam) (*pb.AvailabilityListResponse, error) {
	response, err := b.SVC.ViewAvailabilityService()
	if err != nil {
		return response, err
	}
	return response, nil
}

// AddBooking implements __.AvailabilityServiceServer.
func (b *BookingHandler) BookAppoinment(ctx context.Context, appoinmentpb *pb.Appoinment) (*pb.Response, error) {
	response, err := b.SVC.BookAppointmentService(appoinmentpb)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (b *BookingHandler) ViewAppointment(ctx context.Context, id *pb.ID) (*pb.AppointmentList, error) {
	response, err := b.SVC.ViewAppointmentService(id)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (b *BookingHandler) UserViewAppointment(ctx context.Context, id *pb.ID) (*pb.AppointmentList, error) {
	response, err := b.SVC.UserViewAppointmentService(id)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (b *BookingHandler) ViewAllAppointment(context.Context, *pb.NoParam) (*pb.ViewAppointmentList, error) {
	response, err := b.SVC.ViewAllAppointmentService()
	if err != nil {
		return response, err
	}
	return response, nil
}

func (b *BookingHandler) CancelAppointment(ctx context.Context, appoinmentpb *pb.Cancelappointmentreq) (*pb.Response, error) {
	response, err := b.SVC.CancelAppointmentService(appoinmentpb)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (b *BookingHandler) AddPrescription(ctx context.Context, p *pb.Prescription) (*pb.Response, error) {
	response, err := b.SVC.AddPrescriptionService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (b *BookingHandler) ViewPrescription(ctx context.Context, p *pb.Req) (*pb.PrescriptionListResponse, error) {
	response, err := b.SVC.ViewPrescriptionService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (b *BookingHandler) CreatePayment(ctx context.Context, p *pb.ConfirmAppointment) (*pb.PaymentResponse, error) {
	response, err := b.SVC.CreatePaymentService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (m *BookingHandler) UserPaymentSuccess(ctx context.Context, p *pb.Payment) (*pb.PaymentStatusResponse, error) {
	response, err := m.SVC.PaymentSuccessService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func NewBookingHandler(svc inter.BookingServiceInter) *BookingHandler {
	return &BookingHandler{
		SVC: svc,
	}
}
