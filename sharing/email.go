package sharing

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(recipient, subject, plainTextContent, htmlContent string) error {
	from := mail.NewEmail("user", os.Getenv("SENDGRID_FROM_EMAIL"))
	to := mail.NewEmail("Recipient", recipient)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	sendGridAPIKey := os.Getenv("SENDGRID_API_KEY")
	client := sendgrid.NewSendClient(sendGridAPIKey)
	response, err := client.Send(message)
	if err != nil {
		return err
	}

	// fmt.Printf("Status Code: %d\n", response.StatusCode)
	// fmt.Printf("Body: %s\n", response.Body)
	// fmt.Printf("Headers: %v\n", response.Headers)

	if response.StatusCode != http.StatusOK {
		fmt.Println("Successfully sent email, status code: ", response.StatusCode)
		return nil
	} else {
		return fmt.Errorf("failed to send email, status code: %d", response.StatusCode)
	}
}