package model

import (
	"time"

	"gorm.io/gorm"
)

type Availability struct {
	gorm.Model
	Doctorid  uint32    `gorm:"column:doctorid" json:"doctorid"`
	Date      time.Time `gorm:"type:date" json:"date" time_format:"2006-01-02"`
	StartTime time.Time `gorm:"type:time" json:"star_time" time_format:"15:04:05"`
	EndTime   time.Time `gorm:"type:time" json:"end_time" time_format:"15:04:05"`
}

type Appoinment struct {
	gorm.Model
	Doctorid         uint32    `gorm:"column:doctorid" json:"doctorid"`
	Userid           uint32    `gorm:"column:userid" json:"userid"`
	Date             time.Time `gorm:"type:date" json:"date"`
	StartTime        time.Time `gorm:"type:time" json:"starttime"`
	EndTime          time.Time `gorm:"type:time" json:"endtime"`
	PaymentStatus    string    `gorm:"column:paymentstatus" json:"paymentstatus"`
	Doctorname       string    `json:"doctorname"`
	AppoinmentStatus string    `gorm:"column:appointmentstatus" json:"appointmentstatus"`
	Fees             uint32    `gorm:"column:fees" json:"fees"`
}

type Prescription struct {
	gorm.Model
	Doctorid     uint32     `gorm:"column:doctorid" json:"doctorid"`
	Userid       uint32     `gorm:"column:userid" json:"userid"`
	Appoinmentid uint32     `gorm:"column:appoinmentid" json:"appoinmentid"`
	Appoinment   Appoinment `gorm:"foreignKey:Appoinmentid;references:ID"`
	Medicine     string     `json:"medicine" gorm:"not null"`
	Notes        string     `json:"notes"`
}

type Payment struct {
	PaymentID     string `gorm:"primaryKey;type:varchar(255);not null" json:"PaymentID"` // Make PaymentID the primary key if it holds the Stripe ID
	Appoinmentid  uint   `gorm:"not null" json:"appoinmentid"`
	PaymentAmount uint32 `gorm:"not null" json:"PaymentAmount"`
	Status        string `gorm:"not null" json:"Status"`
	ClientSecret  string `gorm:"not null" json:"ClientSecret"`
	PaymentMethod string `gorm:"not null" json:"PaymentMethod"`
	UserID        uint32 `gorm:"not null" json:"UserID"`
}
