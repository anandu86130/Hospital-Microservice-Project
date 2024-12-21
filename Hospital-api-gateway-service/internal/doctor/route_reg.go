package doctor

import (
	doctorhandler "github.com/anandu86130/Hospital-api-gateway/internal/doctor/handler"
	"github.com/gin-gonic/gin"
)

func (a *Doctor) DoctorSignup(c *gin.Context) {
	doctorhandler.DoctorSignupHandler(c, a.client)
}

func (a *Doctor) DoctorLogin(c *gin.Context) {
	doctorhandler.DoctorLoginHandler(c, a.client)
}

func (u *Doctor) VerifyOTP(c *gin.Context) {
	doctorhandler.DoctorVerifyOTPHandler(c, u.client)
}

func (u *Doctor) ViewProfile(c *gin.Context) {
	doctorhandler.ViewProfileHandler(c, u.client)
}

func (u *Doctor) EditProfile(c *gin.Context) {
	doctorhandler.EditProfileHandler(c, u.client)
}

func (u *Doctor) ChangePassword(c *gin.Context) {
	doctorhandler.ChangePasswordHandler(c, u.client)
}

func (u *Doctor) AddAvailability(c *gin.Context) {
	doctorhandler.AddAvailabilityHandler(c, u.client)
}

func (u *Doctor) ViewAvailability(c *gin.Context) {
	doctorhandler.ViewAvailabilityHandler(c, u.client)
}

func (u *Doctor) UserList(c *gin.Context) {
	doctorhandler.UserListHandler(c, u.client)
}

func (u *Doctor) ViewAppointment(c *gin.Context) {
	doctorhandler.ViewAppointmentHandler(c, u.client)
}

func (u *Doctor) AddPrescription(c *gin.Context) {
	doctorhandler.AddPrescriptionHandler(c, u.client)
}
