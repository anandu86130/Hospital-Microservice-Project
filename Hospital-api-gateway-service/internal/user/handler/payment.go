package handler

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	userpb "github.com/anandu86130/Hospital-api-gateway/internal/user/pbU"
	"github.com/gin-gonic/gin"
)

func UserPaymentHandler(c *gin.Context, client userpb.UserServiceClient) {
	timeOut := time.Second * 100
	_, cancel := context.WithTimeout(c, timeOut)
	defer cancel()

	userIDstring := c.Query("id")
	appointmentIDString := c.Query("appointment_id")
	fmt.Println("===========================", userIDstring)
	fmt.Println("====================", appointmentIDString)

	userID, err := strconv.Atoi(userIDstring)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"Status":  400,
			"Message": "error in converting userID to int",
			"Error":   err.Error(),
		})
		return
	}

	appointmentID, err := strconv.Atoi(appointmentIDString)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"Status":  400,
			"Message": "error in converting orderID to int",
			"Error":   err.Error(),
		})
		return
	}
	log.Println("userid", userIDstring)
	log.Println("appointmentid", appointmentIDString)

	ctx := context.Background()
	response, err := client.CreatePayment(ctx, &userpb.ConfirmAppointment{
		Userid:       uint32(userID),
		Appoinmentid: uint32(appointmentID),
	})

	// Check for any errors in the client response
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"Status":  400,
			"Message": "error in client response",
			"Error":   err.Error(),
		})
		return
	}

	// Check if response is nil
	if response == nil {
		c.AbortWithStatusJSON(500, gin.H{
			"Status":  500,
			"Message": "Payment response is nil",
		})
		return
	}

	// Log the response (for debugging purposes)
	log.Println("User create response:", response)

	// // Proceed with rendering the HTML page
	c.HTML(200, "stripe.html", gin.H{
		"userID":        userID,
		"appointmentID": appointmentID,
		"paymentID":     response.Paymentid,
		"amount":        response.Amount,
		"client":        response.ClientSecret,
	})
}

// Define a struct for the expected request body
type PaymentRequest struct {
	AppointmentID string `json:"appointment_id"`
	UserID        string `json:"user_id"`
	Amount        uint32 `json:"amount"`
	PaymentID     string `json:"paymentID"`
	ClientSecret  string `json:"clientSecret"`
}

func UserPaymentSuccessHandler(c *gin.Context, client userpb.UserServiceClient) {
	// Set a timeout for the request context
	timeOut := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeOut)
	defer cancel()

	// Bind the JSON data from the request body to the struct
	var req PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"Status":  400,
			"Message": "Invalid request body",
			"Error":   err.Error(),
		})
		return
	}
	log.Println("response data", req)

	// Convert userID and orderID from string to int
	userID, err := strconv.Atoi(req.UserID)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"Status":  400,
			"Message": "Error in converting userID to int",
			"Error":   err.Error(),
		})
		return
	}

	appoinmentid, err := strconv.Atoi(req.AppointmentID)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"Status":  400,
			"Message": "Error in converting appointmentid to int",
			"Error":   err.Error(),
		})
		return
	}

	log.Println(appoinmentid)
	// Call the client method to process the payment
	_, err = client.UserPaymentSuccess(ctx, &userpb.UserPayment{
		UserID:        uint32(userID),
		PaymentID:     req.PaymentID,
		Amount:        req.Amount,
		Appointmentid: uint32(appoinmentid),
	})
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"Status":  400,
			"Message": "Error in client response",
			"Error":   err.Error(),
		})
		return
	}

	// Send a successful response
	c.JSON(200, gin.H{
		"status": true,
	})
}

func PaymentSuccessPageHandler(ctx *gin.Context, client userpb.UserServiceClient) {
	// Extract the "payment" query parameter
	paymentID := ctx.Query("paymentID")

	// Render the success page with the payment ID
	ctx.HTML(200, "success.html", gin.H{
		"paymentID": paymentID,
	})
}
