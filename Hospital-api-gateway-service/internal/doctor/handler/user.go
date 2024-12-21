package handler

import (
	"context"
	"net/http"
	"time"

	pb "github.com/anandu86130/Hospital-api-gateway/internal/doctor/pbD"
	"github.com/gin-gonic/gin"
)

func UserListHandler(c *gin.Context, client pb.DoctorServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	response, err := client.UserList(ctx, &pb.NoParam{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while fetching users",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "users fetched successfully",
		"data":    response,
	})
}
