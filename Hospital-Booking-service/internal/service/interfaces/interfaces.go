package interfaces

import (
	pb "booking-service/internal/booking/pbB"
)

type BookingServiceInter interface {
	AddAvailabilityService(a *pb.Availability) (*pb.Response, error)
	ViewAvailabilityService() (*pb.AvailabilityListResponse, error)
	BookAppointmentService(a *pb.Appoinment) (*pb.Response, error)
	ViewAppointmentService(a *pb.ID) (*pb.AppointmentList, error)
	UserViewAppointmentService(a *pb.ID) (*pb.AppointmentList, error)
	ViewAllAppointmentService() (*pb.ViewAppointmentList, error)
	CancelAppointmentService(a *pb.Cancelappointmentreq) (*pb.Response, error)
	AddPrescriptionService(p *pb.Prescription) (*pb.Response, error)
	ViewPrescriptionService(p *pb.Req) (*pb.PrescriptionListResponse, error)
	CreatePaymentService(p *pb.ConfirmAppointment) (*pb.PaymentResponse, error)
	PaymentSuccessService(p *pb.Payment) (*pb.PaymentStatusResponse, error)
}
