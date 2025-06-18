package main

import (
	"fmt"
	"net/http"

	"github.com/dyeghocunha/golang-auth/util"
)

func main() {
	util.SendTestEmail() // dispara ao subir

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Auth service está online!")
	})

	fmt.Println("🚀 Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}
