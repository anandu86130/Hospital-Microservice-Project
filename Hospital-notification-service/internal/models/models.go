package models

import "time"

type PaymentEvent struct {
	PaymentID            string    `json:"payment_id"`
	AppointmentID        uint      `json:"appointment_id"`
	UserID               uint32    `json:"user_id"`
	Email                string    `json:"email"`
	Amount               uint      `json:"amount"`
	Date                 string    `json:"date"`
	AppointmentDate      time.Time `json:"appointmentdate"`
	AppointmentStartTime time.Time `json:"appointmentstarttime"`
	AppointmentEndTime   time.Time `json:"appointmentendtime"`
}
