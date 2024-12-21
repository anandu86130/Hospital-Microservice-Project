package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/anandu86130/Hospital-api-gateway/internal/model"
	pb "github.com/anandu86130/Hospital-api-gateway/internal/user/pbU"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ViewAvailabilityHandler(c *gin.Context, client pb.UserServiceClient) {
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

func BookAppointmentHandler(c *gin.Context, client pb.UserServiceClient) {
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

	var input model.AppoinmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Error when binding JSON: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error when binding JSON",
		})
		return
	}
	u := uint32(userID)
	if u != input.Userid {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "mismatched user id",
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
	appoinment := model.Appoinment{
		Doctorid:  input.Doctorid,
		Userid:    input.Userid,
		Date:      parsedDate,
		StartTime: parsedStartTime,
		EndTime:   parsedEndTime,
	}

	// doctorresponse, err := client.DoctorDetails(ctx, &pb.Doctor{
	// 	Doctorid: input.Doctorid,
	// })
	// Convert to gRPC message
	date := timestamppb.New(appoinment.Date)
	starttime := timestamppb.New(appoinment.StartTime)
	endtime := timestamppb.New(appoinment.EndTime)

	// Make the gRPC call to add availability
	response, err := client.BookAppoinment(ctx, &pb.Appoinment{
		Doctorid:  uint32(appoinment.Doctorid),
		Userid:    appoinment.Userid,
		Date:      date,
		Starttime: starttime,
		Endtime:   endtime,
	})

	if err != nil {
		log.Printf("Failed to book appointment: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Failed to book appointment",
			"error":   err.Error(),
		})
		return
	}

	// If no error, send the success response
	c.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "please complete the payment",
		"data":    response,
	})
}

func ViewAppointmentHandler(c *gin.Context, client pb.UserServiceClient) {
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

	var input model.UserAppointmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Error when binding JSON: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error when binding JSON",
		})
		return
	}
	u := uint32(userID)
	if u != input.UserID {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "mismatched user id",
		})
		return
	}
	// Make the gRPC call to view all availability records
	response, err := client.UserViewAppointment(ctx, &pb.ID{
		ID: uint32(input.UserID),
	})
	if err != nil {
		log.Printf("Failed to retrieve appointment: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve appointment",
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

func CancelAppointmentHandler(c *gin.Context, client pb.UserServiceClient) {
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

	appointmentIDString := c.Query("appointment_id")

	appointmentID, err := strconv.Atoi(appointmentIDString)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"Status":  400,
			"Message": "error in converting orderID to int",
			"Error":   err.Error(),
		})
		return
	}
	// Make the gRPC call to view all availability records
	response, err := client.CancelAppointment(ctx, &pb.Cancelappointmentreq{
		Userid:        uint32(userID),
		Appointmentid: uint32(appointmentID),
	})
	if err != nil {
		log.Printf("Failed to cancel appointment: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to cancel appointment",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "appointment cancelled successfully",
		"data":    response,
	})
}

func ViewPrescriptionHandler(c *gin.Context, client pb.UserServiceClient) {
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

	var input model.UserPrescriptionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Error when binding JSON: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error when binding JSON",
		})
		return
	}
	u := uint32(userID)
	if u != input.UserID {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "mismatched user id",
		})
		return
	}
	// Make the gRPC call to view all availability records
	response, err := client.ViewPrescription(ctx, &pb.Req{
		Userid:       input.UserID,
		Appoinmentid: uint32(input.Appoinmentid),
	})
	if err != nil {
		log.Printf("Failed to retrieve prescription: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve prescription",
			"error":   err.Error(),
		})
		return
	}

	// Convert gRPC response to JSON format and send it
	prescriptions := make([]gin.H, len(response.Profiles))
	for i, prescription := range response.Profiles {
		prescriptions[i] = gin.H{
			"prescription_id":        prescription.Id,
			"doctor_id":              prescription.Doctorid,
			"user_id":                prescription.Userid,
			"medicine":               prescription.Medicine,
			"notes":                  prescription.Notes,
			"appointment_date":       prescription.Date.AsTime().Format("2006-01-02"),
			"appointment_start_time": prescription.Starttime.AsTime().Format("15:04:05"),
			"appointment_end_time":   prescription.Endtime.AsTime().Format("15:04:05"),
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"prescription": prescriptions,
	})
}
