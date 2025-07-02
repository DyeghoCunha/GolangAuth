package controller

import (
	"encoding/json"
	"net/http"

	"github.com/dyeghocunha/golang-auth/repository"
	"github.com/pquerna/otp/totp"
)

type Verify2FARequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func Verify2FAHandler(w http.ResponseWriter, r *http.Request) {
	var req Verify2FARequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Email == "" || req.Code == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	valid := totp.Validate(req.Code, user.TwoFASecret)
	if !valid {
		http.Error(w, "Invalid 2FA code", http.StatusUnauthorized)
		return
	}
	err = repository.Enable2Fa(req.Email)
	if err != nil {
		http.Error(w, "Error enabling 2FA: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "2FA verification successful",
	})

}
