package main

import (
	// "log"
	"fmt"
	"os"
	"strconv"

	"github.com/go-mail/mail"
	"github.com/joho/godotenv"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
			return value
	}
	return fallback
}

func main() {

	fmt.Println("Sendgrid keepalive")

  godotenv.Load()

  smtpHost := getEnv("SMTP_HOST", "smtp.sendgrid.net")
	smtpPortString := getEnv("SMTP_PORT", "587")
  smtpUsername := getEnv("SMTP_USERNAME", "apikey")
	smtpPassword := getEnv("SMTP_PASSWORD", "myPassword")
	smtpFrom := getEnv("SMTP_FROM", "myemail@example.com")
	smtpTo := getEnv("SMTP_TO", "myemail@example.com")

	smtpPort, err := strconv.Atoi(smtpPortString)

	if err != nil {
    panic(err)
  }

	m := mail.NewMessage()

	m.SetHeader("From", smtpFrom)
	m.SetHeader("To", smtpTo)
	m.SetHeader("Subject", "SendGrid Keepalive")
	m.SetBody("text/html", "Hello, this is an email sent so as to keep Sendgrid services active")
		
	d := mail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)
	err = d.DialAndSend(m)
	
	if  err != nil {
		panic(err)
	}

	fmt.Println("Email sent")

  
}