package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	senha := "TesTandoSe@nah234"
	hash := "$2a$10$ouWHKHtT6Dg0aIedsnLPgOjVQzwI3ThxmKMIPu6oBo4vryn9zwRm."

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
	if err != nil {
		fmt.Println("Senha inválida")
	} else {
		fmt.Println("Senha válida")
	}
}
