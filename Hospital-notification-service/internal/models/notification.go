package models

import "time"

type Notification struct {
	ID        string    `bson:"_id,omitempty"`
	UserID    string    `bson:"user_id"`
	Message   string    `bson:"message"`
	CreatedAt time.Time `bson:"created_at"`
	PDFPath   string    `bson:"pdf_path,omitempty"`
}
