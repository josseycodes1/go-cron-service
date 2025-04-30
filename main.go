package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	// create a log file
	logFile, err := os.OpenFile("cron.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	log.Println("Starting Email Cron Service...")

	// setup cron
	c := cron.New()
	_, err = c.AddFunc("@every 10s", sendEmail)
	if err != nil {
		log.Println("Error scheduling job:", err)
		return
	}

	c.Start()
	log.Println("Cron job started. Waiting...")

	// keep the cron service running
	select {}
}

func sendEmail() {
	log.Println("=== sendEmail() triggered ===")
	log.Println("Running job at", time.Now().Format(time.RFC1123))

	// sender email
	from := "adewumijosephine1@gmail.com"
	password := "nvivlcpkyzpbxnpm"

	// receiver email
	to := "akinwumikaliyanu@gmail.com"
	subject := "Assignment done and dusted!"
	body := "Hi, here is my assignment - an automated email sent using a cron job in Go. Let me know when you see it."

	message := []byte("Subject: " + subject + "\r\n\r\n" + body)

	// SMTP server config
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// auth
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// s email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		log.Println("Failed to send email:", err)
	} else {
		log.Println("Email successfully sent to", to)
	}
}
