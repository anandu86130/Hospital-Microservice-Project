package repository

import (
	"booking-service/internal/model"
	inter "booking-service/internal/repository/interfaces"
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type BookingRepository struct {
	DB *gorm.DB
}

// CheckDoctorExistsByLiscenceNumber implements interfaces.DoctorRepository
func NewBookingRepository(db *gorm.DB) inter.BookingRepository {
	return &BookingRepository{
		DB: db,
	}
}

func (r *BookingRepository) CreateAvailability(availability *model.Availability) error {
	if err := r.DB.Create(availability).Error; err != nil {
		log.Printf("Database error: %v", err) // Log the database error
		return err
	}
	return nil
}

func (r *BookingRepository) CheckAvailabilityByDate(doctorID uint32, date time.Time, starttime time.Time, endtime time.Time) (*model.Availability, error) {
	var availability model.Availability

	// Query the database for availability for the given doctor ID and date
	err := r.DB.Where("doctorid = ? AND date = ? AND start_time <= ? AND end_time >= ?", doctorID, date, starttime, endtime).First(&availability).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // No availability found for this date, which is expected
	} else if err != nil {
		return nil, err // Return any other error encountered
	}

	return &availability, nil // Availability found
}

func (r *BookingRepository) CheckOrCreatePrescription(appointmentid uint32, doctorID uint32, userid uint32, medicine string, notes string) (*model.Prescription, error) {
	var prescription model.Prescription

	// Check if the prescription already exists
	err := r.DB.Where("doctorid = ? AND userid = ? AND appoinmentid = ? AND medicine = ?", doctorID, userid, appointmentid, medicine).First(&prescription).Error
	if err == nil {
		// Prescription already exists, return an error
		return nil, fmt.Errorf("prescription already exists for this appointment")
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// Return any other error encountered during the query
		return nil, err
	}

	// Prescription does not exist, create a new one
	newPrescription := model.Prescription{
		Appoinmentid: appointmentid, // Correct spelling if needed
		Doctorid:     doctorID,
		Userid:       userid,
		Medicine:     medicine,
		Notes:        notes,
	}
	if err := r.DB.Create(&newPrescription).Error; err != nil {
		return nil, err // Return any error encountered during creation
	}

	return &newPrescription, nil // Return the newly created prescription
}

// UpdateAvailability updates an existing availability record in the database
func (r *BookingRepository) UpdateAvailability(availability *model.Availability) error {
	// Find the existing availability record
	var existingAvailability model.Availability
	err := r.DB.Where("doctorid = ? AND date = ?", availability.Doctorid, availability.Date).First(&existingAvailability).Error
	if err != nil {
		// If no record is found, return an error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("availability not found for doctor %d on date %v", availability.Doctorid, availability.Date)
		}
		// Handle any other errors from GORM
		return fmt.Errorf("error fetching availability: %v", err)
	}

	// Update the availability fields
	existingAvailability.StartTime = availability.StartTime
	existingAvailability.EndTime = availability.EndTime

	// Save the updated availability
	err = r.DB.Save(&existingAvailability).Error
	if err != nil {
		return fmt.Errorf("error updating availability: %v", err)
	}
	return nil
}

func (r *BookingRepository) GetAllAvailabilities() ([]*model.Availability, error) {
	var availabilities []*model.Availability
	if err := r.DB.Find(&availabilities).Error; err != nil {
		return nil, err
	}
	return availabilities, nil
}

func (r *BookingRepository) GetAppointmentsByDoctorAndDate(doctorID uint32, date time.Time, starttime, endtime time.Time) ([]model.Appoinment, error) {
	var appointments []model.Appoinment

	// Query to find any overlapping appointments for the given doctor and date
	err := r.DB.Where("doctorid = ? AND DATE(date) = ? AND ((start_time < ? AND end_time > ?) OR (start_time < ? AND end_time > ?))",
		doctorID, date.Truncate(24*time.Hour), endtime, starttime, starttime, endtime).Find(&appointments).Error

	if err != nil {
		return nil, err // Return error if encountered
	}

	return appointments, nil // Return the list of overlapping appointments, if any
}

func (r *BookingRepository) CreateAppointment(appointment *model.Appoinment) (uint, error) {
	// Assuming you're using GORM or another ORM for the database interaction
	err := r.DB.Create(appointment).Error
	if err != nil {
		return 0, err // Return 0 or some other value if there's an error
	}

	return appointment.ID, nil // Assuming 'ID' is the field for the appointment's primary key
}

func (r *BookingRepository) GetDoctorAppointment(id uint32, paymentstatus string) ([]*model.Appoinment, error) {
	var appointment []*model.Appoinment
	err := r.DB.Where("doctorid = ? AND paymentstatus = ? AND appointmentstatus = ?", id, paymentstatus, paymentstatus).Find(&appointment).Error
	if err != nil {
		// If no record is found, return an error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("failed to find doctor")
		}
		// Handle any other errors from GORM
		return nil, errors.New("failed to fetch appointment")
	}
	return appointment, nil
}

func (r *BookingRepository) GetDoctorUserAppointment(id uint32, paymentstatus string) ([]*model.Appoinment, error) {
	var appointment []*model.Appoinment
	err := r.DB.Where("userid = ? AND paymentstatus = ? AND appointmentstatus = ?", id, paymentstatus, paymentstatus).Find(&appointment).Error
	if err != nil {
		// If no record is found, return an error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("failed to find user in appointment")
		}
		// Handle any other errors from GORM
		return nil, errors.New("failed to fetch appointment")
	}
	return appointment, nil
}

func (r *BookingRepository) GetAllAppointment() ([]*model.Appoinment, error) {
	var appointments []*model.Appoinment
	if err := r.DB.Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *BookingRepository) CheckAppointment(appointmentid, doctorID, userid uint32) error {
	var appointment model.Appoinment

	// Check if the appointment exists
	err := r.DB.Where("id = ? AND doctorid = ? AND userid = ?", appointmentid, doctorID, userid).First(&appointment).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Appointment does not exist, return an error
		return fmt.Errorf("appointment does not exist for the specified details")
	} else if err != nil {
		// Return any other error encountered during the query
		return err
	}

	return nil // Appointment exists, no error
}

func (r *BookingRepository) ViewPrescription(userid uint32, appoinmentid uint32) ([]*model.Prescription, error) {
	var prescription []*model.Prescription
	err := r.DB.Preload("Appoinment").Where("userid = ? AND appoinmentid = ?", userid, appoinmentid).Find(&prescription).Error

	if err != nil {
		// Handle any errors from GORM
		return nil, errors.New("failed to fetch prescription")
	}

	// Check if no records were found
	if len(prescription) == 0 {
		return nil, errors.New("no prescriptions found for the specified appointment ID")
	}

	return prescription, nil
}

func (b *BookingRepository) SavePayment(payment *model.Payment) error {
	if err := b.DB.Create(&payment).Error; err != nil {
		return fmt.Errorf("failed to save payment: %v", err)
	}
	log.Println("Payment saved successfully")
	return nil
}

func (b *BookingRepository) GetLatestPaymentByAppointmentID(appointmentID int) (model.Payment, error) {
	var payment model.Payment
	query := `
        SELECT payment_id, appoinmentid, payment_amount, status, client_secret, payment_method, user_id
        FROM payments
        WHERE appointmentid = ?
        ORDER BY 
            CASE 
                WHEN status = 'Pending' THEN 1
                WHEN status = 'Completed' THEN 2
                ELSE 3
            END, 
            payment_id DESC
        LIMIT 1`
	result := b.DB.Raw(query, appointmentID).Scan(&payment)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return payment, nil // No payment found for the appointmentid
		}
		return payment, fmt.Errorf("failed to fetch latest payment: %v", result.Error)
	}
	return payment, nil
}

func (b *BookingRepository) FindAppointmentsByID(AppointmentsID uint) (*model.Appoinment, error) {
	var Appointments model.Appoinment
	if err := b.DB.First(&Appointments, AppointmentsID).Error; err != nil {
		return nil, err
	}
	return &Appointments, nil
}

func (b *BookingRepository) CheckAppointmentStatus(userid uint32, appointmentid uint32, appointmentstatus string) (*model.Appoinment, error) {
	var appointment *model.Appoinment
	err := b.DB.Where("userid = ? AND id = ? AND appointmentstatus = ?", userid, appointmentid, appointmentstatus).Find(&appointment).Error
	if err != nil {
		// If no record is found, return an error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("failed to find user in appointment")
		}
		// Handle any other errors from GORM
		return nil, errors.New("failed to check paymentstatus")
	}
	return appointment, nil
}

func (m *BookingRepository) UpdatePaymentAndAppointmentStatus(paymentID string, AppointmentID int, paymentStatus string, appointmentStatus string, Appointmentstatus string) error {
	// Begin a transaction
	tx := m.DB.Begin()

	// Update payment status
	if err := tx.Model(&model.Payment{}).
		Where("payment_id = ?", paymentID).
		Update("status", paymentStatus).Error; err != nil {
		tx.Rollback() // Roll back the transaction on error
		return fmt.Errorf("failed to update payment status: %v", err)
	}

	// Update appointment status
	if err := tx.Model(&model.Appoinment{}).
		Where("id = ?", AppointmentID).
		Update("paymentstatus", appointmentStatus).Error; err != nil {
		tx.Rollback() // Roll back the transaction on error
		return fmt.Errorf("failed to update appointment status: %v", err)
	}

	// Update appointment status
	if err := tx.Model(&model.Appoinment{}).
		Where("id = ?", AppointmentID).
		Update("paymentstatus", appointmentStatus).Error; err != nil {
		tx.Rollback() // Roll back the transaction on error
		return fmt.Errorf("failed to update appointment status: %v", err)
	}

	if err := tx.Model(&model.Appoinment{}).
		Where("id = ?", AppointmentID).
		Update("appointmentstatus", Appointmentstatus).Error; err != nil {
		tx.Rollback() // Roll back the transaction on error
		return fmt.Errorf("failed to update appointment status: %v", err)
	}

	// Commit the transaction if both updates succeed
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	log.Println("Payment and order status updated successfully")
	return nil
}

func (m *BookingRepository) UpdateAppointmentstatus(userid uint32, appointmentid uint32, appointmentstatus string) error {
	var existingappointment model.Appoinment
	err := m.DB.Where("userid = ? AND ID = ? AND appointmentstatus = ?", userid, appointmentid, appointmentstatus).First(&existingappointment).Error
	if err != nil {
		// If no record is found, return an error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("appointment not found for user %d on appointmentid %v", userid, appointmentid)
		}
		// Handle any other errors from GORM
		return fmt.Errorf("error fetching appointment: %v", err)
	}

	// Update the availability fields
	existingappointment.AppoinmentStatus = "Cancelled"

	// Save the updated availability
	err = m.DB.Save(&existingappointment).Error
	if err != nil {
		return fmt.Errorf("error updating appointmentstatus: %v", err)
	}
	return nil
}
