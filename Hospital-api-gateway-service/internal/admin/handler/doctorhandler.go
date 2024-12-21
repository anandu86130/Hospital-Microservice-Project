package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	pb "github.com/anandu86130/Hospital-api-gateway/internal/admin/pb"
	"github.com/gin-gonic/gin"
)

// BlockUserHandler function will send block user request to client.
func BlockDoctorHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	UserIDString := c.Param("id")
	UserID, err := strconv.Atoi(UserIDString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while converting userID to int",
		})
		return
	}

	response, err := client.AdminBlockDoctor(ctx, &pb.AdID{
		ID: uint32(UserID),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while blocking doctor, please check the id and try again",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "Doctor blocked successfully",
		"data":    response,
	})
}

// BlockUserHandler function will send block user request to client.
func UnblockDoctorHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	UserIDString := c.Param("id")
	UserID, err := strconv.Atoi(UserIDString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while converting userID to int",
		})
		return
	}

	response, err := client.AdminUnblockDoctor(ctx, &pb.AdID{
		ID: uint32(UserID),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"Message": "error while unblocking doctor, please check the id and try again",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "doctor unblocked successfully",
		"data":    response,
	})
}

func IsVerifiedHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	UserIDString := c.Param("id")
	UserID, err := strconv.Atoi(UserIDString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while converting userID to int",
		})
		return
	}

	response, err := client.AdminIsVerified(ctx, &pb.AdID{
		ID: uint32(UserID),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while verifying doctor, please check the id and try again",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "Doctor verified successfully",
		"Data":    response,
	})
}

func DoctorListHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	response, err := client.DoctorList(ctx, &pb.NoParam{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while fetching doctors",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "doctors fetched successfully",
		"data":    response,
	})
}
