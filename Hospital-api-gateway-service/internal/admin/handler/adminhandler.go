package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	pb "github.com/anandu86130/Hospital-api-gateway/internal/admin/pb"
	"github.com/anandu86130/Hospital-api-gateway/internal/model"
	"github.com/gin-gonic/gin"
)

// AdminLoginHandler function will send the login request to client.
func AdminLoginHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	var admin model.Login
	if err := c.BindJSON(&admin); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while binding json",
		})
		return
	}

	response, err := client.AdminLoginRequest(ctx, &pb.AdminLogin{
		Email:    admin.Email,
		Password: admin.Password,
	})
	if err != nil {
		log.Printf("Error in AdminLoginRequest gRPC call: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while logging in, please check email and password",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "Login successful",
		"data":    response,
	})
}
