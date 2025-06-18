package main

import (
	"log"
	"net/smtp"
	"os"
)

func main() {
	//from := "autistamajestoso@gmail.com"
	from := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	//pass := "ydlcivwauzhlwlyi" // sem espaços

	to := "dyeghocunha@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Test Email\n\n" +
		"This is a test email sent from Go.\n"

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from,
		[]string{to},
		[]byte(msg),
	)
	if err != nil {
		log.Fatal("Error sending email:", err)
	} else {
		log.Println("Email sent successfully to", to)
	}
}
