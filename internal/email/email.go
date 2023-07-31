package email

import (
	"fmt"
	"log"
	"mygoapp/internal/config"
	"net/smtp"
)

func SendEmail(userConfig config.UserType) {
	// Sender's email address and password
	senderEmail := userConfig.SenderEmail
	senderPassword := userConfig.SenderPassword

	// Recipient's email address
	recipientEmail := userConfig.RecipientEmail

	// SMTP server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	// Message details
	messageSubject := "Monthly Invoice"
	messageBody := "There is a file in the attachments with you montly invoice in PDF format."

	// Compose the email message
	message := fmt.Sprintf("Subject: %s\r\n\r\n%s", messageSubject, messageBody)

	// SMTP authentication and connection
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)
	smtpAddress := fmt.Sprintf("%s:%d", smtpHost, smtpPort)

	// Send the email
	err := smtp.SendMail(smtpAddress, auth, senderEmail, []string{recipientEmail}, []byte(message))
	if err != nil {
		log.Fatal("Error sending email:", err)
	}

	fmt.Println("Email sent successfully.")
}
