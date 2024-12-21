package service

import (
	"context"

	bookingpb "github.com/anandu86130/Hospital-doctor-service/internal/booking/pbB"
	pb "github.com/anandu86130/Hospital-doctor-service/internal/pbD"
)

func (d *DoctorService) AddAvailabilityService(doctorpb *pb.Availability) (*pb.Response, error) {
	ctx := context.Background()
	booking := &bookingpb.Availability{
		Doctorid:  doctorpb.Doctorid,
		Date:      doctorpb.Date,
		Starttime: doctorpb.Starttime,
		Endtime:   doctorpb.Endtime,
	}
	_, err := d.BookingClient.AddAvailability(ctx, booking)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Failed to add availability",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}
	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Availability Added Successfully",
	}, nil
}

func (d *DoctorService) ViewAvailabilityService(p *pb.NoParam) (*pb.AvailabilityListResponse, error) {
	ctx := context.Background()
	availabilityNoParam := &bookingpb.NoParam{}
	availability, err := d.BookingClient.ViewAvailability(ctx, availabilityNoParam)
	if err != nil {
		return nil, err
	}

	//Map the retrieved user data to the gRPC profile format
	// Prepare the UserListResponse
	var availabilitydetails []*pb.AvailabilityList
	for _, doctoravailability := range availability.Availabilities {
		availabilitydetails = append(availabilitydetails, &pb.AvailabilityList{
			Id:        doctoravailability.Id,
			Doctorid:  doctoravailability.Doctorid,
			Date:      doctoravailability.Date,
			Starttime: doctoravailability.Starttime,
			Endtime:   doctoravailability.Endtime,
		})
	}

	return &pb.AvailabilityListResponse{
		Availabilities: availabilitydetails,
	}, nil
}

func (d *DoctorService) ViewAppointmentService(p *pb.ID) (*pb.AppointmentList, error) {
	ctx := context.Background()
	appointment := &bookingpb.ID{
		ID: p.ID,
	}
	appointmentresponse, err := d.BookingClient.ViewAppointment(ctx, appointment)
	if err != nil {
		return nil, err
	}

	//Map the retrieved user data to the gRPC profile format
	// Prepare the UserListResponse
	var appointmentdetails []*pb.Appointment
	for _, doctorappointment := range appointmentresponse.Profiles {
		appointmentdetails = append(appointmentdetails, &pb.Appointment{
			Id:        doctorappointment.Id,
			Doctorid:  doctorappointment.Doctorid,
			Userid:    doctorappointment.Userid,
			Date:      doctorappointment.Date,
			Starttime: doctorappointment.Starttime,
			Endtime:   doctorappointment.Endtime,
		})
	}

	return &pb.AppointmentList{
		Profiles: appointmentdetails,
	}, nil
}

func (d *DoctorService) AddPrescriptionService(doctorpb *pb.Prescription) (*pb.Response, error) {
	ctx := context.Background()
	Prescription := &bookingpb.Prescription{
		Appointmentid: doctorpb.Appointmentid,
		Doctorid:      doctorpb.Doctorid,
		Userid:        doctorpb.Userid,
		Medicine:      doctorpb.Medicine,
		Notes:         doctorpb.Notes,
	}
	_, err := d.BookingClient.AddPrescription(ctx, Prescription)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Failed to prescription availability",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}
	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "prescription Added Successfully",
	}, nil
}

func (d *DoctorService) DoctorDetailsService(doctorpb *pb.Doctor) (*pb.Doctorresponse, error) {
	response, err := d.Repo.DoctorDetails(uint(doctorpb.Doctorid))
	if err != nil {
		return &pb.Doctorresponse{}, err
	}
	return &pb.Doctorresponse{
		Name: response.Name,
		Fees: response.Fees,
	}, nil
}
