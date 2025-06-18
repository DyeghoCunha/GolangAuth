package controller

import (
	"encoding/json"
	"net/http"

	"github.com/dyeghocunha/golang-auth/util"
)

type Setup2FAResponse struct {
	QRCode string `json:"qr_code"`
	Secret string `json:"secret"`
}

type Verify2FARequest struct {
	Code   string `json:"code"`
	Secret string `json:"secret"`
}

func Setup2FAHandler(w http.ResponseWriter, r *http.Request) {
	userEmail := r.Context().Value("user_email").(string)
	qr, secret, err := util.GenerateTOTPKey("PokedexGo", userEmail)
	if err != nil {
		http.Error(w, "Erro ao gerar 2FA", http.StatusInternalServerError)
		return
	}

	resp := Setup2FAResponse{QRCode: qr, Secret: secret}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func Verify2FAHandler(w http.ResponseWriter, r *http.Request) {
	var req Verify2FARequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Requisição inválida", http.StatusBadRequest)
		return
	}

	if util.ValidateTOTPCode(req.Secret, req.Code) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Código válido"))
	} else {
		http.Error(w, "Código inválido", http.StatusUnauthorized)
	}
}
