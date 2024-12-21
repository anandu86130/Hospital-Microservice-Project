package interfaces

import (
	"github.com/anandu86130/Hospital-doctor-service/internal/model"
)

type DoctorRepository interface {
	CreateDoctor(doctor *model.Doctor) error
	CreateOTP(doctor *model.OTP) error
	UpdateOTP(doctor *model.OTP) error
	FindOTPByEmail(email string) (*model.OTP, error)
	VerifyOTPcheck(email string, otp string) error
	FindDoctorByEmail(email string) (*model.Doctor, error)
	UpdateDoctor(user *model.Doctor) error
	DoctorExists(email string) (bool, error)
	FindDoctorByID(doctorID uint) (*model.Doctor, error)
	GetDoctorList() ([]*model.Doctor, error)
	DoctorDetails(doctorid uint) (*model.Doctor, error)
}