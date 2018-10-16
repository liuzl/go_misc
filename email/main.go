package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"EMAIL_ADDR",
		"EMAIL_PASS",
		"smtp.exmail.qq.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.exmail.qq.com:25",
		auth,
		"FROM_EMAIL",
		[]string{"TO_EMAIL"},
		[]byte("This is the email body."),
	)
	if err != nil {
		log.Fatal(err)
	}
}
