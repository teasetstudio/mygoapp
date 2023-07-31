package email

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"mygoapp/internal/config"
	"net/smtp"
	"os"
	"path/filepath"
)

// The function you provided appears to have the necessary components to send an email with
// a PDF attachment. However, there is a problem with the way you include the file content
// in the email message.

// In the current code, you're directly including the byte content of the PDF file in the
// message string using the %s format specifier. This will not work as expected since the
// content may contain special characters that can interfere with the email message formatting.

// To fix this issue, you should encode the file content using the appropriate encoding
// method. In this case, you can use base64 encoding to ensure proper representation of
// the PDF file content in the email message. Here's an updated version of your function
// that incorporates base64 encoding:

// In this updated version, I've made changes to properly encode the PDF file content using
// base64.StdEncoding.EncodeToString. This ensures that the attachment is correctly represented
// in the email message.

func SendFile(filePath string, userConfig *config.UserType) {
	// SMTP server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	smtpUsername := userConfig.SenderEmail
	smtpPassword := userConfig.SenderPassword

	// Sender and recipient email addresses
	from := userConfig.SenderEmail
	to := userConfig.RecipientEmail

	// Create the email message
	subject := "Email with PDF attachment"
	body := "Please find the attached PDF file."
	attachmentPath := filePath

	// Read the file content
	file, err := os.Open(attachmentPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Create the message parts
	header := make(map[string]string)
	header["From"] = from
	header["To"] = to
	header["Subject"] = subject

	message := ""

	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	// Create the MIME parts
	boundary := "boundary"

	message += fmt.Sprintf("MIME-Version: 1.0\r\n"+
		"Content-Type: multipart/mixed; boundary=%s\r\n"+
		"\r\n"+
		"--%s\r\n"+
		"Content-Type: text/plain; charset=utf-8\r\n"+
		"\r\n"+
		"%s\r\n"+
		"\r\n"+
		"--%s\r\n"+
		"Content-Type: application/pdf\r\n"+
		"Content-Disposition: attachment; filename=%s\r\n"+
		"Content-Transfer-Encoding: base64\r\n"+
		"\r\n"+
		"%s\r\n"+
		"--%s--\r\n",
		boundary, boundary, body, boundary, filepath.Base(attachmentPath), base64.StdEncoding.EncodeToString(content), boundary)

	// SMTP authentication setup
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// Connect to the SMTP server
	client, err := smtp.Dial(fmt.Sprintf("%s:%d", smtpHost, smtpPort))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Enable TLS encryption
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // Use this if you're unable to specify ServerName
		// ServerName:         "smtp.example.com", // Uncomment and provide the correct server name if available
	}

	// Start the SMTP connection
	if err = client.StartTLS(tlsConfig); err != nil {
		log.Fatal(err)
	}

	// Authenticate with the server
	if err = client.Auth(auth); err != nil {
		log.Fatal(err)
	}

	// Set the sender and recipient
	if err = client.Mail(from); err != nil {
		log.Fatal(err)
	}
	if err = client.Rcpt(to); err != nil {
		log.Fatal(err)
	}

	// Send the email message
	wc, err := client.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()

	_, err = wc.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Email sent successfully!")
}
