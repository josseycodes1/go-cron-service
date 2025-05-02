package main

import (
	"log"
	"os"
	"time"

	"email-cron-service/utils"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	// Load .env file
	err := godotenv.Load("./utils/.env")
	if err != nil {
		log.Println("Error loading .env file:", err)
		return
	}

	// Set up logging
	logFile, err := os.OpenFile("cron.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("Error creating log file:", err)
		return
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	log.Println("Starting Email Cron Service...")

	// Create cron scheduler
	c := cron.New()

	_, err = c.AddFunc("@every 3s", sendEmail) // Runs every 5 minutes
	if err != nil {
		log.Println("Error scheduling job:", err)
		return
	}

	c.Start()
	log.Println("Cron job started. Waiting...")

	select {} // Keeps the program running
}

func sendEmail() {
	log.Println("=== sendEmail() triggered === to", "ayo@pre.game")
	log.Println("Running at:", time.Now().Format(time.RFC1123))

	// Email config from environment
	from := os.Getenv("EMAIL_USERNAME")
	password := os.Getenv("EMAIL_PASSWORD")
	to := []string{"ayo@pre.game"}
	subject := "Assignment done and dusted!"
	body := "Hi, here is my assignment - an automated email sent using a cron job in Go. Let me know when you see it."

	// Create SMTP server instance
	smtp := utils.SmtpServer{
		Host:     "smtp.gmail.com",
		Port:     465,
		Username: from,
		Password: password,
	}

	err := smtp.SendMail(to, from, subject, []byte(body))
	if err != nil {
		log.Println("Failed to send email:", err)
	} else {
		log.Println("Email sent successfully to:", to)
	}
}
