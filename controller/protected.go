package controller

import (
	"fmt"
	"net/http"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("user_email").(string)
	fmt.Fprintf(w, "Perfil acessado com sucesso. Email: %s\n", email)
}

func SensitiveHandler(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("user_email").(string)
	fmt.Fprintf(w, "Dados sensíveis acessados com 2FA. Email: %s\n", email)
}
