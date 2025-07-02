package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/smtp"
	"os"
)

func SendTestEmail() {
	from := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	to := "dyeghocunha@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Teste SMTP\n\n" +
		"Funcionando dentro do container com Go puro 🎉"

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from,
		[]string{to},
		[]byte(msg),
	)
	if err != nil {
		log.Println("Erro ao enviar:", err)
	} else {
		log.Println("✅ Email enviado com sucesso!")
	}
}

func CheckPasswordHash(hash, password string) bool {
	fmt.Println(hash)
	fmt.Println(password)
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
