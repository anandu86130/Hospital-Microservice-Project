package admin

import (
	"github.com/anandu86130/Hospital-api-gateway/internal/admin/handler"
	"github.com/gin-gonic/gin"
)

func (a *Admin) Login(c *gin.Context) {
	handler.AdminLoginHandler(c, a.Client)
}

func (a *Admin) BlockUser(c *gin.Context) {
	handler.BlockUserHandler(c, a.Client)
}

func (a *Admin) UnBlockUser(c *gin.Context) {
	handler.UnblockUserHandler(c, a.Client)
}

func (a *Admin) BlockDoctor(c *gin.Context) {
	handler.BlockDoctorHandler(c, a.Client)
}

func (a *Admin) UnBlockDoctor(c *gin.Context) {
	handler.UnblockDoctorHandler(c, a.Client)
}

func (a *Admin) IsVerified(c *gin.Context) {
	handler.IsVerifiedHandler(c, a.Client)
}

func (a *Admin) UserList(c *gin.Context) {
	handler.UserListHandler(c, a.Client)
}

func (a *Admin) DoctorList(c *gin.Context) {
	handler.DoctorListHandler(c, a.Client)
}

func (a *Admin) ViewAllAppointment(c *gin.Context) {
	handler.ViewAllAppointmentHandler(c, a.Client)
}
