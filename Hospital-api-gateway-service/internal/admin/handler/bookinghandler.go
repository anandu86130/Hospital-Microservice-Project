package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	pb "github.com/anandu86130/Hospital-api-gateway/internal/admin/pb"
	"github.com/gin-gonic/gin"
)

func ViewAllAppointmentHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 5
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	// Make the gRPC call to view all availability records
	response, err := client.ViewAllAppointment(ctx, &pb.NoParam{})
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
	appointments := make([]gin.H, len(response.Profiles))
	for i, appoinment := range response.Profiles {
		appointments[i] = gin.H{
			"appointment_id": appoinment.Id,
			"user_id":        appoinment.Userid,
			"doctor_id":      appoinment.Doctorid,
			"date":           appoinment.Date.AsTime().Format("2006-01-02"),
			"start_time":     appoinment.Starttime.AsTime().Format("15:04:05"),
			"end_time":       appoinment.Endtime.AsTime().Format("15:04:05"),
			"payment_status": appoinment.Paymentstatus,
			"amount":         appoinment.Amount,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":         http.StatusOK,
		"availabilities": appointments,
	})
}
