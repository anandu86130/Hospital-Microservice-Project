package interfaces

import (
	"booking-service/internal/model"
	"time"
)

type BookingRepository interface {
	// CheckAvailabilityFields() (*model.Availability, error)
	CreateAvailability(pb *model.Availability) error
	CheckAvailabilityByDate(doctorID uint32, date time.Time, starttime time.Time, endtime time.Time) (*model.Availability, error)
	UpdateAvailability(pb *model.Availability) error
	GetAllAvailabilities() ([]*model.Availability, error)
	GetAppointmentsByDoctorAndDate(doctorID uint32, date time.Time, starttime time.Time, endtime time.Time) ([]model.Appoinment, error)
	CreateAppointment(pb *model.Appoinment) (uint, error)
	GetDoctorAppointment(id uint32, paymentstatus string) ([]*model.Appoinment, error)
	GetDoctorUserAppointment(id uint32, paymentstatus string) ([]*model.Appoinment, error)
	GetAllAppointment() ([]*model.Appoinment, error)
	CheckOrCreatePrescription(appointmentid uint32, doctorID uint32, userid uint32, medicine string, notes string) (*model.Prescription, error)
	CheckAppointment(appointmentid uint32, doctorid uint32, userid uint32) error
	ViewPrescription(userid uint32, appoinmentid uint32) ([]*model.Prescription, error)
	CheckAppointmentStatus(userid uint32, appointmentid uint32, paymentstatus string) (*model.Appoinment, error)
	SavePayment(payment *model.Payment) error
	GetLatestPaymentByAppointmentID(appoinmentID int) (model.Payment, error)
	FindAppointmentsByID(AppointmnentID uint) (*model.Appoinment, error)
	UpdatePaymentAndAppointmentStatus(paymentID string, AppointmentID int, paymentStatus string, appointmentStatus string, Appointmentstatus string) error
	UpdateAppointmentstatus(userid uint32, appointmentid uint32, appointmentstatus string) error
}
