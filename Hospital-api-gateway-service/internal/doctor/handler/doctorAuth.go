package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	doctorpb "github.com/anandu86130/Hospital-api-gateway/internal/doctor/pbD"
	"github.com/anandu86130/Hospital-api-gateway/internal/model"
	"github.com/gin-gonic/gin"
)

func DoctorSignupHandler(c *gin.Context, client doctorpb.DoctorServiceClient) {
	ctx, cancel := context.WithTimeout(c, time.Second*3000)
	defer cancel()

	var doctor model.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		log.Printf("Error when binding JSON: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error when binding json",
		})
		return
	}

	response, err := client.DoctorSignup(ctx, &doctorpb.Signup{
		Name:              doctor.Name,
		Email:             doctor.Email,
		Password:          doctor.Password,
		Specialization:    doctor.Specialization,
		YearsOfExperience: uint32(doctor.YearsOfExperience),
		Fees:              uint32(doctor.Fees),
	})
	if err != nil {
		log.Printf("Error when signing up doctor %v: %v", doctor.Email, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error when signing up doctor, please check your input data",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v OTP generated successfully", doctor.Email),
		"data":    response,
	})
}

func DoctorVerifyOTPHandler(c *gin.Context, client doctorpb.DoctorServiceClient) {
	ctx, cancel := context.WithTimeout(c, time.Second*3000)
	defer cancel()

	var request model.OTP
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error when binding JSON: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error when binding json",
		})
		return
	}

	response, err := client.VerifyOTP(ctx, &doctorpb.OTP{
		Email: request.Email,
		Otp:   request.Otp,
	})

	if err != nil {
		log.Printf("Error when verifying OTP for doctor %v: %v", request.Email, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error when verifying otp; please check your otp",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": "OTP verified successfully",
		"data":    response,
	})
}

func DoctorLoginHandler(c *gin.Context, client doctorpb.DoctorServiceClient) {
	ctx, cancel := context.WithTimeout(c, time.Second*3000)
	defer cancel()

	var doctor model.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		log.Printf("Error when binding JSON: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error when binding json",
		})
		return
	}

	fmt.Println("=====================", doctor.Password)
	response, err := client.DoctorLogin(ctx, &doctorpb.Login{
		Email:    doctor.Email,
		Password: doctor.Password,
	})

	if err != nil {
		log.Printf("Error when logging in doctor %v: %v", doctor.Email, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error when logging doctor, please check your email and password",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusAccepted,
		"token":  response,
	})
}
