package utility

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

func SendOTPByEmail(Email, Otp string) error {
	// auth := smtp.PlainAuth(
	// 	"rcdyr",
	// 	"sonusuni2255@gmail.com",
	// 	"ggnhnsxxsnvvonmm",
	// 	"smtp.gmail.com",
	// )

	// msg := []byte(Otp)

	// err := smtp.SendMail(
	// 	"smtp.gmail.com:587",
	// 	auth,
	// 	"sonusuni2255@gmail.com",
	// 	[]string{Email},
	// 	msg,
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// }
	auth := smtp.PlainAuth(
		"",
		"sonusuni2255@gmail.com",
		"ggnhnsxxsnvvonmm",
		"smtp.gmail.com",
	)

	// Connect to the SMTP server
	client, err := smtp.Dial("smtp.gmail.com:587")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer client.Close()

	// Upgrade the connection to TLS
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // Skip verification
		ServerName:         "smtp.gmail.com",
	}

	// Log more details to troubleshoot
	fmt.Println("Attempting to send email to:", Email)
	err = client.StartTLS(tlsConfig)
	if err != nil {
		fmt.Println("Error during StartTLS:", err)
		return err
	}

	err = client.Auth(auth)
	if err != nil {
		fmt.Println("Error during Auth:", err)
		return err
	}

	err = client.Mail("sonusuni2255@gmail.com")
	if err != nil {
		fmt.Println("Error during Mail:", err)
		return err
	}

	err = client.Rcpt(Email)
	if err != nil {
		fmt.Println("Error during Rcpt:", err)
		return err
	}

	wc, err := client.Data()
	if err != nil {
		fmt.Println("Error during Data:", err)
		return err
	}
	defer wc.Close()

	msg := []byte("Subject: Your OTP\n\n" + Otp)
	if _, err = wc.Write(msg); err != nil {
		fmt.Println("Error during Write:", err)
		return err
	}

	return nil
}
