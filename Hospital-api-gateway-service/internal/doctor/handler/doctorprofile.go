package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	pb "github.com/anandu86130/Hospital-api-gateway/internal/doctor/pbD"
	"github.com/anandu86130/Hospital-api-gateway/internal/model"
	"github.com/gin-gonic/gin"
)

// ViewProfileHandler handles the user profile view request.
func ViewProfileHandler(c *gin.Context, client pb.DoctorServiceClient) {
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

	response, err := client.ViewProfile(ctx, &pb.ID{
		ID: uint32(userID),
	})

	if err != nil {
		log.Printf("Error in ViewProfile gRPC call: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error in client response",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "profile fetched successfully",
		"data":    response,
	})
}

// EditProfileHandler handles the user profile edit request.
func EditProfileHandler(c *gin.Context, client pb.DoctorServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	var doctor model.EditDoctor

	if err := c.BindJSON(&doctor); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while binding json",
		})
		return
	}

	id, ok := c.Get("user_id")
	if !ok {
		log.Println("Error: Unable to fetch user ID from context")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while fetching user id from context",
		})
		return
	}

	doctorID, ok := id.(uint)
	if !ok {
		log.Println("Error: User ID type assertion failed")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while user id converting",
		})
		return
	}

	response, err := client.EditProfile(ctx, &pb.DoctorProfile{
		Id:                uint32(doctorID),
		Name:              doctor.Name,
		YearsOfExperience: uint32(doctor.YearsOfExperience),
		Fees:              uint32(doctor.Fees),
	})
	if err != nil {
		log.Printf("Error in EditProfile gRPC call: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    response,
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "Profile edited successfully",
		"data":    response,
	})
}

// ChangePasswordHandler handles the user password change request.
func ChangePasswordHandler(c *gin.Context, client pb.DoctorServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	var doctor model.Password

	if err := c.BindJSON(&doctor); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while binding json",
		})
		return
	}

	id, ok := c.Get("user_id")
	if !ok {
		log.Println("Error: Unable to fetch user ID from context")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while fetching user id from context",
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

	response, err := client.ChangePassword(ctx, &pb.Password{
		User_ID:          uint32(userID),
		Old_Password:     doctor.Old,
		New_Password:     doctor.New,
		Confirm_Password: doctor.Confirm,
	})
	if err != nil {
		log.Printf("Error in ChangePassword gRPC call: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while changing password, please check the passwords",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "Password updated successfully",
		"data":    response,
	})
}
