package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/anandu86130/Hospital-api-gateway/internal/model"
	userpb "github.com/anandu86130/Hospital-api-gateway/internal/user/pbU"
	"github.com/gin-gonic/gin"
)

func UserSignupHandler(c *gin.Context, client userpb.UserServiceClient) {
	ctx, cancel := context.WithTimeout(c, time.Second*3000)
	defer cancel()

	var user model.UserModel
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("error when binding json:%v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "error while binding json",
			"Error":   err.Error(),
		})
		return
	}

	response, err := client.UserSignup(ctx, &userpb.Signup{
		Name:     user.Name,
		Email:    user.Email,
		Gender:   user.Gender,
		Age:      user.Age,
		Number:   user.Number,
		Password: user.Password,
		Address:  user.Address,
	})
	if err != nil {
		log.Printf("error when signing up user %v err: %v", user.Email, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "error in client responsse",
			"Data":    response,
			"Error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "successfully send OTP, verify OTP",
		"Data":    response,
	})

	// c.JSON(http.StatusOK, gin.H{
	// 	"status":  http.StatusAccepted,
	// 	"message": fmt.Sprintf("%v otp generated successfully", user.Email),
	// 	"data":    response,
	// })
}

func UserVerifyOTPHandler(c *gin.Context, client userpb.UserServiceClient) {
	ctx, cancel := context.WithTimeout(c, time.Second*3000)
	defer cancel()

	var request model.VerifyOTPs
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("error when binding json :%v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "error when binding json",
			"Error":   err.Error(),
		})
		return
	}

	response, err := client.VerifyOTP(ctx, &userpb.OTP{
		Email: request.Email,
		Otp:   request.Otp,
	})

	if err != nil {
		log.Printf("error when verifying OTP for user %v err: %v", request.Email, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "error in client response",
			"Data":    response,
			"Error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "OTP verified successfully",
		"Token":   response,
	})
}

func UserLoginHandler(c *gin.Context, client userpb.UserServiceClient) {
	ctx, cancel := context.WithTimeout(c, time.Second*3000)
	defer cancel()

	var user model.UserModel
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("error when binding json: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"Message": "error when binding json",
			"error":   err.Error(),
		})
		return
	}

	response, err := client.UserLogin(ctx, &userpb.Login{
		Email:    user.Email,
		Password: user.Password,
	})

	// if user.IsBlocked {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"status":  http.StatusBadRequest,
	// 		"Message": "user blocked by admin",
	// 	})
	// 	return
	// }

	if err != nil {
		log.Printf("error when loggin in user %v err: %v", user.Email, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "error while logging in",
			"error":   err.Error(),
		})
		return
	}

	// if user.IsBlocked {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"status":  http.StatusBadRequest,
	// 		"Message": "user blocked by admin",
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusAccepted,
		"token":  response,
	})
}
