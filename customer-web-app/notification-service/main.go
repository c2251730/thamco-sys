package main

import (
	"fmt"
	"net/smtp"
	"net/mail"
	"os"
)

type EmailNotification struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
}

func main() {
	fmt.Println("Email Microservice is running...")

	testNotification := EmailNotification{
		To:      "recipient@example.com",
		Subject: "Test Email",
		Text:    "This is a test email notification.",
	}

	err := sendEmail(testNotification)
	if err != nil {
		fmt.Println("Error sending email:", err)
		os.Exit(1)
	}
}

func sendEmail(notification EmailNotification) error {
	smtpServer := "smtp.example.com"
	smtpPort := 587
	smtpUsername := "your-smtp-username"
	smtpPassword := "your-smtp-password"

	from := mail.Address{"", smtpUsername}
	to := mail.Address{"", notification.To}
	subject := notification.Subject
	body := notification.Text

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject

	message := ""
	for key, value := range headers {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	message += "\r\n" + body

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, from.Address, []string{to.Address}, []byte(message))
	if err != nil {
		return err
	}

	fmt.Printf("Email sent to %s with subject: %s\n", notification.To, notification.Subject)
	return nil
}
