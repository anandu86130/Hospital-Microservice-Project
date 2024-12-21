package service

import (
	"booking-service/config"
	pb "booking-service/internal/booking/pbB"
	doctorpb "booking-service/internal/doctor/pbD"
	"booking-service/internal/model"
	inter "booking-service/internal/repository/interfaces"
	interfaces "booking-service/internal/service/interfaces"
	userpb "booking-service/internal/user/pbU"
	"booking-service/internal/utility"
	"booking-service/internal/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookingService struct {
	Repo         inter.BookingRepository
	DoctorClient doctorpb.DoctorServiceClient
	UserClient   userpb.UserServiceClient
	Redis        *config.RedisService
	StripePay    *utils.StripeClient
}

func NewBookingService(repo inter.BookingRepository, doctorclient doctorpb.DoctorServiceClient, userclient userpb.UserServiceClient, redis *config.RedisService, stripepay *utils.StripeClient) interfaces.BookingServiceInter {
	return &BookingService{
		Repo:         repo,
		DoctorClient: doctorclient,
		UserClient:   userclient,
		Redis:        redis,
		StripePay:    stripepay,
	}
}

func (b *BookingService) AddAvailabilityService(availabilitypb *pb.Availability) (*pb.Response, error) {
	log.Println("Checking for existing availability...")

	// Define startDate from the provided date in availabilitypb
	startDate := time.Unix(availabilitypb.Date.Seconds, 0).Truncate(24 * time.Hour)
	starttime := time.Unix(availabilitypb.Starttime.Seconds, 0)
	endtime := time.Unix(availabilitypb.Endtime.Seconds, 0)

	// Check if the provided date is in the future
	currentDate := time.Now().Truncate(24 * time.Hour)
	if !startDate.After(currentDate) {
		log.Printf("Provided date %v is today or in the past", startDate)
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Availability for this date cannot be provided. Please choose another date.",
			Payload: &pb.Response_Error{Error: "The selected date is today or in the past"},
		}, errors.New("the selected date is today or in the past, please select the date from tommorrow onwards")
	}

	// Check for existing availability
	availability, err := b.Repo.CheckAvailabilityByDate(availabilitypb.Doctorid, startDate, starttime, endtime)
	if err != nil {
		log.Printf("Error checking availability: %v", err)
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Failed to check availability",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	if availability != nil {
		log.Printf("Availability already exists for DoctorID %d on Date %v", availabilitypb.Doctorid, startDate)
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "This availability already exists",
			Payload: &pb.Response_Error{Error: "Availability already exists for the selected date and time range"},
		}, errors.New("availability already exists")
	}

	log.Println("No existing availability, proceeding to create a new one...")

	// Create new availability
	availabilityProfile := &model.Availability{
		Doctorid:  availabilitypb.Doctorid,
		Date:      startDate,
		StartTime: starttime,
		EndTime:   endtime,
	}

	err = b.Repo.CreateAvailability(availabilityProfile)
	if err != nil {
		log.Printf("Error saving availability: %v", err)
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Failed to create availability",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	log.Println("Availability added successfully.")
	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Availability added successfully",
	}, nil
}

func (b *BookingService) ViewAvailabilityService() (*pb.AvailabilityListResponse, error) {
	log.Println("Fetching all availabilities...")

	// Fetch all availability records
	availabilityList, err := b.Repo.GetAllAvailabilities()
	if err != nil {
		log.Printf("Error fetching availability: %v", err)
		return &pb.AvailabilityListResponse{
			Availabilities: nil,
		}, err
	}

	// Check if any availability was found
	if len(availabilityList) == 0 {
		log.Println("No availability records found")
		return &pb.AvailabilityListResponse{
			Availabilities: nil,
		}, nil
	}

	// Convert the availability list to the response format
	var availabilityItems []*pb.AvailabilityList
	for _, avail := range availabilityList {
		availabilityItems = append(availabilityItems, &pb.AvailabilityList{
			Id:        uint32(avail.ID),
			Doctorid:  avail.Doctorid,
			Date:      &timestamppb.Timestamp{Seconds: avail.Date.Unix()},
			Starttime: &timestamppb.Timestamp{Seconds: avail.StartTime.Unix()},
			Endtime:   &timestamppb.Timestamp{Seconds: avail.EndTime.Unix()},
		})
	}

	log.Println("Availability fetched successfully.")
	return &pb.AvailabilityListResponse{
		Availabilities: availabilityItems,
	}, nil
}

func (b *BookingService) BookAppointmentService(appointmentpb *pb.Appoinment) (*pb.Response, error) {
	log.Println("Attempting to book an appointment...")

	// Convert protobuf timestamps to Go time.Time
	startDate := time.Unix(appointmentpb.Date.Seconds, 0).Truncate(24 * time.Hour)
	starttime := time.Unix(appointmentpb.Starttime.Seconds, 0)
	endtime := time.Unix(appointmentpb.Endtime.Seconds, 0)

	// Check if the time difference is less than 15 minutes
	timeDifference := endtime.Sub(starttime)
	if timeDifference < 15*time.Minute {
		log.Printf("Error: Selected time slot is less than 15 minutes. Minimum time period is 15 minutes.")
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Selected time slot is less than 15 minutes. Please select a time period of at least 15 minutes.",
		}, errors.New("selected time slot is less than 15 minutes. Please select a time period of at least 15 minutes")
	}

	// Check for existing availability for the requested time
	availability, err := b.Repo.CheckAvailabilityByDate(appointmentpb.Doctorid, startDate, starttime, endtime)
	if err != nil {
		log.Printf("Error checking availability: %v", err)
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Failed to check availability",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	if availability == nil {
		log.Println("No availability found for the requested date and time.")
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "No availability for the selected time slot",
			Payload: &pb.Response_Error{Error: "No availability found"},
		}, errors.New("no availability for the selected time slot")
	}

	// Check if an appointment already exists in this slot
	existingAppointments, err := b.Repo.GetAppointmentsByDoctorAndDate(appointmentpb.Doctorid, startDate, starttime, endtime)
	if err != nil {
		log.Printf("Error retrieving existing appointments: %v", err)
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Failed to retrieve existing appointments",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	for _, appt := range existingAppointments {
		if appt.PaymentStatus == "Completed" {
			log.Println("Conflict detected with another appointment in this time slot.")
			return &pb.Response{
				Status:  pb.Response_ERROR,
				Message: "Time slot is already booked",
				Payload: &pb.Response_Error{Error: "Time slot is already booked"},
			}, errors.New("time slot is already booked")
		}
	}
	doctorDetails := &doctorpb.Doctor{
		Doctorid: appointmentpb.Doctorid, // Convert uint32 into a Doctor struct.
	}

	ctx := context.Background()
	// Now call the method with the correct type
	response, err := b.DoctorClient.DoctorDetails(ctx, doctorDetails)
	if err != nil {
		log.Fatalf("Failed to fetch doctor details: %v", err)
	}

	// Book the appointment
	newAppointment := &model.Appoinment{
		Doctorid:         appointmentpb.Doctorid,
		Userid:           appointmentpb.Userid,
		Date:             startDate,
		StartTime:        starttime,
		EndTime:          endtime,
		Fees:             response.Fees,
		Doctorname:       response.Name,
		PaymentStatus:    "pending",
		AppoinmentStatus: "pending",
	}

	appointmentID, err := b.Repo.CreateAppointment(newAppointment) // Assuming this returns the ID
	if err != nil {
		log.Printf("Error saving appointment: %v", err)
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Failed to create appointment",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	log.Println("Please complete the payment.")

	// Build the response string manually
	dataStr := fmt.Sprintf("appointment_id:%d doctor_name:%s fees:%d", appointmentID, response.Name, response.Fees)

	// Prepare appointment details for Redis storage
	AppointmentData := map[string]interface{}{
		"Appointment":   appointmentID, // Use the generated AppointmentID
		"UserID":        newAppointment.Userid,
		"PaymentAmount": newAppointment.Fees,
		"Status":        newAppointment.PaymentStatus,
	}

	// Serialize the map to JSON

	dataJSON, err := json.Marshal(AppointmentData)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Failed to serialize data",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}
	log.Println("appointment data created================", AppointmentData)

	// Define Redis key for order and store serialized data
	key := fmt.Sprintf("appointment:%d:user:%d", appointmentID, newAppointment.Userid)
	log.Println("marshedlled appointment data========================", dataJSON)
	log.Println("==========================================key is==================", key)

	// Store order details in Redis
	err = b.Redis.SetDataInRedis(key, dataJSON, time.Hour) // setting expiration to 1 hour
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Failed to store order data in Redis",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	return &pb.Response{
		Status: pb.Response_OK,
		Payload: &pb.Response_Data{
			Data: dataStr, // Here, we use the string data directly
		},
	}, nil
}

func (b *BookingService) ViewAppointmentService(id *pb.ID) (*pb.AppointmentList, error) {
	log.Println("Fetching all appointments...")

	// Fetch all availability records
	appointment, err := b.Repo.GetDoctorAppointment(id.ID, "Completed")
	if err != nil {
		log.Printf("Error fetching appointment: %v", err)
		return &pb.AppointmentList{
			Profiles: nil,
		}, err
	}

	// Convert the availability list to the response format
	var appointmentdetails []*pb.Appointment
	for _, avail := range appointment {
		appointmentdetails = append(appointmentdetails, &pb.Appointment{
			Id:        uint32(avail.ID),
			Userid:    avail.Userid,
			Doctorid:  avail.Doctorid,
			Date:      &timestamppb.Timestamp{Seconds: avail.Date.Unix()},
			Starttime: &timestamppb.Timestamp{Seconds: avail.StartTime.Unix()},
			Endtime:   &timestamppb.Timestamp{Seconds: avail.EndTime.Unix()},
		})
	}

	log.Println("appointment fetched successfully.")
	return &pb.AppointmentList{
		Profiles: appointmentdetails,
	}, nil
}

func (b *BookingService) UserViewAppointmentService(id *pb.ID) (*pb.AppointmentList, error) {
	log.Println("Fetching all appointments...")

	// Fetch all availability records
	appointment, err := b.Repo.GetDoctorUserAppointment(id.ID, "Completed")
	if err != nil {
		log.Printf("Error fetching appointment: %v", err)
		return &pb.AppointmentList{
			Profiles: nil,
		}, err
	}

	// Convert the availability list to the response format
	var appointmentdetails []*pb.Appointment
	for _, avail := range appointment {
		appointmentdetails = append(appointmentdetails, &pb.Appointment{
			Id:        uint32(avail.ID),
			Userid:    avail.Userid,
			Doctorid:  avail.Doctorid,
			Date:      &timestamppb.Timestamp{Seconds: avail.Date.Unix()},
			Starttime: &timestamppb.Timestamp{Seconds: avail.StartTime.Unix()},
			Endtime:   &timestamppb.Timestamp{Seconds: avail.EndTime.Unix()},
		})
	}

	log.Println("appointment fetched successfully.")
	return &pb.AppointmentList{
		Profiles: appointmentdetails,
	}, nil
}

func (b *BookingService) ViewAllAppointmentService() (*pb.ViewAppointmentList, error) {
	log.Println("Fetching all availabilities...")

	// Fetch all availability records
	appointment, err := b.Repo.GetAllAppointment()
	if err != nil {
		log.Printf("Error fetching appointment: %v", err)
		return &pb.ViewAppointmentList{
			Profiles: nil,
		}, err
	}

	// Convert the availability list to the response format
	var appointmentdetails []*pb.ViewAppointment
	for _, avail := range appointment {
		appointmentdetails = append(appointmentdetails, &pb.ViewAppointment{
			Id:            uint32(avail.ID),
			Userid:        avail.Userid,
			Doctorid:      avail.Doctorid,
			Date:          &timestamppb.Timestamp{Seconds: avail.Date.Unix()},
			Starttime:     &timestamppb.Timestamp{Seconds: avail.StartTime.Unix()},
			Endtime:       &timestamppb.Timestamp{Seconds: avail.EndTime.Unix()},
			Amount:        avail.Fees,
			Paymentstatus: avail.PaymentStatus,
		})
	}

	log.Println("appointment fetched successfully.")
	return &pb.ViewAppointmentList{
		Profiles: appointmentdetails,
	}, nil
}

func (b *BookingService) AddPrescriptionService(Prescriptionpb *pb.Prescription) (*pb.Response, error) {
	err := b.Repo.CheckAppointment(Prescriptionpb.Appointmentid, Prescriptionpb.Doctorid, Prescriptionpb.Userid)
	if err != nil {
		log.Printf("no appointment exist for this uesr: %v", err)
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "no appointment exists for this user",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	Prescription, err := b.Repo.CheckOrCreatePrescription(Prescriptionpb.Appointmentid, Prescriptionpb.Doctorid, Prescriptionpb.Userid, Prescriptionpb.Medicine, Prescriptionpb.Notes)
	if err != nil {
		log.Printf("Error checking prescription: %v", err)
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "Failed to check prescription",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, err
	}

	if Prescription == nil {
		log.Printf("prescription already exists for this user %d", Prescriptionpb.Doctorid)
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "This prescription already exists",
			Payload: &pb.Response_Error{Error: "prescription already exists for this user"},
		}, errors.New("prescription already exists")
	}

	log.Println("Prescription added successfully.")
	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Prescription added successfully",
	}, nil
}

func (b *BookingService) ViewPrescriptionService(id *pb.Req) (*pb.PrescriptionListResponse, error) {
	log.Println("Fetching all prescription...")

	// Fetch all availability records
	prescriptions, err := b.Repo.ViewPrescription(id.Userid, id.Appoinmentid)
	if err != nil {
		log.Printf("Error fetching prescriptions: %v", err)
		return &pb.PrescriptionListResponse{
			Profiles: nil,
		}, err
	}

	// Convert the prescription list to the response format
	var prescriptionDetails []*pb.PrescriptionList
	for _, avail := range prescriptions {
		// Initialize timestamp for date, start, and end time
		dateTimestamp := &timestamppb.Timestamp{Seconds: avail.Appoinment.Date.Unix()}
		startTimestamp := &timestamppb.Timestamp{}
		endTimestamp := &timestamppb.Timestamp{}

		// Only set start and end time if they're not zero
		if !avail.Appoinment.StartTime.IsZero() {
			startTimestamp = &timestamppb.Timestamp{Seconds: avail.Appoinment.StartTime.Unix()}
		}
		if !avail.Appoinment.EndTime.IsZero() {
			endTimestamp = &timestamppb.Timestamp{Seconds: avail.Appoinment.EndTime.Unix()}
		}
		prescriptionDetails = append(prescriptionDetails, &pb.PrescriptionList{
			Id:            uint32(avail.ID),
			Appointmentid: avail.Appoinmentid,
			Userid:        avail.Userid,
			Doctorid:      avail.Doctorid,
			Medicine:      avail.Medicine,
			Notes:         avail.Notes,
			Date:          dateTimestamp,
			Starttime:     startTimestamp,
			Endtime:       endTimestamp,
		})
	}
	return &pb.PrescriptionListResponse{
		Profiles: prescriptionDetails,
	}, nil
}

func (b *BookingService) CreatePaymentService(p *pb.ConfirmAppointment) (*pb.PaymentResponse, error) {
	key := fmt.Sprintf("appointment:%d:user:%d", p.Appoinmentid, p.Userid)
	log.Printf("Attempting to retrieve appointment data with key: %s", key)

	// Retrieve order data from Redis
	appointmentdata, err := b.Redis.GetFromRedis(key)
	if err != nil {
		if err == redis.Nil {
			log.Printf("No data found in Redis for key: %s", key)
			return nil, fmt.Errorf("no appointment data found for appointment ID %v and user ID %v", p.Appoinmentid, p.Userid)
		}
		return nil, err
	}

	log.Printf("Appointment data retrieved from Redis for key %s: %s", key, appointmentdata)

	var payment model.Payment
	err = json.Unmarshal([]byte(appointmentdata), &payment)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal appointment data: %v", err)
	}

	// Check if a PaymentIntent already exists and if it's incomplete
	if payment.PaymentID != "" && payment.Status != "Completed" {
		log.Printf("Existing PaymentIntent found with ID: %s", payment.PaymentID)
	} else {
		// If no PaymentIntent exists or if the status is "Completed," create a new one
		amountInCents := int(payment.PaymentAmount)
		if amountInCents < 100 { // Convert to cents if necessary
			amountInCents = int(payment.PaymentAmount * 100)
		}

		// Create the PaymentIntent with Stripe
		paymtID, clientSecret, err := b.StripePay.CreatePaymentIntent(int64(amountInCents), "usd")
		if err != nil {
			return nil, fmt.Errorf("failed to create payment intent: %v", err)
		}
		log.Printf("Payment Intent created with ID: %s", paymtID)

		// Update payment details
		payment.PaymentID = paymtID
		payment.Appoinmentid = uint(p.Appoinmentid)
		payment.ClientSecret = clientSecret
		payment.PaymentMethod = "stripe"
		payment.Status = "Pending" // Set status to "Pending" for the new payment intent

		// Update the order data in Redis
		updatedPaymentData, err := json.Marshal(payment)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal updated payment data: %v", err)
		}
		err = b.Redis.SetDataInRedis(key, updatedPaymentData, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to update payment data in Redis: %v", err)
		}

		// Save the new payment intent details in the database
		err = b.Repo.SavePayment(&payment)
		if err != nil {
			return nil, fmt.Errorf("failed to save payment data in the database: %v", err)
		}
	}
	log.Println("==============================", payment)

	// Return the payment response using the existing or newly created PaymentIntent
	response := &pb.PaymentResponse{
		Paymentid:    payment.PaymentID,
		Clientsecret: payment.ClientSecret,
		Appoinmentid: p.Appoinmentid,
		Amount:       uint32(payment.PaymentAmount),
	}
	return response, nil
}

func (b *BookingService) PaymentSuccessService(p *pb.Payment) (*pb.PaymentStatusResponse, error) {
	key := fmt.Sprintf("appointment:%d:user:%d", p.AppointmentID, p.UserID)
	log.Printf("Retrieving payment data from Redis for key: %s", key)

	// Step 1: Fetch payment data from Redis
	paymentData, err := b.Redis.GetFromRedis(key)
	var payment model.Payment
	if err == redis.Nil {
		log.Printf("No payment data / payment paid already; fetching from database for appointment_id: %d", p.AppointmentID)

		// Fetch the latest payment from the database
		payment, err = b.Repo.GetLatestPaymentByAppointmentID(int(p.AppointmentID))
		if err != nil {
			return &pb.PaymentStatusResponse{
				Status:  pb.PaymentStatusResponse_FAILED,
				Message: fmt.Sprintf("Failed to fetch payment data from the database: %v", err),
			}, nil
		}
	} else if err != nil {
		return &pb.PaymentStatusResponse{
			Status:  pb.PaymentStatusResponse_FAILED,
			Message: fmt.Sprintf("Error fetching payment data from Redis: %v", err),
		}, nil
	} else {
		json.Unmarshal([]byte(paymentData), &payment)
	}

	// Step 2: Check if payment has already been completed
	if payment.Status == "Completed" {
		log.Printf("Payment already completed for payment_id: %s", payment.PaymentID)
		return &pb.PaymentStatusResponse{
			Status:  pb.PaymentStatusResponse_SUCCESS,
			Message: "Payment has already been successfully processed.",
		}, nil
	}

	// Step 3: Verify payment status with Stripe
	paymentStatus, err := b.StripePay.VerifyPaymentStatus(payment.PaymentID)
	if err != nil {
		return &pb.PaymentStatusResponse{
			Status:  pb.PaymentStatusResponse_FAILED,
			Message: fmt.Sprintf("Failed to verify payment status: %v", err),
		}, nil
	}

	// Step 4: If payment is not successful, return failure
	if paymentStatus != "succeeded" {
		return &pb.PaymentStatusResponse{
			Status:  pb.PaymentStatusResponse_FAILED,
			Message: "Payment failed or not completed.",
		}, nil
	}

	// Step 5: Check if the order is already marked as completed before updating
	appointments, err := b.Repo.FindAppointmentsByID(uint(p.AppointmentID))
	if err != nil {
		return &pb.PaymentStatusResponse{
			Status:  pb.PaymentStatusResponse_FAILED,
			Message: fmt.Sprintf("Error fetching appointment data from the database: %v", err),
		}, nil
	}

	// Prevent updating the payment if the order is already completed
	if appointments.PaymentStatus == "Completed" {
		log.Printf("Appointment already completed, skipping payment update for appointment_id: %d", p.AppointmentID)
		return &pb.PaymentStatusResponse{
			Status:  pb.PaymentStatusResponse_SUCCESS,
			Message: "Appointment has already been completed and paid.",
		}, nil
	}

	// Step 6: Update payment and order status in the database
	payment.Status = "Completed"
	err = b.Repo.UpdatePaymentAndAppointmentStatus(payment.PaymentID, int(p.AppointmentID), payment.Status, "Completed", "Completed")
	if err != nil {
		return &pb.PaymentStatusResponse{
			Status:  pb.PaymentStatusResponse_FAILED,
			Message: fmt.Sprintf("Failed to update payment and appointment status: %v", err),
		}, nil
	}

	// type appointmentdetails struct {
	// 	Doctorid      uint32
	// 	Userid        uint32
	// 	Date          time.Time
	// 	StartTime     time.Time
	// 	EndTime       time.Time
	// 	PaymentStatus string
	// 	Doctorname    string
	// 	Fees          uint32
	// }
	var Appointmentpayload []utility.AppointmentPayload

	// Create a single AppointmentPayload
	appointmentpayloads := []utility.AppointmentPayload{
		{
			AppointmentID: p.AppointmentID,
			Doctorid:      appointments.Doctorid,
			Userid:        appointments.Userid,
			Date:          appointments.Date,
			StartTime:     appointments.StartTime,
			EndTime:       appointments.EndTime,
		},
	}

	// Append each element of appointmentpayloads to Appointmentpayload
	Appointmentpayload = append(Appointmentpayload, appointmentpayloads...)

	// // Call the function to handle the cutting result notification
	// err = utility.HandleAppointmentResultNotification(p.AppointmentID, Appointmentpayloads)
	// if err != nil {
	// 	log.Printf("Error notifying cutting result for item ID %d: %v", p.Appoi, err)
	// }
	userDetails := &userpb.User{
		Userid: p.UserID, // Convert uint32 into a User struct.
	}

	ctx := context.Background()
	// Now call the method with the correct type
	response, err := b.UserClient.UserDetails(ctx, userDetails)
	if err != nil {
		log.Fatalf("Failed to fetch doctor details: %v", err)
	}

	log.Println("=========================================================================", response.Email)
	log.Println("=========================================================================", payment.Appoinmentid)
	log.Println("=========================================================================", payment.PaymentID)
	log.Println("=========================================================================", payment.PaymentAmount)

	startTimestamp := &timestamppb.Timestamp{Seconds: appointments.StartTime.Unix()}

	// Convert EndTime to timestamp
	endTimestamp := &timestamppb.Timestamp{Seconds: appointments.EndTime.Unix()}
	// Now you can pass endTimestamp to your event handling function

	startTimeConverted := startTimestamp.AsTime()
	endTimeConverted := endTimestamp.AsTime()

	// Publish Payment Success Event
	err = utility.HandlePaymentNotification(payment.PaymentID, payment.Appoinmentid, payment.UserID, response.Email, payment.PaymentAmount, startTimeConverted, endTimeConverted, appointments.Date, time.Now())
	if err != nil {
		log.Printf("Failed to publish payment event: %v", err)
	}

	err = utility.HandleAppointmentResultNotification(uint32(payment.Appoinmentid), Appointmentpayload)
	if err != nil {
		log.Printf("Error notifying appointment result for appointmentid ID %d: %v", payment.Appoinmentid, err)
	}

	updatedPaymentData, _ := json.Marshal(payment)
	err = b.Redis.SetDataInRedis(key, updatedPaymentData, 0) // Ensure that the cache is updated
	if err != nil {
		log.Printf("Failed to update Redis cache: %v", err)
	}

	// Optionally remove the payment data from Redis if no longer needed
	err = b.Redis.DeleteDataFromRedis(key)
	if err != nil {
		log.Printf("Failed to delete Redis cache after payment update: %v", err)
	}

	log.Printf("Payment successfully completed and updated for payment_id: %s", payment.PaymentID)
	return &pb.PaymentStatusResponse{
		Status:  pb.PaymentStatusResponse_SUCCESS,
		Message: "Payment successfully completed.",
	}, nil
}

func (b *BookingService) CancelAppointmentService(p *pb.Cancelappointmentreq) (*pb.Response, error) {
	appointmentstatus := "Completed"
	err := b.Repo.UpdateAppointmentstatus(p.Userid, p.Appointmentid, appointmentstatus)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: fmt.Sprintf("failed to cancel appointment: %v", err),
		}, err
	}
	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "successfully cancelled appointment appointmentstatus:",
	}, nil
}