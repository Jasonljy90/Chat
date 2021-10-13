package main

import (
	"fmt"
	"net/smtp"
	"os"
)

func sendEmail(email string, link string) {
	// Get environment variables
	getEnvVars()
	resetEmailAccount := os.Getenv("Email_Account")
	resetEmailPassword := os.Getenv("Email_Password")

	// sender data
	from := resetEmailAccount //ex: "John.Doe@gmail.com"
	password := resetEmailPassword

	// receiver address
	toEmail := email // ex: "Jane.Smith@yahoo.com"

	to := []string{toEmail}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	// message
	message := []byte("To: " + email + "\r\n" +
		"Subject: Password Reset Request\r\n" +
		"\r\n" +
		"You have received this email because a password reset request for Foodpanda account was received. The reset link will only be valid for 30mins. Click the link to reset your password: \r\n" + link)

	// athentication data
	auth := smtp.PlainAuth("", from, password, host)

	// send mail
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}
