package user

import (
	"github.com/anandu86130/Hospital-api-gateway/internal/user/handler"
	"github.com/gin-gonic/gin"
)

func (u *User) UserSignup(c *gin.Context) {
	handler.UserSignupHandler(c, u.client)
}

func (u *User) UserLogin(c *gin.Context) {
	handler.UserLoginHandler(c, u.client)
}

func (u *User) VerifyOTP(c *gin.Context) {
	handler.UserVerifyOTPHandler(c, u.client)
}

func (u *User) ViewProfile(c *gin.Context) {
	handler.ViewProfileHandler(c, u.client)
}

func (u *User) EditProfile(c *gin.Context) {
	handler.EditProfileHandler(c, u.client)
}

func (u *User) ChangePassword(c *gin.Context) {
	handler.ChangePasswordHandler(c, u.client)
}

func (u *User) ViewDoctor(c *gin.Context) {
	handler.ViewDoctorProfileHandler(c, u.client)
}

func (u *User) ViewAvailability(c *gin.Context) {
	handler.ViewAvailabilityHandler(c, u.client)
}

func (u *User) BookAppoinment(c *gin.Context) {
	handler.BookAppointmentHandler(c, u.client)
}

func (u *User) ViewAppointment(c *gin.Context) {
	handler.ViewAppointmentHandler(c, u.client)
}

func (u *User) CancelAppointment(c *gin.Context){
	handler.CancelAppointmentHandler(c, u.client)
}

func (u *User) ViewPrescription(c *gin.Context) {
	handler.ViewPrescriptionHandler(c, u.client)
}

func (u *User) Payment(c *gin.Context) {
	handler.UserPaymentHandler(c, u.client)
}

func (u *User) PaymentSuccess(ctx *gin.Context) {
	handler.UserPaymentSuccessHandler(ctx, u.client)
}

func (u *User) PaymentSuccessPage(ctx *gin.Context) {
	handler.PaymentSuccessPageHandler(ctx, u.client)
}