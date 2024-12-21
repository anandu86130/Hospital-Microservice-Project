package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/anandu86130/Hospital-api-gateway/internal/model"
	pb "github.com/anandu86130/Hospital-api-gateway/internal/user/pbU"
	"github.com/gin-gonic/gin"
)

// ViewProfileHandler handles the user profile view request.
func ViewProfileHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id from context",
			"Error":   "",
		})
		return
	}

	userID, ok := id.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id converting",
			"Error":   ""})
		return
	}

	response, err := client.ViewProfile(ctx, &pb.ID{
		ID: uint32(userID),
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Data":    response,
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "profile fetched successfully",
		"Data":    response,
	})
}

// EditProfileHandler handles the user profile edit request.
func EditProfileHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	var user model.UserModel

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while binding json",
			"Error":   err.Error()})
		return
	}

	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id from context",
			"Error":   ""})
		return
	}

	userID, ok := id.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id converting",
			"Error":   ""})
		return
	}

	response, err := client.EditProfile(ctx, &pb.Profile{
		User_ID: uint32(userID),
		Name:    user.Name,
		Gender:  user.Gender,
		Age:     user.Age,
		Number:  user.Number,
		Address: user.Address,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Data":    response,
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "Profile edited succesfully",
		"Data":    response,
	})
}

// ChangePasswordHandler handles the user password change request.
func ChangePasswordHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	var user model.Password

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while binding json",
			"Error":   err.Error()})
		return
	}

	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id from context",
			"Error":   ""})
		return
	}

	userID, ok := id.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id converting",
			"Error":   ""})
		return
	}

	response, err := client.ChangePassword(ctx, &pb.Password{
		User_ID:          uint32(userID),
		Old_Password:     user.Old,
		New_Password:     user.New,
		Confirm_Password: user.Confirm,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Data":    response,
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "PAssword updated succesfully",
		"Data":    response,
	})
}

func ViewDoctorProfileHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	response, err := client.DoctorList(ctx, &pb.NoParam{})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Data":    response,
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "profile fetched successfully",
		"Data":    response,
	})
}
