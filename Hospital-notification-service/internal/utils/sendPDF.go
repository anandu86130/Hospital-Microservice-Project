package utilis

import (
	"fmt"

	"github.com/anandu86130/hospital-notification-service/config"
	"github.com/anandu86130/hospital-notification-service/internal/models"
	"gopkg.in/gomail.v2"
)

// Send email function using GoMail
func SendNotificationToEmail(event models.PaymentEvent, subject, body string) error {
	filePath := "invoice.pdf"

	// Generate the payment invoice
	err := GeneratePaymentInvoicePDF(event.PaymentID, event.AppointmentDate, event.AppointmentStartTime, event.AppointmentEndTime, event.AppointmentID, event.Amount, event.Date, filePath)
	if err != nil {
		return fmt.Errorf("failed to generate PDF: %v", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "sonusuni2255@gmail.com")
	m.SetHeader("To", event.Email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	m.Attach(filePath)

	d := gomail.NewDialer("smtp.gmail.com", 587, config.LoadConfig().APPEMAIL, config.LoadConfig().APPPASSWORD)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
