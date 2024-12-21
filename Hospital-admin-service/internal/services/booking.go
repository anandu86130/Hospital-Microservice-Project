package services

import (
	"context"

	bookingpb "github.com/anandu86130/Hospital-admin-service/internal/booking/pbB"
	pb "github.com/anandu86130/Hospital-admin-service/internal/pb"
)

func (d *AdminService) ViewAllAppointmentService(p *pb.NoParam) (*pb.ViewAppointmentList, error) {
	ctx := context.Background()
	appointment := &bookingpb.NoParam{}
	appointmentresponse, err := d.BookingClient.ViewAllAppointment(ctx, appointment)
	if err != nil {
		return nil, err
	}

	//Map the retrieved user data to the gRPC profile format
	// Prepare the UserListResponse
	var appointmentdetails []*pb.ViewAppointment
	for _, doctorappointment := range appointmentresponse.Profiles {
		appointmentdetails = append(appointmentdetails, &pb.ViewAppointment{
			Id:            doctorappointment.Id,
			Doctorid:      doctorappointment.Doctorid,
			Userid:        doctorappointment.Userid,
			Date:          doctorappointment.Date,
			Starttime:     doctorappointment.Starttime,
			Endtime:       doctorappointment.Endtime,
			Paymentstatus: doctorappointment.Paymentstatus,
			Amount:        doctorappointment.Amount,
		})
	}

	return &pb.ViewAppointmentList{
		Profiles: appointmentdetails,
	}, nil
}