package models

import "time"

type AppointmentResultEvent struct {
	Appointments []AppointmentPayload `json:"appointments"`
}

type AppointmentPayload struct {
	AppointmentID uint32    `json:"appointment_id"`
	Doctorid      uint32    `gorm:"column:doctorid" json:"doctorid"`
	Userid        uint32    `gorm:"column:userid" json:"userid"`
	Date          time.Time `gorm:"type:date" json:"date"`
	StartTime     time.Time `gorm:"type:time" json:"starttime"`
	EndTime       time.Time `gorm:"type:time" json:"endtime"`
}
