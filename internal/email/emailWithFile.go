package email

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/smtp"
	"net/textproto"
	"path/filepath"
)

func SendEmailWithFile(filePath string) {
	// Sender's email address and password
	senderEmail := "ivocabulary9000@gmail.com"
	senderPassword := "ohfnffkbyiesjcda"

	// Recipient's email address
	recipientEmail := "i.tichkevitch@godeltech.com"

	// SMTP server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	// Message details
	messageSubject := "Test Email with Attachment"
	messageBody := "This is a test email with an attachment."

	// Read the file content
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	// Create a new message with attachment
	message := createMessageWithAttachment(senderEmail, recipientEmail, messageSubject, messageBody, filePath, fileContent)

	// SMTP authentication and connection
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)
	smtpAddress := fmt.Sprintf("%s:%d", smtpHost, smtpPort)

	// Send the email
	err = smtp.SendMail(smtpAddress, auth, senderEmail, []string{recipientEmail}, message)
	if err != nil {
		log.Fatal("Error sending email:", err)
	}

	fmt.Println("Email sent successfully.")
}

func createMessageWithAttachment(senderEmail, recipientEmail, subject, body, filePath string, fileContent []byte) []byte {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	// Create the main email body
	header := make(textproto.MIMEHeader)
	header.Set("Content-Type", "text/plain; charset=utf-8")
	bodyPart, err := writer.CreatePart(header)
	if err != nil {
		log.Fatal("Error creating email body:", err)
	}
	bodyPart.Write([]byte(body))

	// Create the attachment
	header = make(textproto.MIMEHeader)
	header.Set("Content-Type", getMimeType(filePath))
	header.Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filepath.Base(filePath)))
	attachmentPart, err := writer.CreatePart(header)
	if err != nil {
		log.Fatal("Error creating attachment:", err)
	}
	attachmentPart.Write(fileContent)

	// Close the multipart writer
	writer.Close()

	// Create the email message
	message := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: multipart/mixed; boundary=%s\r\n\r\n%s",
		senderEmail, recipientEmail, subject, writer.Boundary(), buf.Bytes()))

	return message
}

func getMimeType(filePath string) string {
	extension := filepath.Ext(filePath)
	switch extension {
	case ".txt":
		return "text/plain"
	case ".pdf":
		return "application/pdf"
	case ".doc", ".docx":
		return "application/msword"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	default:
		return "application/octet-stream"
	}
}
