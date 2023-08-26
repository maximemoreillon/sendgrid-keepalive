package main

import (
    // "log"
    "os"
		"strconv"
		"fmt"


    "github.com/joho/godotenv"
		"github.com/go-mail/mail"

)

func main() {
  err := godotenv.Load()
	if err != nil {
    panic(err)
  }

  smtpHost := os.Getenv("SMTP_HOST")
  smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpFrom := os.Getenv("SMTP_FROM")
	smtpTo := os.Getenv("SMTP_TO")

	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))

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