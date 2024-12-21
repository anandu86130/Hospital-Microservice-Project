package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	pb "github.com/anandu86130/Hospital-api-gateway/internal/doctor/pbD"
	"github.com/anandu86130/Hospital-api-gateway/internal/model"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func AddAvailabilityHandler(c *gin.Context, client pb.DoctorServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	id, ok := c.Get("user_id")
	if !ok {
		log.Println("Error: Unable to fetch user ID from context")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while fetching user id",
		})
		return
	}
	userID, ok := id.(uint)
	if !ok {
		log.Println("Error: User ID type assertion failed")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while user id converting",
		})
		return
	}

	var input model.AvailabilityInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Error when binding JSON: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error when binding JSON",
		})
		return
	}

	u := uint32(userID)
	if u != input.Doctorid {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "mismatched doctor id",
		})
		return
	}
	// Parse the date and time fields
	parsedDate, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		log.Printf("Error parsing date: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid date format",
		})
		return
	}

	parsedStartTime, err := time.Parse("15:04:05", input.StartTime)
	if err != nil {
		log.Printf("Error parsing start time: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid start time format",
		})
		return
	}

	parsedEndTime, err := time.Parse("15:04:05", input.EndTime)
	if err != nil {
		log.Printf("Error parsing end time: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid end time format",
		})
		return
	}

	// Construct the availability struct
	availability := model.Availability{
		Doctorid:  input.Doctorid,
		Date:      parsedDate,
		StartTime: parsedStartTime,
		EndTime:   parsedEndTime,
	}

	// Convert to gRPC message
	date := timestamppb.New(availability.Date)
	starttime := timestamppb.New(availability.StartTime)
	endtime := timestamppb.New(availability.EndTime)

	// Make the gRPC call to add availability
	response, err := client.AddAvailability(ctx, &pb.Availability{
		Doctorid:  uint32(availability.Doctorid),
		Date:      date,
		Starttime: starttime,
		Endtime:   endtime,
	})

	if err != nil {
		log.Printf("Failed to add availability: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Failed to add availability",
			"error":   err.Error(),
		})
		return
	}

	// If no error, send the success response
	c.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "Availability added successfully",
		"data":    response,
	})
}

func ViewAvailabilityHandler(c *gin.Context, client pb.DoctorServiceClient) {
	timeout := time.Second * 5
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	// Make the gRPC call to view all availability records
	response, err := client.ViewAvailability(ctx, &pb.NoParam{})
	if err != nil {
		log.Printf("Failed to retrieve availability: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve availability",
			"error":   err.Error(),
		})
		return
	}

	// Convert gRPC response to JSON format and send it
	availabilities := make([]gin.H, len(response.Availabilities))
	for i, availability := range response.Availabilities {
		availabilities[i] = gin.H{
			"doctor_id":  availability.Doctorid,
			"date":       availability.Date.AsTime().Format("2006-01-02"),
			"start_time": availability.Starttime.AsTime().Format("15:04:05"),
			"end_time":   availability.Endtime.AsTime().Format("15:04:05"),
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":         http.StatusOK,
		"availabilities": availabilities,
	})
}

func ViewAppointmentHandler(c *gin.Context, client pb.DoctorServiceClient) {
	timeout := time.Second * 5
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	id, ok := c.Get("user_id")
	if !ok {
		log.Println("Error: Unable to fetch user ID from context")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while fetching user id",
		})
		return
	}
	userID, ok := id.(uint)
	if !ok {
		log.Println("Error: User ID type assertion failed")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while user id converting",
		})
		return
	}

	var input model.DoctorAppointmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Error when binding JSON: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error when binding JSON",
		})
		return
	}
	u := uint32(userID)
	if u != input.DoctorID {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "mismatched doctor id",
		})
		return
	}
	// Make the gRPC call to view all availability records
	response, err := client.ViewAppointment(ctx, &pb.ID{
		ID: uint32(input.DoctorID),
	})
	if err != nil {
		log.Printf("Failed to retrieve appointment: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve availability",
			"error":   err.Error(),
		})
		return
	}

	// Convert gRPC response to JSON format and send it
	appointments := make([]gin.H, len(response.Profiles))
	for i, appointment := range response.Profiles {
		appointments[i] = gin.H{
			"appointment_id": appointment.Id,
			"doctor_id":      appointment.Doctorid,
			"user_id":        appointment.Userid,
			"date":           appointment.Date.AsTime().Format("2006-01-02"),
			"start_time":     appointment.Starttime.AsTime().Format("15:04:05"),
			"end_time":       appointment.Endtime.AsTime().Format("15:04:05"),
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"appointments": appointments,
	})
}

func AddPrescriptionHandler(c *gin.Context, client pb.DoctorServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	id, ok := c.Get("user_id")
	if !ok {
		log.Println("Error: Unable to fetch user ID from context")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while fetching user id",
		})
		return
	}
	userID, ok := id.(uint)
	if !ok {
		log.Println("Error: User ID type assertion failed")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while user id converting",
		})
		return
	}

	var input model.PrescriptionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Error when binding JSON: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error when binding JSON",
		})
		return
	}

	u := uint32(userID)
	if u != input.Doctorid {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "mismatched doctor id",
		})
		return
	}
	// Construct the availability struct
	Prescription := model.PrescriptionInput{
		Appoinmentid: input.Appoinmentid,
		Doctorid:     input.Doctorid,
		Userid:       input.Userid,
		Medicine:     input.Medicine,
		Notes:        input.Notes,
	}

	// Make the gRPC call to add availability
	_, err := client.AddPrescription(ctx, &pb.Prescription{
		Appointmentid: Prescription.Appoinmentid,
		Doctorid:      uint32(Prescription.Doctorid),
		Userid:        Prescription.Userid,
		Medicine:      Prescription.Medicine,
		Notes:         Prescription.Notes,
	})

	if err != nil {
		log.Printf("Failed to add prescription: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "failed to add prescriptuon",
			"error":   err.Error(),
		})
		return
	}

	// If no error, send the success response
	c.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "Prescription added successfully",
	})
}
