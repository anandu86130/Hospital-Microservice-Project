package model

import (
	"time"

	"gorm.io/gorm"
)

// type Availability struct {
// 	gorm.Model
// 	Doctorid  uint32    `gorm:"column:doctorid" json:"doctorid"`
// 	Date      time.Time `gorm:"type:date" json:"date" time_format:"2006-01-02"`
// 	StartTime time.Time `gorm:"type:time" json:"starttime" time_format:"15:04:05"`
// 	EndTime   time.Time `gorm:"type:time" json:"endtime" time_format:"15:04:05"`
// }

type AvailabilityInput struct {
	Doctorid  uint32 `json:"doctorid" binding:"required"`
	Date      string `json:"date" binding:"required"`       // Expecting format "YYYY-MM-DD"
	StartTime string `json:"start_time" binding:"required"` // Expecting format "HH:MM:SS"
	EndTime   string `json:"end_time" binding:"required"`   // Expecting format "HH:MM:SS"
}

type Availability struct {
	Doctorid  uint32
	Date      time.Time
	StartTime time.Time
	EndTime   time.Time
}

type AppoinmentInput struct {
	Doctorid  uint32 `json:"doctorid" binding:"required"`
	Userid    uint32 `json:"userid" binding:"required"`
	Date      string `json:"date" binding:"required"`      // Expecting format "YYYY-MM-DD"
	StartTime string `json:"starttime" binding:"required"` // Expecting format "HH:MM:SS"
	EndTime   string `json:"endtime" binding:"required"`   // Expecting format "HH:MM:SS"
}

type Appoinment struct {
	gorm.Model
	Doctorid  uint32
	Userid    uint32
	Date      time.Time
	StartTime time.Time
	EndTime   time.Time
}

type DoctorAppointmentInput struct {
	DoctorID uint32
}

type UserAppointmentInput struct {
	UserID uint32
}

type PrescriptionInput struct {
	Appoinmentid uint32 `json:"appoinmentid" binding:"required"`
	Doctorid     uint32 `json:"doctorid" binding:"required"`
	Userid       uint32 `json:"userid" binding:"required"`
	Medicine     string `json:"medicine" binding:"required"`
	Notes        string `json:"notes"`
}

type UserPrescriptionInput struct {
	UserID       uint32
	Appoinmentid uint32
}
