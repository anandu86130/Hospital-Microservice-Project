package utilis

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/phpdave11/gofpdf"
)

func GeneratePaymentInvoicePDF(paymentID string, appointmentdate time.Time, appointmentstarttime time.Time, appointmentendtime time.Time, AppointmentID uint, amount uint, date string, filePath string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Set text color to black
	pdf.SetTextColor(0, 0, 0) // RGB for black
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, "Appointment_booking")
	pdf.Ln(6)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, "Calicut, Kerala, 636528")
	pdf.Ln(6)
	pdf.Cell(0, 10, "www.BMH_HOSPITAL.com")
	pdf.Ln(15)

	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, "INVOICE")
	pdf.Ln(15)

	pdf.SetFont("Arial", "", 12)

	// Table Header
	pdf.CellFormat(50, 10, "Field", "1", 0, "L", false, 0, "")
	pdf.CellFormat(100, 10, "Details", "1", 1, "L", false, 0, "")

	// Payment ID
	pdf.CellFormat(50, 10, "Payment ID:", "1", 0, "L", false, 0, "")
	pdf.CellFormat(100, 10, paymentID, "1", 1, "L", false, 0, "")

	// Appointment ID
	pdf.CellFormat(50, 10, "Appointment ID:", "1", 0, "L", false, 0, "")
	pdf.CellFormat(100, 10, strconv.Itoa(int(AppointmentID)), "1", 1, "L", false, 0, "")

	// Amount
	pdf.CellFormat(50, 10, "Amount:", "1", 0, "L", false, 0, "")
	pdf.CellFormat(100, 10, fmt.Sprintf("$%d", amount), "1", 1, "L", false, 0, "")

	// Date
	pdf.CellFormat(50, 10, "Date:", "1", 0, "L", false, 0, "")
	pdf.CellFormat(100, 10, date, "1", 1, "L", false, 0, "")

	// Appointment Date
	pdf.CellFormat(50, 10, "Appointment Date:", "1", 0, "L", false, 0, "")
	pdf.CellFormat(100, 10, appointmentdate.Format("2006-01-02"), "1", 1, "L", false, 0, "")

	// Appointment Start Time
	pdf.CellFormat(50, 10, "Start Time:", "1", 0, "L", false, 0, "")
	pdf.CellFormat(100, 10, appointmentstarttime.Format("03:04 PM"), "1", 1, "L", false, 0, "")

	// Appointment End Time
	pdf.CellFormat(50, 10, "End Time:", "1", 0, "L", false, 0, "")
	pdf.CellFormat(100, 10, appointmentendtime.Format("03:04 PM"), "1", 1, "L", false, 0, "")

	// Text thanking the user
	pdf.SetFont("", "", 10)
	pdf.CellFormat(90, 10, "Thank you for booking appointment. Welcome back again!", "", 0, "C", false, 0, "")
	pdf.CellFormat(90, 10, "Congrats!!! YOUR APPOINTMENT IS CONFIRMED", "", 0, "C", false, 0, "")
	pdf.Ln(12)

	// Output the PDF
	err := pdf.OutputFileAndClose(filePath)
	if err != nil {
		log.Printf("Failed to generate PDF: %v", err)
		return err
	}

	log.Println("PDF generated successfully:", filePath)
	return nil
}
