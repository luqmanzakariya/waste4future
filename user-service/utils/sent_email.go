// using SendGrid's Go Library
// https://github.com/sendgrid/sendgrid-go
package utils

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(toEmail, subject, content string) error {
	from := mail.NewEmail("Waste4Future", os.Getenv("EMAIL_SENDER"))
	to := mail.NewEmail("Recipient", toEmail)
	message := mail.NewSingleEmail(from, subject, to, content, content)

	apiKey := os.Getenv("SENDGRID_API_KEY")
	client := sendgrid.NewSendClient(apiKey)

	response, err := client.Send(message)
	if err != nil {
		return err
	}

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		fmt.Println("Email sent successfully!")
	} else {
		fmt.Printf("Failed to send email. Status Code: %d, Body: %s\n", response.StatusCode, response.Body)
	}

	return nil
}
